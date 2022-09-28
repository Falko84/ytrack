package main

import "fmt"

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail = n
	}
	n.Next = l.Head
	l.Head = n
}

func ListSize(l *List) int {
	var res int
	it := l.Head
	for it != nil {
		res++
		it = it.Next
	}
	return res
}
