package feburary

import "strings"

/*
 * @lc app=leetcode id=2864 lang=golang
 *
 * [2864] Maximum Odd Binary Number
 */

// @lc code=start
func maximumOddBinaryNumber(s string) string {
	return strings.Repeat("1", strings.Count(s, "1")-1) + strings.Repeat("0", strings.Count(s, "0")) + "1"
}
// @lc code=end

