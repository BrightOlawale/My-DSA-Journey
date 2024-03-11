package march

func customSortString(order string, s string) string {
	charCount := make(map[rune]int)

	// Count occurrences of characters in s
	for _, char := range s {
		charCount[char]++
	}

	result := ""

	// Append characters from order to result
	for _, char := range order {
		count, found := charCount[char]
		if found {
			for i := 0; i < count; i++ {
				result += string(char)
			}
			delete(charCount, char)
		}
	}

	// Append remaining characters from s
	for char, count := range charCount {
		for i := 0; i < count; i++ {
			result += string(char)
		}
	}

	return result
}