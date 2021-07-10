package main

func removeDuplicates(llist *SinglyLinkedListNode) *SinglyLinkedListNode {

	curr := llist
	for curr.next != nil {
		if curr.data == curr.next.data {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
	return llist
}

func main() {
	/*
		Auto Generated IO
	*/
}
