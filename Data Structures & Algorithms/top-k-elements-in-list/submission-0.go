func topKFrequent(nums []int, k int) []int {
    numMap := make(map[int]int)
    for _, num := range nums {
        numMap[num]++
    }

    result := make([]int, 0, k)
    for i := 0; i < k; i++ {
        looped := false
        highestFreq := 0
        highestNum := 0
        for num, freq := range numMap {
            if !looped || freq > highestFreq {
                highestFreq = freq
                highestNum = num
            }
            looped = true
        }
        result = append(result, highestNum)
        delete(numMap, highestNum)
    }
    
    return result
}
