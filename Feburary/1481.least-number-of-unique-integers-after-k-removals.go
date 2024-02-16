/*
 * @lc app=leetcode id=1481 lang=golang
 *
 * [1481] Least Number of Unique Integers after K Removals
 */

// @lc code=start
func findLeastNumOfUniqueInts(arr []int, k int) int {
    // Count the frequency of each element
    freqMap := make(map[int]int)
    for _, elem := range arr {
        freqMap[elem]++
    }

    // Extract frequencies into a slice
    freqs := make([]int, 0, len(freqMap))
    for _, freq := range freqMap {
        freqs = append(freqs, freq)
    }

    // Sort frequencies in ascending order
    sort.Ints(freqs)

    // Remove the lowest frequency integers until k elements are removed or no more unique integers left
    removed := 0
    for _, freq := range freqs {
        if k >= freq {
            // If there are enough removals available, remove this frequency of integers
            k -= freq
            removed++
        } else {
            // If not enough removals available, stop iteration
            break
        }
    }

    // Return the number of remaining unique integers
    return len(freqs) - removed
}
// @lc code=end

