package main

import "testing"

func TestNextPerm(t *testing.T) {
	nums := []int{1, 2, 3, 3}

	for i := 0; i < 24; i++ {
		nums = nextPerm(nums)
		t.Error(nums)
	}

	//t.Error(nextPerm([]int{1, 2, 3, 4}))
	//t.Error(nextPerm([]int{4, 3, 2, 1}))
}
