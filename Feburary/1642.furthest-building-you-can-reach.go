package feburary

/*
 * @lc app=leetcode id=1642 lang=golang
 *
 * [1642] Furthest Building You Can Reach
 */

// @lc code=start
import (
	"container/heap"
)

// MinHeap implements the heap.Interface for a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
    pq := &MinHeap{}
    heap.Init(pq)
    
    for i := 1; i < len(heights); i++ {
        diff := heights[i] - heights[i-1]
        if diff > 0 {
            heap.Push(pq, diff)
            if pq.Len() > ladders {
                bricks -= heap.Pop(pq).(int)
                if bricks < 0 {
                    return i - 1
                }
            }
        }
    }
    
    return len(heights) - 1
}

// @lc code=end

