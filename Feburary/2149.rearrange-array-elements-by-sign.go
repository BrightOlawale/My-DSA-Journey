package feburary

/*
 * @lc app=leetcode id=2149 lang=golang
 *
 * [2149] Rearrange Array Elements by Sign
 */

// @lc code=start
func rearrangeArray(nums []int) []int {
    positiveArr := make([]int, 0, len(nums)/2)
    negativeArr := make([]int, 0, len(nums)/2)

    for _, elem := range nums {
        if elem > 0 {
            positiveArr = append(positiveArr, elem)
        } else {
            negativeArr = append(negativeArr, elem)
        }
    }
    sign := true
    for idx, _ := range nums {
        if sign {
            if len(positiveArr) > 0 {
                nums[idx] = positiveArr[0]
                positiveArr = positiveArr[1:]
            }
        } else {
            if len(negativeArr) > 0 {
                nums[idx] = negativeArr[0]
                negativeArr = negativeArr[1:]
            }
        }
        sign= !sign
    }
    return nums
}
// @lc code=end

