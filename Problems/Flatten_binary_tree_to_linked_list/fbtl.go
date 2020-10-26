/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */

package Flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatter(root *TreeNode) {
	if root != nil && !(root.Left == nil && root.Right == nil) {
		flatter(root.Right)
		flatter(root.Left)
		if tmp := root.Left; tmp != nil {
			for ; tmp.Right != nil; tmp = tmp.Right {
			}
			tmp.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
	}
}
