package test

import (
	"gofaker/common"
)

type testSource struct {
	constant int64
}

func (ts *testSource) Int63() int64 {
	return ts.constant
}

func (ts *testSource) Seed(seed int64) {
	ts.constant = seed
}

func newTestSource(constant int32) *testSource {
	ts := testSource{}
	ts.Seed(int64(constant) << 32)
	return &ts
}

func ConstantRand(constant int32) *common.Rand {
	return common.New(newTestSource(constant))
}
