package main

func hourglassSum(arr [][]int32) (answer int32) {
	var temp int32
	answer = -1e9
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			temp = arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] +
				arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			if temp >= answer {
				answer = temp
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
