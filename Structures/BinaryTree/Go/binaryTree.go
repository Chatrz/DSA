package main

import "fmt"

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	key    int
}

type Tree struct {
	Root *Node
}

func CreateTree() *Tree {
	tree := &Tree{
		Root: nil,
	}
	return tree
}

func (tree *Tree) Insert(data int) {
	newNode := &Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		key:    data,
	}
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		tmp := tree.Root
		for {
			holder := tmp
			if data > tmp.key {
				tmp = tmp.Right
				if tmp == nil {
					holder.Right = newNode
					newNode.Parent = holder
					break
				}
			} else {
				tmp = tmp.Left
				if tmp == nil {
					holder.Left = newNode
					newNode.Parent = holder
					break
				}
			}
		}
	}
}

func (node *Node) Search(key int) *Node {
	if node == nil {
		fmt.Println("KEY DOES NOT EXIST !")
		return nil
	}
	if node.key == key {
		return node
	}
	if key > node.key {
		return node.Right.Search(key)
	} else {
		return node.Left.Search(key)
	}
}

func (tree *Tree) GetMax() int {
	tmp := tree.Root
	for {
		holder := tmp
		tmp = tmp.Right
		if tmp == nil {
			return holder.key
		}
	}
}

func (tree *Tree) GetMin() int {
	tmp := tree.Root
	for {
		holder := tmp
		tmp = tmp.Left
		if tmp == nil {
			return holder.key
		}
	}
}

//TODO completing remove and adding show
func (tree *Tree) Remove(key int) bool {
	node := tree.Root.Search(key)
	if node == nil {
		fmt.Println("KEY DOES NOT EXIST !")
		return false
	} else {
		if node.Left != nil {
			node.Left.Parent = node.Parent
		}
		if node.Right != nil {
			node.Right.Parent = node.Parent
		}
		return true
	}
}

func main() {
	tree := CreateTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(-1)
	tree.Insert(-6)
	fmt.Println(tree.Root.Search(4).key)
	tree.Remove(4)
	fmt.Println(tree.Root.Search(5).Parent.key)
}
