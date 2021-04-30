package sdlkit

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

// RNG returns a new rand.Rand with the current unix time as source.
func RNG() *rand.Rand {
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return rng
}
