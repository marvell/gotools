package gotools

import "testing"

func TestValueWithDefault(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		val      *int
		def      int
		expected int
	}{
		{"Nil value, return default", nil, 10, 10},
		{"Non-nil value, return value", ptr(5), 10, 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ValueWithDefault(tc.val, tc.def)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}

// Helper function to create a pointer
func ptr[T any](value T) *T {
	return &value
}
