package prob91

var cache map[int]int

func NumDecodings(s string) int {
	cache = make(map[int]int, len(s))
	return decode(0, s)
}

func decode(start int, s string) int {
	strLen := len(s)
	if start == strLen {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	if val, exists := cache[start]; exists {
		return val
	}

	combinations := decode(start+1, s)
	// filter out edge cases
	if start+1 < strLen && (s[start] == '1' || s[start] == '2' && s[start+1] < '7') {
		combinations += decode(start+2, s)
	}

	cache[start] = combinations
	return cache[start]
}
