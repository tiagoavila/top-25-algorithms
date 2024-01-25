package kpmpatternsearch

import (
	"reflect"
	"testing"
)

func TestComputeLPSArray(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		want    []int
	}{
		{"Test1", "AABAACAABAA", []int{0, 1, 0, 1, 2, 0, 1, 2, 3, 4, 5}},
		{"Test2", "AAAA", []int{0, 1, 2, 3}},
		{"Test3", "ABCDE", []int{0, 0, 0, 0, 0}},
		{"TestFromLogicFirstChannel", "AABAAAC", []int{0, 1, 0, 1, 2, 2, 0}},
		{"Test1FromAbdulBariChannel", "ABCDABEABF", []int{0, 0, 0, 0, 1, 2, 0, 1, 2, 0}},
		{"Test2FromAbdulBariChannel", "ABCDEABFABC", []int{0, 0, 0, 0, 0, 1, 2, 0, 1, 2, 3}},
		{"Test3FromAbdulBariChannel", "AABCADAABE", []int{0, 1, 0, 0, 1, 0, 1, 2, 3, 0}},
		{"Test4FromAbdulBariChannel", "AAAABAACD", []int{0, 1, 2, 3, 0, 1, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeLPSArray(tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeLPSArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
