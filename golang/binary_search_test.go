package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9}
	target := 5

	assert.Equal(t, 3, BinarySearch(arr, 0, len(arr)-1, target))
}

func TestVariedBinarySearch(t *testing.T) {
	arr := []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{3, 4, 5, 6, 7, 8, 9, 1, 2}

	assert.Equal(t, -1, VariedBinarySearch(arr, 0, len(arr)-1, 10))
	assert.Equal(t, 5, VariedBinarySearch(arr, 0, len(arr)-1, 4))
	assert.Equal(t, 1, VariedBinarySearch(arr, 0, len(arr)-1, 9))
	assert.Equal(t, 5, VariedBinarySearch(arr2, 0, len(arr2)-1, 8))
	assert.Equal(t, 1, VariedBinarySearch(arr2, 0, len(arr2)-1, 4))
}
