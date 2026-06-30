func groupAnagrams(strs []string) [][]string {
    // create map using "dictionary" array as key
    groups := make(map[[26]int][]string)

    for _, str := range strs {
        // create "dictionary" array - one element per letter of alphabet
        counts := [26]int{}
        for _, char := range str {
            counts[char-'a']++
        }
        // add each word to the group using the count of their characters as the key
        groups[counts] = append(groups[counts], str)
    }

    result := make([][]string, 0, len(groups))

    for _, group := range groups {
        result = append(result, group)
    }

    return result

}