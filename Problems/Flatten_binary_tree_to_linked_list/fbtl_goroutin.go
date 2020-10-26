/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */
//this is my solution for https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package fbtl

import "sync"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatter(root *TreeNode, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if root != nil && !(root.Left == nil && root.Right == nil) {
		var cwg sync.WaitGroup
		cwg.Add(2)
		go flatter(root.Right, &cwg)
		go flatter(root.Left, &cwg)
		cwg.Wait()
		if tmp := root.Left; tmp != nil {
			for ; tmp.Right != nil; tmp = tmp.Right {
			}
			tmp.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
	}
}

func flatten(root *TreeNode) {
	flatter(root, nil)
}
