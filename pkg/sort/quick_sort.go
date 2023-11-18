package sort

import (
	"cmp"
	"math/rand"
)

// QuickSort sort slice of elements which is possible to compare.
func QuickSort[S ~[]E, E cmp.Ordered](arr S) {
	QuickSortFunc(arr, func(a, b E) int {
		switch {
		case a > b:
			return 1
		case a < b:
			return -1
		default:
			return 0
		}
	})
}

// QuickSortFunc sort slice of elements, according to compare func.
func QuickSortFunc[S ~[]E, E any](arr S, cmp func(a, b E) int) {
	if len(arr) <= 1 {
		return
	}

	pivot := chooseRandPivot(arr)
	less, greater := partition(arr, pivot, cmp)
	QuickSortFunc(less, cmp)
	QuickSortFunc(greater, cmp)
}

func chooseRandPivot[S ~[]E, E any](arr S) E {
	return arr[rand.Intn(len(arr))]
}

func partition[S ~[]E, E any](arr S, pivot E, cmp func(a, b E) int) (less, greater S) {
	eqStartIdx := 0
	for i := range arr {
		if cmp(arr[i], pivot) < 0 {
			arr[i], arr[eqStartIdx] = arr[eqStartIdx], arr[i]
			eqStartIdx++
		}
	}

	eqEndIdx := eqStartIdx
	for i := eqEndIdx; i < len(arr); i++ {
		if cmp(arr[i], pivot) == 0 {
			arr[i], arr[eqEndIdx] = arr[eqEndIdx], arr[i]
			eqEndIdx++
		}
	}

	return arr[:eqStartIdx], arr[eqEndIdx:]
}
