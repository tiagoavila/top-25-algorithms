package quickselect

import "testing"

// TestQuickSelect tests the QuickSelect function with different scenarios
func TestQuickSelect(t *testing.T) {
	tests := []struct {
		arr       []int
		kPosition int
		expect    int
	}{
		{arr: []int{7, 10, 4, 3, 20, 15}, kPosition: 3, expect: 7},
		{arr: []int{7, 10, 4, 3, 20, 15}, kPosition: 4, expect: 10},
	}

	for _, test := range tests {
		result := QuickSelect(test.arr, 0, len(test.arr)-1, test.kPosition-1)
		if result != test.expect {
			t.Errorf("QuickSelect(%v, 0, %d, %d) = %d; want %d", test.arr, len(test.arr)-1, test.kPosition, result, test.expect)
		}
	}
}
