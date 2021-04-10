package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func rotateLeft(d int32, arr []int32) []int32 {
    for ; d > 0 ; d-- {
        left := arr[0]
        arr = arr[1:]
        arr = append(arr, left)        
    }
    return arr
}

func main() {
	/*
	 HackerRank-Auto-Generated-IO & Function Call
	*/
}