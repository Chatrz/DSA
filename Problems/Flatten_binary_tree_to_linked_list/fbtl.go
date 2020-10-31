/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */
//this is my solution for https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package Flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil && !(root.Left == nil && root.Right == nil) {
		flatten(root.Right)
		flatten(root.Left)
		if tmp := root.Left; tmp != nil {
			for ; tmp.Right != nil; tmp = tmp.Right {
			}
			tmp.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
	}
}
