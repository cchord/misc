package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func schedule(adjTable map[string][]string) {
	inDegrees := make(map[string]int, 0)
	for node, adjNodes := range adjTable {
		if _, ok := inDegrees[node]; !ok {
			inDegrees[node] = 0
		}

		for _, adjNode := range adjNodes {
			inDegrees[adjNode]++
		}
	}

	queue := make([]string, 0, len(adjTable))
	for node, degree := range inDegrees {
		if degree != 0 {
			continue
		}
		queue = append(queue, node)
	}

	jobc := make(chan []string, 1)
	done := make(chan struct{})
	go func() {
		defer close(done)

		for joblist := range jobc {
			var wg sync.WaitGroup
			for _, node := range joblist {
				wg.Add(1)
				go func(node string) {
					defer wg.Done()
					start := time.Now()
					time.Sleep(time.Second * 2)
					fmt.Printf("Node %s finished, started at %v, cost %v\n", node, start, time.Since(start))
				}(node)
			}
			wg.Wait()
		}
	}()

	for len(queue) != 0 {
		jobc <- queue

		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[i]
			for _, adjNode := range adjTable[node] {
				inDegrees[adjNode]--
				if inDegrees[adjNode] == 0 {
					queue = append(queue, adjNode)
				}
			}
		}
		queue = queue[count:]
	}
	close(jobc)

	<-done
}

func main() {
	input := `{"A":["B","C"],"B":["D","E"],"D":["F"],"E":["F"]}`
	dependencies := make(map[string][]string, 0)
	err := json.Unmarshal([]byte(input), &dependencies)
	if err != nil {
		panic(err)
	}

	adjacentTable := make(map[string][]string, 0)
	for node, parents := range dependencies {
		for _, parent := range parents {
			adjacentTable[parent] = append(adjacentTable[parent], node)
		}
	}

	schedule(adjacentTable)
}
