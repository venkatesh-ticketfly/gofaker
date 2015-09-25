package test

import "math/rand"

type testSource struct {
	constant int64
}

func (ts *testSource) Int63() int64 {
	return ts.constant
}

func (ts *testSource) Seed(seed int64) {
	panic("Cannot seed a ConstanRand. Inteded to be used by test.")
}

func ConstantRand(constant int32) *rand.Rand {
	return rand.New(&testSource{int64(constant) << 32})
}
