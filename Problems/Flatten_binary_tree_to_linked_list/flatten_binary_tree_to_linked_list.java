/*
the answer of https://leetcode.com/problems/flatten-binary-tree-to-linked
which i guess turned to be really good 
*/

import java.util.*;
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    // Stack<TreeNode> stack=new Stack<TreeNode>();
    public void flatten(TreeNode root) {
        if(root==null)return;
        flattenByPointer(root);
    }
    public TreeNode flattenByPointer(TreeNode root){
        if (root.right==null&&root.left==null)return root;
        else if(root.right==null){
            root.right=root.left;
            root.left=null;
            TreeNode rightEnd=flattenByPointer(root.right);
            return rightEnd;
        }else if(root.left==null){
            TreeNode rightEnd=flattenByPointer(root.right);
            return rightEnd;
        }else{
            TreeNode leftEnd=flattenByPointer(root.left);
            TreeNode rightEnd=flattenByPointer(root.right);
            TreeNode rightTree=root.right;
            root.right=root.left;
            leftEnd.right=rightTree;
            root.left=null;
            return rightEnd;
        }
    }
}