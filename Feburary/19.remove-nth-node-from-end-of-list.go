/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 dummy := &ListNode{0, head}
 fast := dummy
 slow := dummy

 // Move fast pointer n steps ahead
 for i := 0; i <= n; i++ {
	 fast = fast.Next
 }

 // Move both fast and slow pointers until fast pointer reaches the end
 for fast != nil {
	 fast = fast.Next
	 slow = slow.Next
 }

 // Remove the nth node by adjusting pointers
 slow.Next = slow.Next.Next

 return dummy.Next
// @lc code=end

