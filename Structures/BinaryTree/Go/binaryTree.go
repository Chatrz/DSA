package binaryTree

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
	/*fmt.Print("SEARCHING FOR : ")
	fmt.Println(key)
	fmt.Print("ON KEY ")
	fmt.Println(node.key)*/
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
func (tree *Tree) DeleteUseKey(key int) {
	node := tree.Root.Search(key)
	node.DeleteNode()
}
func (node *Node) DeleteNode() bool {
	if node == nil {
		return false
	} else {
		if node.IsLeaf() { //node has no children
			if node.IsRightChild() {
				node.Parent.Left = nil

			} else {
				node.Parent.Right = nil
			}
			node.Parent = nil
		} else if node.Right != nil && node.Left == nil { //node has one children at right
			node.Right.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Right
			} else {
				node.Parent.Left = node.Right
			}
		} else if node.Left != nil && node.Right == nil { //node has one children at left
			node.Left.Parent = node.Parent
			if node.IsRightChild() {
				node.Parent.Right = node.Left
			} else {
				node.Parent.Left = node.Left
			}
		} else { //node has two children
			sucNod := node.GetSuccessor()
			fmt.Println(sucNod.key)
			holder := sucNod.key
			sucNod.key = node.key
			node.key = holder
			fmt.Println(sucNod.key)
			sucNod.DeleteNode()
		}
		return true
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


func main() {
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
	pic := &TreePicture{pic: ""}
	DisplayTree(pic, "", "", tree.Root)
	fmt.Println(pic.pic)
}
