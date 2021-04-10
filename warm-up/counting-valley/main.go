package main

func countingValleys(steps int32, path string) (res int32) {
	temp := 0
	for idx := 0; idx < len(path); idx++ {
		if path[idx] == 'U' {
			temp += 1
			if temp == 0 {
				res++
			}
		} else {
			temp -= 1
		}
	}
	return
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
