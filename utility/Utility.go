package utility

import (
	"time"
	"math/rand"
)

func GetRandBinaryArray(length int, trueCount int) []bool {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	array := make([]bool, length)

	c := 0
	for c < trueCount {
		index := r.Intn(length)
		if !array[index] {
			array[index] = true
			c++
		}
	}

	return array
}
