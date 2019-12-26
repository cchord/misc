package main

func nextPerm(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	i := n - 1
	for i > 0 {
		j := i
		i--
		if !(nums[i] > nums[j]) {
			continue
		}

		k := n - 1
		for nums[k] > nums[i] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
		reverse(nums[j:])
		return nums
	}
	reverse(nums)
	return nums
}

func reverse(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j := 0, n-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
