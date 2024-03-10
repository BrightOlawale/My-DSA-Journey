package feburary

/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
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
 func findBottomLeftValue(root *TreeNode) int {
    var bottomLeftValue int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size:=len(queue)
        bottomLeftValue = queue[0].Val

        for i:=0; i<size; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            
        }
    }

    return bottomLeftValue
}
// @lc code=end

