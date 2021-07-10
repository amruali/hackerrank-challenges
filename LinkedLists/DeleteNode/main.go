package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		llist = llist.next
		return llist
	}
	var cnt int32 = 0
	curr := llist
	var prev *SinglyLinkedListNode
	for cnt <= position && curr != nil {
		if cnt == position {
			prev.next = curr.next
			return llist
		}
		prev = curr
		curr = curr.next
		cnt++
	}
	return llist
}

func main() {}
