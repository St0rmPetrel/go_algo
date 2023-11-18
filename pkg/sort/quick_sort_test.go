package sort_test

import (
	"log"
	"testing"

	"github.com/St0rmPetrel/go_algo/pkg/sort"

	"github.com/stretchr/testify/require"
)

var logger = log.Default()

func TestQuickSort(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		slice  []int
		sorted []int
	}{
		"base_case_1": {
			slice:  []int{},
			sorted: []int{},
		},
		"base_case_2": {
			slice:  []int{1},
			sorted: []int{1},
		},
		"sorted_1": {
			slice:  []int{1, 2},
			sorted: []int{1, 2},
		},
		"sorted_eq_1": {
			slice:  []int{2, 2},
			sorted: []int{2, 2},
		},
		"unsorted_1": {
			slice:  []int{2, 1},
			sorted: []int{1, 2},
		},
		"sorted": {
			slice:  []int{1, 2, 3},
			sorted: []int{1, 2, 3},
		},
		"not_sorted": {
			slice:  []int{2, 1, 3},
			sorted: []int{1, 2, 3},
		},
		"random": {
			slice:  []int{32, 23, 4, 5, 33, 2, 3, 55, 23, 34, 6},
			sorted: []int{2, 3, 4, 5, 6, 23, 23, 32, 33, 34, 55},
		},
	}

	for name, testCase := range tests {
		test := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			sort.QuickSort(test.slice)
			require.Equal(t, test.sorted, test.slice)
		})
	}
}
