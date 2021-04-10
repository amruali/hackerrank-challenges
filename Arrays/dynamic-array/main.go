package main

func dynamicArray(n, q int32, queries [][]int32) []int32 {
	var i int32
	var lastAnswer int32
	var temp int32
	res := make([]int32, 0)
	myMap := make(map[int32][]int32)
	for i = 0; i < q; i++ {
		temp = 0
		x := queries[i][1]
		y := queries[i][2]
		switch Op := queries[i][0]; Op {
		case 1:
			idx := (x ^ lastAnswer) % n
			myMap[idx] = append(myMap[idx], y)
		case 2:
			idx := (x ^ lastAnswer) % n
			temp = myMap[idx][y%int32(len(myMap[idx]))]
		}

		if temp != 0 {
			lastAnswer = temp
			res = append(res, lastAnswer)
		}

	}

	return res
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}
