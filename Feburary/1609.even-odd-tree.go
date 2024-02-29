package feburary

/*
 * @lc app=leetcode id=1609 lang=golang
 *
 * [1609] Even Odd Tree
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
 func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		prev := -1

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if (level%2 == 0 && node.Val%2 == 0) || (level%2 != 0 && node.Val%2 != 0) {
				return false // Condition violated: Even level with even value or odd level with odd value
			}

			if prev != -1 && ((level%2 == 0 && prev >= node.Val) || (level%2 != 0 && prev <= node.Val)) {
				return false // Condition violated: Values not strictly increasing (even level) or decreasing (odd level)
			}

			// Update the previous value for the next iteration
			prev = node.Val

			// Enqueue the children of the current node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Move to the next level
		level++
	}

	return true
}
// @lc code=end

