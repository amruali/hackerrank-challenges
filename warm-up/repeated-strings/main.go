package main

import (
	"strings"
)

func repeatedString(s string, n int64) (res int64) {
	if len(s) == 1 {
		if s[0] == 'a' {
			return n
		}
		return 0
	} else if int64(len(s)) > n {
		var i int64
		for i = 0; i < n; i++ {
			if s[i] == 'a' {
				res++
			}
		}
		return res
	} else {
		var rep = n / int64(len(s))
		var extra = n - (rep * int64(len(s)))
		return int64(strings.Count(s, "a"))*rep + int64(strings.Count(s[:extra], "a"))
	}

}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
