package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUniq(t *testing.T) {
	//Arrange
	testTable := []struct {
		numbers  []float32
		expected float32
	}{{
		numbers:  []float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0},
		expected: 2.0,
	},
		{
			numbers:  []float32{4.0, 4.0, 4.0, 0.0, 4.0, 4.0},
			expected: 9.0,
		}, {
			numbers:  []float32{4.0, 4.0, 4.0, 9.0, 4.0, 4.0},
			expected: 9.0,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := FindUniq(testCase.numbers)
		t.Logf("Calling FinfUniq(%v), result %f\n", testCase.numbers, result)
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorect result. Expected: %f, got %f", testCase.expected, result))
		if result != testCase.expected {

		}
	}
	//Assert

}
