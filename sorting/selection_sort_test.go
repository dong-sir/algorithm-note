/*
Copyright 2023 dong-sir.

selectionSort 选择排序

*/

package sorting

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		desc  string
		slice []int
		want  []int
	}{
		{
			desc:  "NoDuplicates",
			slice: []int{5, 8, 6, 3, 9, 2, 1, 7},
			want:  []int{1, 2, 3, 5, 6, 7, 8, 9},
		},
		{
			desc:  "MultipleDuplicates",
			slice: []int{5, 8, 6, 3, 9, 2, 1, 1, 7, 6, 6},
			want:  []int{1, 1, 2, 3, 5, 6, 6, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := SelectionSort(tc.slice)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("SelectionSort(%v) = %v, want: %v", tc.slice, got, tc.want)
			}
		})
	}

}
