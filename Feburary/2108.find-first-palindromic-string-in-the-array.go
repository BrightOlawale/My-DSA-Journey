package feburary

/*
 * @lc app=leetcode id=2108 lang=golang
 *
 * [2108] Find First Palindromic String in the Array
 */

// @lc code=start
func firstPalindrome(words []string) string {
    for _, word := range words {
        half := len(word) / 2
        isPalindrome := true
        for i := 0; i < half; i++ {
            if word[i] != word[len(word)-1-i] {
                isPalindrome = false
                break
            }
        }
        if isPalindrome {
            return word
        }
    }
    return ""
}
// @lc code=end

