package main

import (
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Add(data string) {
	node := &Node{Data: data}
	if list.Head == nil {
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (list *LinkedList) Remove(data string) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (list *LinkedList) length() int {
	result := 0
	for current := list.Head; current != nil; current = current.Next {
		result++
	}
	return result
}

func main() {
	names := LinkedList{}

	names.Add("Tom")
	names.Add("Alice")
	names.Add("Jacky")
	names.Add("Lucky")

	for current := names.Head; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}
	fmt.Printf("Length: %d\n", names.length())

	names.Remove("Jacky")

	for current := names.Head; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}

	fmt.Printf("Length: %d\n", names.length())
}
