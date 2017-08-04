package golang

// 一个递增数组 [1, 2, 3, 4, 5, 6, 7, 8, 9] 取其中任意一点切分成两段,
// 把左边部分移到最右边成为新数组并命名为 a2,
// 例 a2 = [8, 9, 1, 2, 3, 4, 5, 6, 7],
// 给定一个数字 t ,
// 判断数字 t 是否在数组 a2 中

func pMid(pL, pR int) int {
	return pL + (pR-pL)/2
}

func VariedBinarySearch(arr []int, pL, pR int, target int) int {
	if pR < pL {
		return -1
	}

	mid := pMid(pL, pR)

	// fmt.Println(pL, mid, pR)
	// [5, 6, 1, 2, 3, 4]
	if arr[mid] <= arr[pR] {
		if target >= arr[mid] && target <= arr[pR] {
			return BinarySearch(arr, mid, pR, target)
		}
		return VariedBinarySearch(arr, pL, mid, target)
	}
	// [3, 4, 5, 6, 1, 2]
	if arr[mid] > arr[pL] {
		if target <= arr[mid] && target >= arr[pL] {
			return BinarySearch(arr, pL, mid, target)
		}
		return VariedBinarySearch(arr, mid, pR, target)
	}

	return -1
}

func BinarySearch(arr []int, pL, pR int, target int) int {
	if pR < pL {
		return -1
	}

	mid := pMid(pL, pR)

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		pR = mid - 1
	} else {
		pL = mid + 1
	}

	return BinarySearch(arr, pL, pR, target)
}
