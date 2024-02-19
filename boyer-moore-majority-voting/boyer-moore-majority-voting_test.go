package boyermooremajorityvoting

import (
	"testing"
)

func TestFindMajorityElement(t *testing.T) {
	testCases := []struct {
		name     string
		array    []int
		expected int
	}{
		{"MajorityElementExists", []int{3, 3, 4, 2, 4, 4, 2, 4, 4}, 4},
		{"NoMajorityElement", []int{3, 3, 4, 2, 4, 4, 2, 4}, -1},
		{"AllSameElements", []int{1, 1, 1, 1}, 1},
		{"SingleElement", []int{1}, 1},
		{"EmptyArray", []int{}, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMajorityElement(tc.array)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
