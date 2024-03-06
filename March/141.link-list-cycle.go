package march

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    visited := make(map[*ListNode]bool)

    for head != nil {
        if visited[head] {
            return true
        }
        visited[head] = true
        head = head.Next
    }
    
    return false
}