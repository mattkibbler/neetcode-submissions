func isAnagram(s string, t string) bool {
	if s == t {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, c := range s {
		sMap[c] += 1
	}
	for _, c := range t {
		tMap[c] += 1
	}

	for c, _ := range sMap {
		if sMap[c] != tMap[c] {
			return false
		}
	}


	return true
}
