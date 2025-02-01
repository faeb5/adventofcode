package main

import (
	"reflect"
	"testing"
)

func TestSplitInputLines(t *testing.T) {
	tests := []struct {
		name            string
		inputLines      []string
		wantRuleLines   RuleLines
		wantUpdateLines UpdateLines
		wantErr         bool
	}{
		{
			name:            "Valid input",
			inputLines:      []string{"rule1", "rule2", "", "update1", "update2"},
			wantRuleLines:   RuleLines{"rule1", "rule2"},
			wantUpdateLines: UpdateLines{"udpate1", "update2"},
			wantErr:         false,
		},
		{
			name:            "No empty line",
			inputLines:      []string{"rule1", "update1"},
			wantRuleLines:   RuleLines{},
			wantUpdateLines: UpdateLines{},
			wantErr:         true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			ruleLines, updateLines, err := splitInputLines(test.inputLines)

			if (err != nil) != test.wantErr {
				t.Errorf("splitInputLines() error = %q, want %v", err, test.wantErr)
			}

			if !reflect.DeepEqual(ruleLines, test.wantRuleLines) {
				t.Errorf("splitInputLines() RuleLines = %v, want %v", ruleLines, test.wantRuleLines)
			}

			if !reflect.DeepEqual(ruleLines, test.wantRuleLines) {
				t.Errorf("splitInputLines() UpdateLines = %v, want %v", updateLines, test.wantUpdateLines)
			}
		})
	}
}

func TestParseRules(t *testing.T) {
	tests := []struct {
		name      string
		input     RuleLines
		wantRules []Rule
		wantErr   bool
	}{
		{
			name:      "Valid rules",
			input:     RuleLines{"97|13", "45|16"},
			wantRules: []Rule{{97, 13}, {45, 16}},
			wantErr:   false,
		},
		{
			name:      "Rule contains NaN",
			input:     RuleLines{"97|ERROR"},
			wantRules: []Rule{},
			wantErr:   true,
		},
		{
			name:      "Missing split sign",
			input:     RuleLines{"9713"},
			wantRules: []Rule{},
			wantErr:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			rules, err := parseRules(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("parseRules() error = %q, want %v", err, test.wantErr)
			}

			if !reflect.DeepEqual(rules, test.wantRules) {
				t.Errorf("parseRules() []Rule = %v, want %v", rules, test.wantRules)
			}
		})
	}
}

func TestParseUpdates(t *testing.T) {
	tests := []struct {
		name        string
		input       UpdateLines
		wantUpdates []Update
		wantErr     bool
	}{
		{
			name:        "Valid updates",
			input:       UpdateLines{"75,47", "97,61"},
			wantUpdates: []Update{{75, 47}, {97, 61}},
			wantErr:     false,
		},
		{
			name:        "Update contains NaN",
			input:       UpdateLines{"97, ERROR, 44"},
			wantUpdates: []Update{},
			wantErr:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			updates, err := parseUpdates(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("parseUpdates() error = %q, want %v", err, test.wantErr)
			}

			if !reflect.DeepEqual(updates, test.wantUpdates) {
				t.Errorf("parseUpdates() []Update = %v, want %v", updates, test.wantUpdates)
			}
		})
	}
}
