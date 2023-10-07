package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
	found := false
	for i := 0; i < len(*aPtr); i++ {
		if (*aPtr)[i] == elem {
			*aPtr = append((*aPtr)[0:i], (*aPtr)[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		*aPtr = append(*aPtr, elem)
	}
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}
