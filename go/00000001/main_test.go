package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplesSum(t *testing.T) {
	var limit = 10
	var want = 23
	got := multiplesSum(limit)
	require.Equal(t, want, got)
}
