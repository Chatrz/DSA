package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	key    int
}

type Tree struct {
	Root *Node
}

type TreePicture struct {
	pic string
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

func (root *Node) GetMax() *Node {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Right
		if tmp == nil {
			return holder
		}
	}
}

func (root *Node) GetMin() *Node {
	tmp := root
	for {
		holder := tmp
		tmp = tmp.Left
		if tmp == nil {
			return holder
		}
	}
}

func (node *Node) GetPredecessor() *Node {
	if node.Left != nil {
		return node.Left.GetMax()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for ; tmp2 != nil; {
			if tmp != tmp2.Left {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func (node *Node) GetSuccessor() *Node {
	if node.Right != nil {
		return node.Right.GetMin()
	} else {
		tmp := node
		tmp2 := tmp.Parent
		for ; tmp2 != nil; {
			if tmp != tmp2.Right {
				break
			}
			tmp = tmp2
			tmp2 = tmp2.Parent
		}
		return tmp2
	}
}

func DisplayTree(res *TreePicture, padding string, pointer string, node *Node) {
	if node != nil {
		res.pic = res.pic + padding
		res.pic = res.pic + pointer
		res.pic = res.pic + strconv.Itoa(node.key)
		res.pic = res.pic + "\n"

		paddingBuilder := padding + "│  "
		paddingForBoth := paddingBuilder
		pointerForRight := "└──"
		pointerForLeft := ""
		if node.Right != nil {
			pointerForLeft = "├──"
		} else {
			pointerForLeft = "└──"
		}

		DisplayTree(res, paddingForBoth, pointerForLeft, node.Left)
		DisplayTree(res, paddingForBoth, pointerForRight, node.Right)
	}
}

//TODO adding successor and predecessor finder for keys (after saturday class)
//correcting remove

func main() {
	tree := CreateTree()
	tree.Insert(-300)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(-65)
	tree.Insert(3)
	tree.Insert(-60)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(-70)
	tree.Insert(-69)
	tree.Insert(-400)
	tree.Insert(-1)
	tree.Insert(-6)
	tree.Insert(-200)
	tree.Insert(7)
	pic := &TreePicture{pic: ""}
	DisplayTree(pic, "", "", tree.Root)
	fmt.Println(pic.pic)
	fmt.Println(tree.Root.Search(-60).GetPredecessor().key)
	fmt.Println(tree.Root.Search(-400).GetSuccessor().key)
}
