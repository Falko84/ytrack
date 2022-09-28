package main

import "fmt"

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{
		Data: data_ref,
		Next: &NodeI{},
	}
	if l == nil {
		return n
	}
	if data_ref < l.Data {
		n.Next = l
		return n
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < data_ref {
		iterator = iterator.Next
	}
	n.Next = iterator.Next
	iterator.Next = n
	return l
}
