/*
 * +------------------------------------+
 * | Author:        hadi abbasi	        |
 * |                                    |
 * | Link:          github.com/hawwwdi  |
 * +------------------------------------+
 */
//this is my solution(java version) for https://leetcode.com/problems/flatten-binary-tree-to-linked-list/


/* 
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
} */

class Solution {
    public void flatten(TreeNode root) {
        if (root!= null && !(root.left == null && root.right == null)) {
            flatten(root.right);
            flatten(root.left);
            TreeNode tmp = root.left;
            if (tmp != null) {
                for(;tmp.right != null; tmp = tmp.right);
                tmp.right = root.right;
                root.right = root.left;
                root.left = null;
            }
        }
    }
}