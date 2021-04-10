package main

func jumpingOnClouds(c []int32) (res int32) {
	for i := 0; i < len(c); i++ {
		if c[i] == 1 && i == len(c)-1 {
			return
		}
		if c[i] == 0 {
			if i+1 == len(c)-1 {
				if c[i+1] == 0 {
					res++
					return
				}
				return
			}
			for j := i; j < len(c)-1; j++ {
				if c[j+1]+c[j+2] == 0 {
					res++
					i += 1
					break
				} else {
					res++
					break
				}

			}

		}
	}
	return

}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
