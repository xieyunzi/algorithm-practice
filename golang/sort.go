package golang

// interface from golang "sort" package
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func InsertionSort(data Interface) Interface {
	len := data.Len()

	for i := 0; i < len; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}

	return data
}

func BubbleSort(data Interface) Interface {
	len := data.Len()

	for i := len - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}

	return data
}

func SelectionSort(data Interface) Interface {
	len := data.Len()

	var larger int
	for i := len - 1; i > 0; i-- {
		larger = i
		for j := 0; j < i; j++ {
			if data.Less(larger, j) {
				larger = j
			}
		}
		data.Swap(larger, i)
	}

	return data
}

func ShellSort(data Interface) Interface {
	len := data.Len()

	for gap := len / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len; i++ {
			for j := i; j >= gap && data.Less(j, j-gap); j -= gap {
				data.Swap(j, j-gap)
			}
		}
	}

	return data
}

func partition(data Interface, low, high int) int {
	pivot := low

	for i := low; i < high; i++ {
		if data.Less(i, high) {
			data.Swap(pivot, i)
			pivot++
		}
	}
	data.Swap(pivot, high)
	return pivot
}

func QuickSort(data Interface, lowhigh ...int) Interface {
	var low, high int
	if len(lowhigh) == 2 {
		low, high = lowhigh[0], lowhigh[1]
	} else {
		low, high = 0, data.Len()-1
	}

	if low < high {
		pi := partition(data, low, high)
		QuickSort(data, low, pi-1)
		QuickSort(data, pi+1, high)
	}

	return data
}

func reverse(data Interface, start, end int) {
	end--
	for i := 0; start+i < end-i; i++ {
		data.Swap(start+i, end-i)
	}
}

func rotate(data Interface, a, p, b int) {
	if a < p && p < b {
		reverse(data, a, p)
		reverse(data, p, b)
		reverse(data, a, b)
	}
}

func merge(data Interface, l, m, r int) {
	var step int
	for i, j := l, m; i < j && j < r; {
		for i < j && data.Less(i, j) {
			i++
		}
		step = 0
		for j < r && data.Less(j, i) {
			j++
			step++
		}
		rotate(data, i, j-step, j)
		i = i + step
	}
}

// func lower(data Interface, from, to, pivot int) int {
// 	for from < to {
// 		mid := from + (to-from)/2
// 		if data.Less(mid, pivot) {
// 			from = mid + 1
// 		} else {
// 			to = mid
// 		}
// 	}
// 	return from
// }

// func upper(data Interface, from, to, pivot int) int {
// 	for from < to {
// 		mid := from + (to-from)/2
// 		if data.Less(pivot, mid) {
// 			to = mid
// 		} else {
// 			from = mid + 1
// 		}
// 	}
// 	return from
// }

// func merge(data Interface, l, pivot, r int) {
// 	if l < pivot && pivot < r && data.Less(pivot, pivot-1) {
// 		if r-l == 2 {
// 			reverse(data, l, r)
// 		} else {
// 			var firstCut, secondCut int
// 			if pivot-l > r-pivot {
// 				firstCut = l + (pivot-l)/2
// 				secondCut = lower(data, pivot, r, firstCut)
// 			} else {
// 				secondCut = pivot + (r-pivot)/2
// 				firstCut = upper(data, l, pivot, secondCut)
// 			}
// 			rotate(data, firstCut, pivot, secondCut)
// 			middle := firstCut + secondCut - pivot
// 			merge(data, middle, secondCut, r)
// 			merge(data, l, firstCut, middle)
// 		}
// 	}
// }

func InplaceMergeSort(data Interface, leftright ...int) Interface {
	var l, r int
	if len(leftright) == 2 {
		l, r = leftright[0], leftright[1]
	} else {
		l, r = 0, data.Len()
	}

	if r-l > 1 {
		m := l + (r-l)/2

		InplaceMergeSort(data, l, m)
		InplaceMergeSort(data, m, r)

		merge(data, l, m, r)
	}

	return data
}

func HeapSort(data Interface) Interface {
	return data
}
