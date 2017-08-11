package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Arr []int

func (a Arr) Len() int           { return len(a) }
func (a Arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Arr) Less(i, j int) bool { return a[i] < a[j] }

func TestSort(t *testing.T) {
	raw := Arr{7, 9, 8, 5, 3, 4, 6, 1, 2, 0}
	sorted := Arr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	original := make(Arr, raw.Len())
	copy(original, raw)

	// fmt.Println("original: ", original)
	// fmt.Println("sorted: ", sorted)

	copy(original, raw)
	assert.Equal(t, sorted, InsertionSort(original))
	copy(original, raw)
	assert.Equal(t, sorted, BubbleSort(original))
	copy(original, raw)
	assert.Equal(t, sorted, SelectionSort(original))
	copy(original, raw)
	assert.Equal(t, sorted, ShellSort(original))
	copy(original, raw)
	assert.Equal(t, sorted, QuickSort(original))
	copy(original, raw)
	assert.Equal(t, sorted, InplaceMergeSort(original))
	// copy(original, raw)
	// assert.Equal(t, sorted, HeapSort(original))
}
