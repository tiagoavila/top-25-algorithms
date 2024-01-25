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
		{"Test5FromAbdulBariChannel", "ABABD", []int{0, 0, 1, 2, 0}},
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

func TestKMPPatternSearch(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pattern string
		want    []int
	}{
		{"Test1", "ABABCABCABABABD", "ABABD", []int{10}},
		{"Test2", "ABABABABCABABABD", "ABABD", []int{11}},
		{"Test3", "AABAACAADAABAABA", "AABA", []int{0, 9, 12}},
		{"Test4", "ABCABCABCABC", "ABC", []int{0, 3, 6, 9}},
		{"Test5", "GEEKSFORGEEKS", "GEEKS", []int{0, 8}},
		{"Test5", "AAAAAA", "AA", []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindIndexOfSubStrings(tt.text, tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeLPSArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
