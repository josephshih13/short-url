package base62

import (
	"strings"
)

const (
	alphabet = "i6jXnuZfTtg0z2SwQvm74haVdCKpE1lsWAYbGN3qrIPF8xUyHke9oRDc5MLBJO"
	length   = len(alphabet)
)

func Encode(number int) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(5)

	for idx := 0; idx < 5; number, idx = number/length, idx+1 {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

// func Decode(encoded string) (uint64, error) {
// 	var number uint64

// 	for i, symbol := range encoded {
// 		alphabeticPosition := strings.IndexRune(alphabet, symbol)

// 		if alphabeticPosition == -1 {
// 			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
// 		}
// 		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
// 	}

// 	return number, nil
// }
