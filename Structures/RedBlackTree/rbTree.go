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

const(
  Black Color = iota
  Red
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
  color Color
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

func (tree *Tree)DisplayTree()  {
	pic := &TreePicture{pic: ""}
	GetTreePic(pic, "", "", tree.Root)
	fmt.Println(pic.pic)
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
