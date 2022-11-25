package utility

import (
	"time"
	"math/rand"
)

type Position struct {
	X int
	Y int
}

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

func GetHoursToSeconds(d time.Duration) (int, int, int) {
	s := int(d.Seconds())

	h := s / 3600
	s -= h * 3600

	m := s / 60
	s -= m * 60
	
	return h, m, s
}
