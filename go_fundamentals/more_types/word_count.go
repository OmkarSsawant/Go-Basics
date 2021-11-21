package moretypes

import "strings"

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)
	words := strings.Fields(s) //similar to strings.Split(`\s+`)
	for _, word := range words {
		countMap[word] = countMap[word] + 1
	}
	return countMap
}
