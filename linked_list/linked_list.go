package linked_list

import (
	"fmt"
)

type LinkedList struct {
	HeadNodeRef *Node
	length      int
}

type Node struct {
	Key     int
	NextRef *Node
}

func (l *LinkedList) Prepend(n *Node) {
	n.NextRef = l.HeadNodeRef
	l.HeadNodeRef = n
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.HeadNodeRef
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.Key)
		toPrint = toPrint.NextRef
		l.length--
	}
	fmt.Println("\n")
}

func (l *LinkedList) Delete(value int) {
	if l.length == 0 {
		return
	}
	if l.HeadNodeRef.Key == value {
		l.HeadNodeRef = l.HeadNodeRef.NextRef
		l.length--
		return
	}
	previousNode := l.HeadNodeRef
	for previousNode.NextRef.Key != value {
		if previousNode.NextRef.NextRef == nil {
			return
		}
		previousNode = previousNode.NextRef
	}
	previousNode.NextRef = previousNode.NextRef.NextRef
	l.length--
}
