package main

import (
	"reflect"
	"testing"
)

func TestGetDistances(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	want := []int{2, 1, 0, 1, 2, 5}

	got := getDistances(left, right)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("getDistances(left, right) = %v, want %v", got, want)
	}
}

func TestSumOfInts(t *testing.T) {
	ints := []int{2, 1, 0, 1, 2, 5}

	want := 11
	got := sumOfInts(ints)

	if got != want {
		t.Fatalf("sumOfInts(ints) = %v, want %v", got, want)
	}
}
