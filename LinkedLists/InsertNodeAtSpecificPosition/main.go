package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{data: data}
	if position == 0 {
		newNode.next = llist
		return newNode
	}

	var cnt int32 = 0
	curr := llist
	var prev *SinglyLinkedListNode
	for cnt <= position && curr != nil {
		if cnt == position {
			newNode.next = curr
			prev.next = newNode
			return llist
		}
		prev = curr
		curr = curr.next
		cnt++
	}
	prev.next = newNode
	return llist
}

func main() {
	/*
	   Auto Generated IO
	*/
}
