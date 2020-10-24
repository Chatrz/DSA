package binaryTree

import "testing"

func TestCreateTree(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if tree.Root.Right.Left.key!=24{
		t.Error("CreateTree not working !!")
	}
}
func TestTree_Insert(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if tree.Root.Right.Left.key!=24{
		t.Error("Tree_Insert not working !!")
	}
}

func TestNode_Search(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if tree.Root.Search(24)==nil{
		t.Error("Node_Search not working !!")
	}
}

func TestNode_GetMax(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if tree.Root.GetMax().key!=25{
		t.Error("GetMax not working !!")
	}
}

func TestNode_GetMin(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if tree.Root.GetMin().key!=20{
		t.Error("GetMin not working !!")
	}
}
func TestNode_IsLeaf(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if !tree.Root.Right.Left.IsLeaf(){
		t.Error("IsLeaf not working !!")
	}
}
func TestNode_GetPredecessor(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(18)
	tree.Insert(19)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(5)
	if tree.Root.GetPredecessor().key!=19{
		t.Error("GetPredecessor not working !!")
	}
}
func TestNode_GetSuccessor(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(18)
	tree.Insert(19)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(5)
	if tree.Root.GetSuccessor().key!=23{
		t.Error("GetSuccessor not working !!")
	}
}

func TestNode_IsRightChild(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	if !tree.Root.Right.IsRightChild(){
		t.Error("IsRightChild not working !!")
	}
}

func TestTree_DeleteUseKey(t *testing.T) {
	tree := CreateTree()
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(24)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(23)
	tree.Insert(27)
	tree.Insert(15)
	tree.Insert(17)
	tree.Insert(18)
	tree.Insert(19)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(5)
	tree.DeleteUseKey(24)
	if tree.Root.Search(24)!=nil{
		t.Error("DeleteUseKey not working !!")
	}
}
