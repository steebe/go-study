package maps

import "strings"

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, v := range strings.Fields(s) {
		count := result[v]
		if count == 0 {
			result[v] = 1
		} else {
			count++
			result[v] = count
		}
	}
	return result
}
