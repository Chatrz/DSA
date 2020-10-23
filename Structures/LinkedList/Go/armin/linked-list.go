package linkedList

import (
	"fmt"
)

type Node struct {
	key  interface{}
	next *Node
	prev *Node
}

type LinkedList struct {
	size int
	head *Node
	tail *Node
}

func CreateLinkedList() *LinkedList {
	list := LinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
	return &list
}

func (list *LinkedList) AddFront(key interface{}) {
	if list.size == 0 {
		list.firstInitialize(key)
	} else if list.size == 1 {
		node := &Node{
			key:  key,
			prev: list.head,
			next: nil,
		}
		list.head.next = node
		list.tail = node
	} else {
		node := &Node{
			key:  key,
			prev: list.tail,
			next: nil,
		}
		list.tail.next = node
		list.tail = node
	}
	list.size = list.size + 1
}

func (list *LinkedList) AddBack(key interface{}) {
	if list.size == 0 {
		list.firstInitialize(key)
	} else if list.size == 1 {
		node := &Node{
			key:  key,
			prev: nil,
			next: list.head,
		}
		list.head.prev = node
		list.tail = list.head
		list.head = node
	} else {
		node := &Node{
			key:  key,
			prev: nil,
			next: list.head,
		}
		list.head.prev = node
		list.head = node
	}
	list.size = list.size + 1
}

// adding first element to linked list
func (list *LinkedList) firstInitialize(key interface{}) {
	node := &Node{
		key:  key,
		prev: nil,
		next: nil,
	}
	list.head = node
}

func (list *LinkedList) Delete(index int) *Node {
	node := list.Search(index)
	if node == nil {
		return nil
	} else {
		list.size = list.size - 1
		if node.next == nil { //node is tail of list
			list.tail = node.prev
			node.prev.next = nil
		} else if node.prev == nil { //nod is head of list
			list.head = node.next
			node.next.prev = nil
		} else {
			node.next.prev = node.prev
			node.prev.next = node.next
		}
		return node
	}
}

func (list *LinkedList) Display() {
	rooter := list.head
	fmt.Print("HEAD OF LIST : ")
	for {
		fmt.Println(rooter.key)
		rooter = rooter.next
		if rooter.next == nil {
			fmt.Print("TAIL OF LIST : ")
			fmt.Println(rooter.key)
			break
		}
	}

}

func (list *LinkedList) Search(index int) *Node {
	if index < 0 || index > list.size {
		fmt.Println("INVALID INDEX !")
		return nil
	} else {
		node := list.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		return node
	}
}

func main() {
	link := CreateLinkedList()
	link.AddFront("head")
	link.AddFront(30)
	link.AddFront(40)
	link.AddFront(50)
	link.AddBack("new head")
	link.AddBack("new new head")
	link.AddFront("tail")
	link.Display()
	fmt.Println("######################################################3")
	link.Delete(10)
	link.Display()
}
