package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func reversePrint(llist *SinglyLinkedListNode) {

	s := make([]int32, 0)
	sLen := 0
	curr := llist
	for curr != nil {
		s = append(s, curr.data)
		curr = curr.next
		sLen++
	}
	for i := sLen - 1; i >= 0; i-- {
		fmt.Println(s[i])
	}

}

func main() {}
