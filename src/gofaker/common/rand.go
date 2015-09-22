package common

import (
	"math/rand"
	"time"
)

type Rand struct {
	*rand.Rand
}

func Default() *Rand {
	return New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
}

func New(src rand.Source) *Rand {
	return &Rand{rand.New(src)}
}

func (r *Rand) Choose(strSlice []string) string {
	return strSlice[r.Intn(len(strSlice))]
}
