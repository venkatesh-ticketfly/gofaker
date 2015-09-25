package common

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

func NewRand() *rand.Rand {
	return rand.New(&lockedSource{src: rand.NewSource(1)})
}

func Choose(strSlice []string, rand *rand.Rand) string {
	return strSlice[rand.Intn(len(strSlice))]
}
