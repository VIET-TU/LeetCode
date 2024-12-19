package _1512_Number_of_Good_Pairs

func numIdenticalPairs(nums []int) int {
	result := 0
	numberCountMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count, isExist := numberCountMap[nums[i]]
		if isExist {
			result += count
			numberCountMap[nums[i]] = count + 1
		} else {
			numberCountMap[nums[i]] = 1
		}
	}
	return result
}
