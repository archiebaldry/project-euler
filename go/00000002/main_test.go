package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumEvenFibSequence(t *testing.T) {
	limit := 10
	want := 10
	got := SumEvenFibSequence(limit)
	require.Equal(t, want, got)
}
