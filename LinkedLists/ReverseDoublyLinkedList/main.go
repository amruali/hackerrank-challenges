package main

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {

	curr := llist
	prev := llist.prev

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	return prev

}

func main() {}
