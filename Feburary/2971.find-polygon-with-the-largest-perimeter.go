package feburary

import "sort"

/*
 * @lc app=leetcode id=2971 lang=golang
 *
 * [2971] Find Polygon With the Largest Perimeter
 */

// @lc code=start
func largestPerimeter(nums []int) int64 {
    // Sort the array in ascending order
    sort.Ints(nums)

    // Initialize variables
    var previousElementsSum int64
    var ans int64 = -1

    // Iterate through the array
    for i := 0; i < len(nums); i++ {
        // Check if enough elements for a triangle and current element is smaller than sum of previous two
        if i >= 2 && int64(nums[i]) < previousElementsSum {
            // Calculate potential perimeter and update if larger
            perimeter := int64(nums[i]) + previousElementsSum
            if perimeter > ans {
                ans = perimeter
            }
        }
        // Add current element to previous elements sum
        previousElementsSum += int64(nums[i])
    }

    return ans
}

// @lc code=end

