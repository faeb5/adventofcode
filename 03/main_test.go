package main

import (
	"testing"
)

const (
	example_one = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	example_two = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
)

func TestGetSum(t *testing.T) {
	want := 161
	got := getSum(example_one)
	if got != want {
		t.Fatalf("getSum(line) = %v, got %v", got, want)
	}
}

func TestSanitizeLine(t *testing.T) {
	wantLine := "mul(2,4)mul(8,5)"
	wantEnabled := true
	line, enabled := sanitizeLine(example_two, true)
	if line != wantLine {
		t.Fatalf("sanitizeLine(line, enabled) string = %v, want %v", line, wantLine)
	}
	if enabled != wantEnabled {
		t.Fatalf("sanitizeLine(line, enabled) bool = %v, want %v", enabled, wantEnabled)
	}
}
