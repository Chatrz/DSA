package linkedList

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	list:=CreateLinkedList()
	if list.size!=0{
		t.Error()
	}
}

func TestLinkedList_AddBack(t *testing.T) {
	link := CreateLinkedList()
	link.AddBack(20)
	link.AddBack(30)
	link.AddBack(40)
	if link.head.key!=40{
		t.Error("AddBack func is not working correctly ! ")
	}
}
func TestLinkedList_AddFront(t *testing.T) {
	link:=CreateLinkedList()
	link.AddFront(20)
	link.AddFront(30)
	link.AddFront(40)
	if link.tail.key!=40{
		t.Error("AddFront func is not working correctly !")
	}
}

func TestLinkedList_Search(t *testing.T) {
	link:=CreateLinkedList()
	link.AddFront(20)
	link.AddFront(30)
	link.AddFront(40)
	if link.Search(2).key!=30{
		t.Error("search func is not working correctly !")
	}
}

func TestLinkedList_Delete(t *testing.T) {
	link:=CreateLinkedList()
	link.AddFront(20)
	link.AddFront(30)
	link.AddFront(40)
	link.Delete(3)
	link.Display()
	if link.tail.key==40{
		t.Error("Delete func is not working correctly !")
	}
}


