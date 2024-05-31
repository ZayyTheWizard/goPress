package huffman

import (
	"reflect"
	"testing"
)

func testBuildTable(t *testing.T) {
	var Test = []struct {
		input string
		want  map[rune]int
	}{
		{"AAABBCCCCD", map[rune]int{'A': 3, 'B': 2, 'C': 4, 'D': 1}},
		{"KKKK!!OO)))PPPP", map[rune]int{'K': 4, '!': 2, 'O': 2, ')': 3, 'P': 4}},
		{"AaBbCcDdEeFf", map[rune]int{'A': 1, 'a': 1, 'B': 1, 'b': 1, 'C': 1, 'c': 1, 'D': 1, 'd': 1, 'E': 1, 'e': 1, 'F': 1, 'f': 1}},
		{"_-0+=123444567890_+=", map[rune]int{'_': 2, '-': 1, '+': 2, '1': 1, '2': 1, '3': 1, '4': 3, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1, '0': 2, '=': 2}},
	}

	for _, tt := range Test {
		t.Run("Test Build Table", func(t *testing.T) {
			ans := buildTable(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
