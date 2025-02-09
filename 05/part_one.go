package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	val    int
	before *[]int
	after  *[]int
}

func solvePartOne(content string) {
	sum := 0

	rules, updates := parseRulesAndUpdates(content)
	for _, update := range updates {
		if isUpdateOK(update, rules) {
			sum += getMedianOf(update)
		}
	}

	fmt.Println(sum)
}

func getMedianOf(update []int) int {
	return update[len(update)/2]
}

func isUpdateOK(update []int, rules map[int]Rule) bool {
	for i, val := range update {
		rule, ok := rules[val]
		if !ok {
			return false
		}

		// check before
		for j := 0; j < i; j++ {
			if !slices.Contains(*rule.before, update[j]) {
				return false
			}
		}

		// check after
		for j := i + 1; j < len(update); j++ {
			if !slices.Contains(*rule.after, update[j]) {
				return false
			}
		}
	}
	return true
}

func parseRulesAndUpdates(content string) (rules map[int]Rule, updates [][]int) {
	parts := strings.Split(content, "\n\n")

	// rules
	ruleLines := strings.Split(parts[0], "\n")
	rules = make(map[int]Rule, len(ruleLines)-1)
	for _, line := range ruleLines {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, "|")

		lhs, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}

		rhs, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}

		rule, ok := rules[lhs]
		if !ok {
			rules[lhs] = Rule{
				val:    lhs,
				before: &[]int{},
				after:  &[]int{rhs},
			}
		} else {
			*rule.after = append(*rule.after, rhs)
		}

		rule, ok = rules[rhs]
		if !ok {
			rules[rhs] = Rule{
				val:    rhs,
				before: &[]int{lhs},
				after:  &[]int{},
			}
		} else {
			*rule.before = append(*rule.before, lhs)
		}
	}

	// sort after and before arrays for all rules
	for _, rule := range rules {
		sort.Ints(*rule.after)
		sort.Ints(*rule.before)
	}

	// updates
	updateLines := strings.Split(parts[1], "\n")
	updates = make([][]int, len(updateLines)-1)
	for i, line := range updateLines {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, ",")
		update := make([]int, len(nums))
		for j, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			update[j] = num
		}
		updates[i] = update
	}

	return rules, updates
}
