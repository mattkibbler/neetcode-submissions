func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, x := range nums {
        m[x] += 1
    }
    for _, val := range m {
        if val > 1 {
            return true
        }
    }
    return false
}
