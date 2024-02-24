package random

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"
)

func Random(n int) int {
	var seed int64
	err := binary.Read(crand.Reader, binary.BigEndian, &seed)
	if err != nil {
		panic(err)
	}
	rng := rand.New(rand.NewSource(seed))
	return rng.Intn(n)
}

func Duration(n int) time.Duration {
	return time.Duration(Random(n)) * time.Millisecond
}
