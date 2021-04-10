package main


func reverseArray(a []int32) []int32 {
	n := len(a)
	for i := 0; i < n/2; i++ {
		temp := a[i]
		a[i] = a[n-i-1]
		a[n-i-1] = temp
	}
	return a
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
