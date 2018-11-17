package removeduplicate

// RemoveDuplicate is a function to remove duplicate word in slice
func RemoveDuplicate(input []string) []string {
	mapStr := map[string]bool{}
	output := []string{}
	for _, v := range input {
		mapStr[v] = false
	}
	for i := 0; i < len(input); i++ {
		for k, v := range mapStr {
			if input[i] == k && !v {
				output = append(output, input[i])
				mapStr[k] = true
			}
		}
	}
	return output
}
