package utility

import (
	"time"
	"math/rand"
)

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) Position {
	return Position{ X: x, Y: y }
}

func GetRandBinaryArray(length int, trueCount int) []bool {
	return GetRandBinaryArrayWithSeed(length, trueCount, time.Now().UnixNano())
}

func GetRandBinaryArrayWithSeed(length int, trueCount int, seed int64) []bool {
	s := rand.NewSource(seed)
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

func SplitDuration(d time.Duration) (int, int, int) {
	s := int(d.Seconds())

	h := s / 3600
	s -= h * 3600

	m := s / 60
	s -= m * 60
	
	return h, m, s
}
