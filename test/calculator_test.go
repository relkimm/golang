package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plus(a, b int) int {
	return a + b
}

func TestPlus(t *testing.T) {
	tests := []struct {
		actual   int
		expected int
	}{
		{plus(5, 10), 15},
		{plus(3, 5), 8},
		{plus(2, 3), 5},
		{plus(10, 23), 33},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.actual)
	}
}

func minus(a, b int) int {
	return a - b
}

func TestMinus(t *testing.T) {
	tests := []struct {
		actual   int
		expected int
	}{
		{minus(10, 5), 5},
		{minus(6, 2), 4},
		{minus(8, 5), 3},
		{minus(4, 2), 2},
		{minus(7, 6), 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.actual)
	}
}

func divide(a, b float32) float32 {
	return a / b
}

func TestDivide(t *testing.T) {
	tests := []struct {
		actual   float32
		expected float32
	}{
		{divide(3, 2), 1.5},
		{divide(5, 2), 2.5},
		{divide(3, 1), 3},
		{divide(8, 4), 2},
		{divide(8, 2), 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.actual)
	}
}

func times(a, b int) int {
	return a * b
}

func TestTimes(t *testing.T) {
	tests := []struct {
		actual   int
		expected int
	}{
		{times(3, 4), 12},
		{times(2, 3), 6},
		{times(1, 3), 3},
		{times(4, 5), 20},
		{times(7, 6), 42},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.actual)
	}
}
