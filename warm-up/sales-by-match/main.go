package main

func sockMerchant(n int32, ar []int32) (result int32) {
	m := make(map[int32]int32, n)
	var i int32
	for i = 0; int(i) < len(ar); i++ {
		m[ar[i]] += 1
	}
	for _, val := range m {
		var temp int32
		temp = val / 2
		if temp > 0 {
			result += temp
		}
	}
	return
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
