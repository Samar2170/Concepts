package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	n := Node{val: val, next: nil}
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) InsertAt(pos int, val int) {
	n := Node{val: val, next: nil}

	if pos == 0 {
		n.next = l.head
		l.head = &n
		l.len++
		return
	}

	ptr := l.head
	for i := 0; i < pos-1; i++ {
		ptr = ptr.next
	}
	n.next = ptr.next
	ptr.next = &n
	l.len++

}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("Empty List")
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}

func (l *LinkedList) SearchNode(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.val == val {
			return ptr.val
		}
		ptr = ptr.next
	}
	return -1
}

func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	for i := 0; i < pos-1; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *LinkedList) DeleteAt(pos int) error {
	prevNode := l.GetAt(pos - 1)
	prevNode.next = l.GetAt(pos + 1)
	l.len--
	return nil
}
