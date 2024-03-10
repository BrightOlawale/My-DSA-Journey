package feburary

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start

func rangeBitwiseAnd(left, right int) int {
	shift := 0
	for left < right {
		left >>= 1 // Right shift left
		right >>= 1 // Right shift right
		shift++
	}
	return left << shift
}


func rangeBitwiseAndII(left, right int) int {
    for left < right {
        right &= (right - 1) // Clear the rightmost set bit of right
    }
    return right
}


// @lc code=end

