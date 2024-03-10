package march

func intersection(nums1 []int, nums2 []int) []int {
    // Create a map to store the unique elements of nums1
    numSet := make(map[int]bool)
    for _, num := range nums1 {
        numSet[num] = true
    }
    
    // Iterate through nums2 and check for intersection with nums1
    var result []int
    for _, num := range nums2 {
        // If the element is present in nums1, add it to the result and remove it from numSet
        if numSet[num] {
            result = append(result, num)
            delete(numSet, num)
        }
    }
    
    return result
}