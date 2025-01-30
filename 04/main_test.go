package main

import (
	"reflect"
	"testing"
)

func TestToRuneArray(t *testing.T) {
	input := []string{"ABCD", "EFGH"}
	want := [][]rune{{65, 66, 67, 68}, {69, 70, 71, 72}}
	got := toRuneArray(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("toRuneArray(input) = %v, want %v", got, want)
	}
}

func TestSearchRight(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune("......XMAS")},
			point: [2]int{0, 6},
			want:  true,
		},
		{
			name:  "Length too short",
			runes: [][]rune{[]rune("......XMA")},
			point: [2]int{0, 6},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune("......XMAX")},
			point: [2]int{0, 6},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchRight(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchRight(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchLeft(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune("SAMX......")},
			point: [2]int{0, 3},
			want:  true,
		},
		{
			name:  "Length of runes too short",
			runes: [][]rune{[]rune("AMX.......")},
			point: [2]int{0, 2},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune("MAMX......")},
			point: [2]int{0, 3},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchLeft(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchLeft(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchDown(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune(".........."), []rune(".X........"), []rune(".M........"), []rune(".A........"), []rune(".S........")},
			point: [2]int{1, 1},
			want:  true,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune(".........."), []rune(".X........"), []rune(".M........"), []rune(".A........")},
			point: [2]int{1, 1},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune(".........."), []rune(".X........"), []rune(".M........"), []rune(".A........"), []rune(".M........")},
			point: [2]int{1, 1},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchDown(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchDown(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchTop(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune("..S......."), []rune("..A......."), []rune("..M......."), []rune("..X......."), []rune("..........")},
			point: [2]int{3, 2},
			want:  true,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune("..A......."), []rune("..M......."), []rune("..X......."), []rune("..........")},
			point: [2]int{2, 2},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune("..M......."), []rune("..A......."), []rune("..M......."), []rune("..X......."), []rune("..........")},
			point: [2]int{3, 2},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchTop(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchTop(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchDownRight(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune(".........."), []rune("......X..."), []rune(".......M.."), []rune("........A."), []rune(".........S")},
			point: [2]int{1, 6},
			want:  true,
		},
		{
			name:  "Length too short",
			runes: [][]rune{[]rune(".........."), []rune(".........."), []rune(".......X.."), []rune("........M."), []rune(".........A")},
			point: [2]int{2, 7},
			want:  false,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune(".........."), []rune(".........."), []rune("......X..."), []rune(".......M.."), []rune("........A.")},
			point: [2]int{2, 6},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune(".........."), []rune("......X..."), []rune(".......M.."), []rune("........A."), []rune(".........M")},
			point: [2]int{1, 6},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchDownRight(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchDownRight(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchTopRight(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune(".........S"), []rune("........A."), []rune(".......M.."), []rune("......X..."), []rune("..........")},
			point: [2]int{3, 6},
			want:  true,
		},
		{
			name:  "Length too short",
			runes: [][]rune{[]rune(".........."), []rune(".........A"), []rune("........M."), []rune(".......X.."), []rune("..........")},
			point: [2]int{3, 7},
			want:  false,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune("........A."), []rune(".......M.."), []rune("......X..."), []rune(".........."), []rune("..........")},
			point: [2]int{2, 6},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune(".........M"), []rune("........A."), []rune(".......M.."), []rune("......X..."), []rune("..........")},
			point: [2]int{3, 6},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchTopRight(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchTopRight(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchDownLeft(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune(".........."), []rune("...X......"), []rune("..M......."), []rune(".A........"), []rune("S.........")},
			point: [2]int{1, 3},
			want:  true,
		},
		{
			name:  "Length too short",
			runes: [][]rune{[]rune(".........."), []rune("..X......."), []rune(".M........"), []rune("A........."), []rune("..........")},
			point: [2]int{1, 2},
			want:  false,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune(".........."), []rune(".........."), []rune("...X......"), []rune("..M......."), []rune(".A........")},
			point: [2]int{2, 3},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune(".........."), []rune("...X......"), []rune("..M......."), []rune(".A........"), []rune("M.........")},
			point: [2]int{1, 3},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchDownLeft(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchDownLeft(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSearchTopLeft(t *testing.T) {
	tests := []struct {
		name  string
		runes [][]rune
		point [2]int
		want  bool
	}{
		{
			name:  "Exact hit",
			runes: [][]rune{[]rune("S........."), []rune(".A........"), []rune("..M......."), []rune("...X......"), []rune("..........")},
			point: [2]int{3, 3},
			want:  true,
		},
		{
			name:  "Length too short",
			runes: [][]rune{[]rune(".........."), []rune("A........."), []rune(".M........"), []rune("..X......."), []rune("..........")},
			point: [2]int{3, 2},
			want:  false,
		},
		{
			name:  "Height too short",
			runes: [][]rune{[]rune("A........."), []rune(".M........"), []rune("..X......."), []rune(".........."), []rune("..........")},
			point: [2]int{2, 2},
			want:  false,
		},
		{
			name:  "No match",
			runes: [][]rune{[]rune("M........."), []rune(".A........"), []rune("..M......."), []rune("...X......"), []rune("..........")},
			point: [2]int{3, 3},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := searchTopLeft(test.runes, []rune("XMAS"), test.point[0], test.point[1])
			if got != test.want {
				t.Fatalf("searchTopLeft(runes, pattern, i, j) = %v, want %v", got, test.want)
			}
		})
	}
}
