package euclideanfindgcd

import "testing"

func TestFindGCD(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Name of the test case
		a, b     int    // Input values
		expected int    // Expected result
	}{
		{"GCD of 270 and 192", 270, 192, 6},
		{"GCD of 10 and 15", 10, 15, 5},
		{"GCD of 10 and 5", 10, 5, 5},
		{"GCD of 15 and 20", 15, 20, 5},
		{"GCD of 3 and 7", 3, 7, 1},
		{"GCD of 0 and 5", 0, 5, 5},
		{"GCD of 5 and 0", 5, 0, 5},
		{"GCD of 0 and 0", 0, 0, 0}, // Depending on the mathematical definition you follow, this might be undefined.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			result := FindGCD(tt.a, tt.b)

			// Check if the result is as expected
			if result != tt.expected {
				t.Errorf("FindGCD(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
