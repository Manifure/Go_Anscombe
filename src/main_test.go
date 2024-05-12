package main

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	test := mean(nums)
	var correct float64 = 8
	if test != correct {
		t.Errorf("mean = %f, want %f", test, correct)
	}
}

func TestMedian(t *testing.T) {
	nums := []int{79, 26, 81, 48, 52, 74, 95, 36, 66, 1, 110, 24, 31, 85}
	test := median(nums)
	var correct float64 = 59
	if test != correct {
		t.Errorf("median = %f, want %f", test, correct)
	}
}

func TestMode(t *testing.T) {
	nums := []int{48, 43, 38, 32, 15, 12, 38, 15, 21, 25}
	test := mode(nums)
	var correct int = 15
	if test != correct {
		t.Errorf("mode = %d, want %d", test, correct)
	}
}

func TestSD(t *testing.T) {
	nums := []int{8, 12, 23, 22, 16, 21, 20, 24}
	test := sd(nums)
	test = math.Round(test*1000) / 1000
	var correct float64 = 5.356
	if test != correct {
		t.Errorf("SD = %f, want %f", test, correct)
	}
}
