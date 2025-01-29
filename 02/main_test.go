package main

import (
	"reflect"
	"testing"
)

func TestGetDistances(t *testing.T) {
	tests := []struct {
		name      string
		report    []int
		wantValid bool
	}{
		{
			name:      "Valid report",
			report:    []int{7, 6, 4, 2, 1},
			wantValid: true,
		},
		{
			name:      "Increase is too big",
			report:    []int{1, 2, 7, 8, 9},
			wantValid: false,
		},
		{
			name:      "Decrease too big",
			report:    []int{9, 7, 6, 2, 1},
			wantValid: false,
		},
		{
			name:      "Switch from increase to decrease",
			report:    []int{1, 3, 2, 4, 5},
			wantValid: false,
		},
		{
			name:      "Neither increase nor decrease",
			report:    []int{8, 6, 4, 4, 1},
			wantValid: false,
		},
		{
			name:      "Valid report",
			report:    []int{1, 3, 6, 7, 9},
			wantValid: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			valid := isReportValid(test.report)
			if valid != test.wantValid {
				t.Fatalf("isReportValid(report) = %v, want %v", test.wantValid, valid)
			}
		})
	}
}

func TestRemoveLevelAt(t *testing.T) {
	report := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 4, 5}
	got := removeLevelAt(report, 2)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("removeLevelAt(report, index) = %v, want %v", got, want)
	}
}

func TestIsDampenedReportValid(t *testing.T) {
	tests := []struct {
		name      string
		report    []int
		wantValid bool
	}{
		{
			name:      "Valid report",
			report:    []int{7, 6, 4, 2, 1},
			wantValid: true,
		},
		{
			name:      "Increase is too big",
			report:    []int{1, 2, 7, 8, 9},
			wantValid: false,
		},
		{
			name:      "Decrease is too big",
			report:    []int{9, 7, 6, 2, 1},
			wantValid: false,
		},
		{
			name:      "Valid report",
			report:    []int{1, 3, 2, 4, 5},
			wantValid: true,
		},
		{
			name:      "Valid report",
			report:    []int{8, 6, 4, 4, 1},
			wantValid: true,
		},
		{
			name:      "Valid report",
			report:    []int{1, 3, 6, 7, 9},
			wantValid: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			valid := isDampenedReportValid(test.report)
			if valid != test.wantValid {
				t.Fatalf("isDampenedReportValid(report) = %v, want %v", test.wantValid, valid)
			}
		})
	}
}
