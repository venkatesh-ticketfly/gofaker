package common

import "strconv"

func Numerify(pattern string, numerifyRune rune, rand *Rand) string {
	result := []byte{}
	for _, char := range pattern {
		if char == numerifyRune {
			result = strconv.AppendInt(result, int64(rand.Int31n(10)), 10)
		} else {
			result = append(result, byte(char))
		}
	}

	return string(result)
}
