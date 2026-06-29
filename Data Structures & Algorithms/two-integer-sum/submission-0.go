func twoSum(nums []int, target int) []int {
    numsLen := len(nums)
	result := make([]int, 2)
	for i := 0; i < numsLen; i++ {
		for i2 := i + 1; i2 < numsLen; i2++ {
			if nums[i] + nums[i2] == target {
				result[0] = i
				result[1] = i2
			}
		}
	}
	return result
}
