package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		wantErr   bool
		wantInput Input
	}{
		{
			name:    "Valid input",
			lines:   []string{"42|69", "69|42", "", "12,13,14", "2,3,4"},
			wantErr: false,
			wantInput: Input{
				rules: Rules{
					"42|69": true,
					"69|42": true,
				},
				updates: []Update{
					{12, 13, 14},
					{2, 3, 4},
				},
			},
		},
		{
			name:      "Missing blank line",
			lines:     []string{"42|69", "69|42", "42|69", "69|42"},
			wantErr:   true,
			wantInput: Input{},
		},
		{
			name:      "Missing updates",
			lines:     []string{"42|69", "69|42"},
			wantErr:   true,
			wantInput: Input{},
		},
		{
			name:      "Missing rules",
			lines:     []string{"12,13,14", "2,3,4"},
			wantErr:   true,
			wantInput: Input{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			input, err := parseInput(test.lines)
			if (err != nil) != test.wantErr {
				t.Errorf("parseInput() error = %v, want %v", err, test.wantErr)
			}

			if !reflect.DeepEqual(input.rules, test.wantInput.rules) {
				t.Errorf("parseInput() Input.rules = %v, want %v", input.rules, test.wantInput.rules)
			}

			if !reflect.DeepEqual(input.updates, test.wantInput.updates) {
				t.Errorf("parseInput() Input.updates = %v, want %v", input.updates, test.wantInput.updates)
			}
		})
	}
}
