package feburary

/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
    // Create maps to store the count of people trusting each person
    // and the count of people each person trusts
    trusting := make(map[int]int)
    trusted := make(map[int]int)

    // Iterate through the trust relationships to update the counts
    for _, relation := range trust {
        trusting[relation[0]]++
        trusted[relation[1]]++
    }

    // Check if there is a person who is trusted by all others
    // except themselves and who trusts nobody
    for i := 1; i <= n; i++ {
        if trusted[i] == n-1 && trusting[i] == 0 {
            return i
        }
    }

    // If no such person is found, return -1
    return -1
}
// @lc code=end

