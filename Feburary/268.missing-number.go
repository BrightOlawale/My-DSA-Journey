package feburary

import "sort"

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
    sort.Ints(nums)
    for n:=len(nums); n >= 0; n-- {
        if n-1 >= 0 && n == nums[n-1]{
            continue
        } else{
            return n
        }
    }
    return 0
}

func missingNumber(nums []int) int {
    missing := len(nums) 
    
    for i, num := range nums {
        missing ^= i ^ num
    }
    
    return missing
}
// @lc code=end

