package _528_Shuffle_String

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices))
	for index, value := range indices {
		result[value] = s[index]
	}
	return string(result)
}
