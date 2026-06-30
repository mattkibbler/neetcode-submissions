func topKFrequent(nums []int, k int) []int {

    numMap := make(map[int]int)
    for _, num := range nums {
        numMap[num]++
    }

    // use slice to store the number frequencies
    // each index represent the frequency and stores a slice of the numbers appearing
    // that many times
    counts := make([][]int, len(nums)+1)
    for num, freq := range numMap {
        counts[freq] = append(counts[freq], num)
    }

    result := make([]int, 0, k)
    for i := len(counts) - 1; i >= 0; i-- {
        for i2 := 0; i2 < len(counts[i]); i2++ {
            if(len(result) == k) {
                return result
            }
            result = append(result, counts[i][i2])
        }
    }

    return result
}
