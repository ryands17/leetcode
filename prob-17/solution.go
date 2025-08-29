package prob17

var digitToLetters = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	var result []string

	revDigits := reverseString(digits)
	for i, d := range revDigits {
		letters := digitToLetters[byte(d)]

		var tempResult []string
		for _, l := range letters {
			if i == 0 {
				result = append(result, string(l))
				continue
			}

			for _, r := range result {
				tempResult = append(tempResult, string(l)+r)
			}
		}

		if i != 0 {
			result = make([]string, len(tempResult))
			copy(result, tempResult)
		}
	}

	return result
}
