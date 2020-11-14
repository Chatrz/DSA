package main

////////////////////////////////
//    Author : Armin Goodarzi //
//    GitHub :                //
//      github.com/Armingodiz //
//                            //
////////////////////////////////

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

///////////////////////////////////////// types :

type ColorNode int

const (
	Black ColorNode = iota
	Red
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Color  ColorNode
	key    int
}

type Tree struct {
	Root  *Node
	dummy *Node
}

type TreePicture struct {
	pic string
}

//////////////////////////////////////////////

/////////////////////////////////////// create funcs :

func (tree *Tree) CreateNode(key int) *Node {
	return &Node{
		Parent: nil,
		Left:   tree.dummy,
		Right:  tree.dummy,
		Color:  Red,
		key:    key,
	}
}

func CreateTree() *Tree {
	tree := &Tree{
		Root:  nil,
		dummy: nil,
	}
	tree.dummy = tree.CreateNode(-1234)
	tree.dummy.Color = Black
	return tree
}

/////////////////////////////////////////////

////////////////////////////////////////////////////////// getter funcs :

func (node *Node) getUncle() *Node {
	if node.Parent.IsRightChild() {
		return node.Parent.Parent.Left
	}
	return node.Parent.Parent.Right
}

func (node *Node) getGrandParent() *Node {
	return node.Parent.Parent
}

/////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////// other funcs :

func (node *Node) Search(tree *Tree, key int) *Node {
	if node == tree.dummy {
		fmt.Println("KEY DOES NOT EXIST !")
		return nil
	}
	if node.key == key {
		return node
	}
	if key > node.key {
		return node.Right.Search(tree, key)
	} else {
		return node.Left.Search(tree, key)
	}
}

func (node *Node) IsRightChild() bool {
	if node.key > node.Parent.key {
		return true
	}
	return false
}

/////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////// Rotations funcs :

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

////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////// insertion :

func (tree *Tree) Insert(data int) {
	newNode := tree.CreateNode(data)
	if tree.Root == nil {
		tree.Root = newNode
		tree.Root.Color = Black
	} else {
		tmp := tree.Root
		for {
			holder := tmp
			if data > tmp.key {
				tmp = tmp.Right
				if tmp == tree.dummy {
					holder.Right = newNode
					newNode.Parent = holder
					break
				}
			} else {
				tmp = tmp.Left
				if tmp == tree.dummy {
					holder.Left = newNode
					newNode.Parent = holder
					break
				}
			}
		}
		tree.fixRbViolations(newNode)
	}
}

func (tree *Tree) fixRbViolations(node *Node) {
	for node != tree.Root && node.Color == Red && node.Parent.Color == Red {
		if !node.Parent.IsRightChild() { // for LL and LR cases :
			if node.getUncle() != nil && node.getUncle().Color == Red { // case L.1 :
				node.Parent.Color = Black
				node.getUncle().Color = Black
				node.getGrandParent().Color = Red
				node = node.getGrandParent()
			} else { // case L.2 & L.3
				if node.IsRightChild() { // case L.2 = LR
					node = node.Parent
					tree.Rotate_left(node)
				} //case L.3 = LL :
				node.Parent.Color = Black
				node.getGrandParent().Color = Red
				tree.Rotate_right(node.getGrandParent())
			}
		} else { // for RR and RL cases :
			if node.getUncle() != nil && node.getUncle().Color == Red { // case R.1 :
				node.Parent.Color = Black
				node.getUncle().Color = Black
				node.getGrandParent().Color = Red
				node = node.getGrandParent()
			} else { // case R.2 & R.3
				if !node.IsRightChild() { // case R.2 = RL
					node = node.Parent
					tree.Rotate_right(node)
				} //case R.3 = RR :
				node.Parent.Color = Black
				node.getGrandParent().Color = Red
				tree.Rotate_left(node.getGrandParent())
			}
		}
	}
	tree.Root.Color = Black
}

/////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////// TODOs :
//TODO: implement deletion

/////////////////////////////////////////////////////////////////// displaying tree :

func GetTreePic(tree *Tree, res *TreePicture, padding string, pointer string, node *Node) {
	if node != tree.dummy {
		res.pic = res.pic + padding
		res.pic = res.pic + pointer
		key := strconv.Itoa(node.key)
		if node.Color == Red {
			red := color.New(color.FgRed).SprintFunc()
			key = red(key)
		}
		res.pic = res.pic + key
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

		GetTreePic(tree, res, paddingForBoth, pointerForLeft, node.Left)
		GetTreePic(tree, res, paddingForBoth, pointerForRight, node.Right)
	}
}

func (tree *Tree) DisplayTree() {
	pic := &TreePicture{pic: ""}
	GetTreePic(tree, pic, "", "", tree.Root)
	fmt.Println(pic.pic)
}

/////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////// testing implemention :
func main() {
	tree := CreateTree()
	tree.Insert(10)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(18)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(7)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(15)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(16)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(30)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(25)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(40)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(60)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(2)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(1)
	tree.DisplayTree()
	fmt.Println("######################################")
	tree.Insert(70)
	tree.DisplayTree()
	fmt.Println("######################################")
}
