package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	tempSlice := make([]int32, 0)
	curr := llist
	for curr != nil {
		tempSlice = append(tempSlice, curr.data)
		curr = curr.next
	}
	newList := &SinglyLinkedList{}
	newList.head = &SinglyLinkedListNode{data: tempSlice[len(tempSlice)-1]}
	currentHeadNode := newList.head
	for i := len(tempSlice) - 2; i >= 0; i-- {
		newNode := &SinglyLinkedListNode{data: tempSlice[i]}
		currentHeadNode.next = newNode
		currentHeadNode = currentHeadNode.next
	}
	return newList.head
}

func main() {

}
