package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/samber/lo"
	"testing"
)

func Test_Filter(t *testing.T) {
	even := lo.Filter([]int{1, 2, 3, 4, 5, 6}, func(x int, index int) bool {
		return x%2 == 0
	})

	want := []int{2, 4, 6}
	if diff := cmp.Diff(want, even); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func Test_Filter_Normal(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	var even []int

	for _, x := range input {
		if x%2 == 0 {
			even = append(even, x)
		}
	}

	want := []int{2, 4, 6}
	if diff := cmp.Diff(want, even); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func Test_Map(t *testing.T) {
	doubled := lo.Map([]int{1, 2, 3, 4, 5}, func(item int, index int) int {
		return item * 2
	})

	want := []int{2, 4, 6, 8, 10}
	if diff := cmp.Diff(want, doubled); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func Test_Map_Normal(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	doubled := make([]int, len(input))

	for i, item := range []int{1, 2, 3, 4, 5} {
		doubled[i] = item * 2
	}

	want := []int{2, 4, 6, 8, 10}
	if diff := cmp.Diff(want, doubled); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}

func Test_UniqMpa(t *testing.T) {
	uniq := lo.Uniq([]int{1, 1, 2, 2, 3, 4, 5, 5})

	want := []int{1, 2, 3, 4, 5}
	if diff := cmp.Diff(want, uniq); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}
