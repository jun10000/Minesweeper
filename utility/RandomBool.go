package utility

import (
	"math/rand"
    "time"
)

type RandomBool struct {
    src       rand.Source
    cache     int64
    remaining int
}

func NewRandBool() *RandomBool {
    return &RandomBool{ src: rand.NewSource(time.Now().UnixNano()) }
}

func (b *RandomBool) Get() bool {
    if b.remaining == 0 {
        b.cache, b.remaining = b.src.Int63(), 63
    }

    result := b.cache&0x01 == 1
    b.cache >>= 1
    b.remaining--

    return result
}
