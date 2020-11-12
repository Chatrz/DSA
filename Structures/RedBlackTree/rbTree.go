package main

////////////////////////////////
//    Author : Armin Goodarzi //
//    GitHub :                //
//      github.com/Armingodiz //
//                            //
////////////////////////////////

import (
	"fmt"
	"strconv"
)

type Color int

const (
	Black Color = iota
	Red
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	color  Color
	key    int
}

type Tree struct {
	Root *Node
}

func CreateNode(key int) *Node {
	return &Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		color:  Red,
		key:    key,
	}
}

func CreateTree() *Tree {
	tree := &Tree{
		Root: nil,
	}
	return tree
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

func (node *Node) IsLeaf() bool {
	if node.Right == nil && node.Left == nil {
		return true
	}
	return false
}

func (node *Node) IsRightChild() bool {
	if node.key > node.Parent.key {
		return true
	}
	return false
}

func (node *Node) getUncle() *Node {
	if node.Parent.IsRightChild() {
		return node.Parent.Parent.Left
	}
	return node.Parent.Parent.Right
}

func (node *Node) getGrandParent() *Node {
	return node.Parent.Parent
}

func (tree *Tree) Rotate_right(node *Node) {
	// setting targetNode.right as node.left
	// node as and targetNode.right
	targetNode := node.Left
	node.Left = targetNode.Right
	node.Left.Parent = node
	targetNode.Right = node
	setParrentsAfterRotation(tree, node, targetNode)
}

func (tree *Tree) Rotate_left(node *Node) {
	// setting targetNode.Left as node.Right
	// and setting node as targetNode.Left
	targetNode := node.Right
	node.Right = targetNode.Left
	node.Right.Parent = targetNode
	targetNode.Left = node
	setParrentsAfterRotation(tree, node, targetNode)
}

func setParrentsAfterRotation(tree *Tree, node, targetNode *Node) {
	// setting targetNode.parent
	if node == tree.Root { //node was the root of the tree
		tree.Root = targetNode
		targetNode.Parent = nil
	} else {
		targetNode.Parent = node.Parent
		if node.IsRightChild() {
			node.Parent.Right = targetNode
		} else {
			node.Parent.Left = targetNode
		}
	}
	// setting targetNode as node's parent
	node.Parent = targetNode
}

func main() {
	tree := CreateTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(7)
	fmt.Println("before right rotation :")
	tree.DisplayTree()
	fmt.Println("After right rotation :")
	tree.Rotate_right(tree.Root)
	tree.DisplayTree()
  fmt.Println("After left rotation :")
  tree.Rotate_left(tree.Root)
  tree.DisplayTree()
}

////////////////////////////////////////////////////////////////// TODOs :
//TODO implement insertion

func (tree *Tree) Insert(data int) {
	newNode := CreateNode(data)
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

//TODO: implement deletion

/////////////////////////////////////////////////////////////////// displaying tree :
type TreePicture struct {
	pic string
}

func GetTreePic(res *TreePicture, padding string, pointer string, node *Node) {
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

		GetTreePic(res, paddingForBoth, pointerForLeft, node.Left)
		GetTreePic(res, paddingForBoth, pointerForRight, node.Right)
	}
}

func (tree *Tree) DisplayTree() {
	pic := &TreePicture{pic: ""}
	GetTreePic(pic, "", "", tree.Root)
	fmt.Println(pic.pic)
}
