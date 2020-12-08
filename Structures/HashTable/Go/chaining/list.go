package chaining

import (
	"fmt"
)

type Node struct {
	Key  Item
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func CreateLinkedList() *LinkedList {
	list := LinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	return &list
}

func (list *LinkedList) AddBack(key Item) {
	if list.Size == 0 {
		list.firstInitialize(key)
	} else if list.Size == 1 {
		node := &Node{
			Key:  key,
			Prev: nil,
			Next: list.Head,
		}
		list.Head.Prev = node
		list.Tail = list.Head
		list.Head = node
	} else {
		node := &Node{
			Key:  key,
			Prev: nil,
			Next: list.Head,
		}
		list.Head.Prev = node
		list.Head = node
	}
	list.Size = list.Size + 1
}

// adding first element to linked list
func (list *LinkedList) firstInitialize(key Item) {
	node := &Node{
		Key:  key,
		Prev: nil,
		Next: nil,
	}
	list.Head = node
}

func (list *LinkedList) Delete(index int) *Node {
	node := list.Search(index)
	if node == nil {
		return nil
	} else {
		if node.Next == nil && node.Prev == nil { // it is the only node
			list.Head = nil
		} else if node.Next == nil { //node is tail of list
			list.Tail = node.Prev
			node.Prev.Next = nil
		} else if node.Prev == nil { //nod is head of list
			list.Head = node.Next
			node.Next.Prev = nil
		} else {
			node.Next.Prev = node.Prev
			node.Prev.Next = node.Next
		}
		list.Size = list.Size - 1
		return node
	}
}

func (list *LinkedList) Display() {
	rooter := list.Head
	fmt.Print("list is :   ")
	for rooter != nil {
		fmt.Print(rooter.Key)
		fmt.Print("  ")
		rooter = rooter.Next
	}
	fmt.Println()
}

func (list *LinkedList) Search(index int) *Node {
	node := list.Head
	if node.Key.key == index {
		return node
	}
	for node != list.Tail {
		node = node.Next
		if node == nil || node.Key.key == index {
			return node
		}
	}
	return nil
}
