package huffman

import (
	"testing"
)

// TestHuffmanCompression tests the HuffmanCompression function.
// func TestHuffmanCompression(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		value    string
// 		expected string
// 	}{
// 		//{"empty string", "", ""},
// 		{"Test from GeeksForGeeks", "dcffebffedeadadefccfdefeaefbfcfedfbdfdebcdcfcfedfffbfcfaefefffcbeffefffdfbffbfffdffceefffcbafdfcffff", ""},
// 		//{"Text from post", "Huffman coding is a data compression algorithm.", ""}, // Uncomment and update when the function is implemented
// 		// {"multiple characters", "abc", "some_expected_result"}, // Uncomment and update when the function is implemented
// 		// Add more test cases as necessary
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := HuffmanCompression(tt.value); got != tt.expected {
// 				t.Errorf("HuffmanCompression(%v) = %v, want %v", tt.value, got, tt.expected)
// 			}
// 		})
// 	}
// }

func TestHuffmanEncodeWithSpecificCodesGeeksForGeeks(t *testing.T) {
	sampleString := "dcffebffedeadadefccfdefeaefbfcfedfbdfdebcdcfcfedfffbfcfaefefffcbeffefffdfbffbfffdffceefffcbafdfcffff"

	expectedCodes := map[string]string{
		"f": "0",
		"c": "100",
		"d": "101",
		"a": "1100",
		"b": "1101",
		"e": "111",
	}

	encoded := GetCodes(sampleString)

	for char, expectedCode := range expectedCodes {
		if code, exists := encoded[char]; !exists || code != expectedCode {
			t.Errorf("Code for '%s' is incorrect, got: %s, want: %s", char, code, expectedCode)
		}
	}
}
