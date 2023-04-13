package lists

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Length() int {
	length := 0
	for node := ll.head; node != nil; node = node.next {
		length++
	}
	return length
}

func (ll *LinkedList) Head() (int, error) {
	if ll.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	return ll.head.value, nil
}

func (ll *LinkedList) GetAt(index int) (int, error) {
	if ll.head == nil {
		return 0, fmt.Errorf("index %d out of range", index)
	}
	node := ll.head
	for i := 1; i <= index; i++ {
		node = node.next
		if node == nil {
			return 0, fmt.Errorf("index %d out of range", index)
		}
	}
	return node.value, nil
}

func (ll *LinkedList) Tail() (int, error) {
	if ll.head == nil {
		return 0, fmt.Errorf("empty list")
	}
	node := ll.head
	for node.next != nil {
		node = node.next
	}
	return node.value, nil
}

func (ll *LinkedList) InsertBeforeHead(value int) {
	currentHead := ll.head
	newNode := &Node{value, currentHead}
	ll.head = newNode
}

func (ll *LinkedList) InsertAfterTail(value int) {
	newNode := &Node{value, nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		node := ll.head
		for node.next != nil {
			node = node.next
		}
		node.next = newNode
	}
}

func (ll *LinkedList) DeleteHead() error {
	if ll.head == nil {
		return fmt.Errorf("empty list")
	}
	ll.head = ll.head.next
	return nil
}

func (node *Node) String() string {
	return strconv.Itoa(node.value)
}

func (ll *LinkedList) String() string {
	builder := strings.Builder{}
	builder.WriteByte('[')
	for node := ll.head; node != nil; node = node.next {
		builder.WriteString(node.String())
		if node.next != nil {
			builder.WriteString(", ")
		}
	}
	builder.WriteByte(']')
	return builder.String()
}

