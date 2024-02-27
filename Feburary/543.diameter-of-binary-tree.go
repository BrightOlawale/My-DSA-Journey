package feburary

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    diameter:= 0
    depth(root, &diameter)
    return diameter
}

func depth(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }

    leftDepth := depth(node.Left, diameter)
    rightDepth := depth(node.Right, diameter)

    *diameter = max(*diameter, leftDepth+rightDepth)

    return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

