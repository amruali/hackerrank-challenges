package main

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	if llist == nil {
		return &DoublyLinkedListNode{data: data}
	}
	newNode := &DoublyLinkedListNode{data: data}
	if llist.data > data {
		newNode.next = llist
		return newNode
	}
	if llist.next == nil {
		llist.next = newNode
		return llist
	}
	curr := llist
	for curr.next != nil {
		if curr.next.data > data {
			newNode.next = curr.next
			curr.next = newNode
			return llist
		}
		curr = curr.next
	}
	curr.next = newNode
	return llist
}

func main() {

}
