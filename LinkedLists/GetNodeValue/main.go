package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	_ "strings"
)

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
    head *SinglyLinkedListNode
    tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
    node := &SinglyLinkedListNode {
        next: nil,
        data: nodeData,
    }

    if singlyLinkedList.head == nil {
        singlyLinkedList.head = node
    } else {
        singlyLinkedList.tail.next = node
    }

    singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}


func getNode(llist *SinglyLinkedListNode, positionFromTail int32) int32 {
    // Write your code here
    s := make([]int32, 0)
    curr := llist
    for curr != nil {
        s = append(s, curr.data)
        curr = curr.next
    }
    return s[len(s) - int(positionFromTail) - 1]
}

func main() {
