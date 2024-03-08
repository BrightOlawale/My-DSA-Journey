package march

func maxFrequencyElements(nums []int) int {
    frequencyMap := make(map[int]int)
    maxFrequency := 0
    for _, element := range nums{
        frequencyMap[element]++
        if frequencyMap[element] > maxFrequency{
            maxFrequency = frequencyMap[element]
        }
    }

    count := 0
    for _, frequency := range frequencyMap{
        if frequency == maxFrequency{
            count += frequency
        }
    }

    return count
}