package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestMain(m *testing.M) {
	main()
}
