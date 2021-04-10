package main

import (
	"strings"
)

func matchingStrings(stringss []string, queries []string) []int32 {
	res := make([]int32, len(queries))
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(stringss); j++ {
			res[i] += int32(strings.Count(queries[i], stringss[j]))
		}
	}
	return res
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
