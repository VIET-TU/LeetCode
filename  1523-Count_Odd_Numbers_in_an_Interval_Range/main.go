package _1523_Count_Odd_Numbers_in_an_Interval_Range

func countOdds(low int, high int) int {
	result := 0
	numberOfElement := high - low + 1
	result = numberOfElement / 2
	if numberOfElement%2 == 1 && low%2 == 1 {
		result++
	}
	return result
}
