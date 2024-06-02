package huffman

import (
	"reflect"
	"testing"
)

// Tree structure for Huffman encoding
func TestEncodeWithTree(t *testing.T) {
	// Build a sample Huffman Tree
	root := &Tree{
		left: &Tree{
			Letter:    'a',
			Frequency: 9,
		},
		right: &Tree{
			left: &Tree{
				Letter:    'b',
				Frequency: 2,
			},
			right: &Tree{
				Letter:    'c',
				Frequency: 3,
			},
		},
	}

	tests := []struct {
		content string
		want    []rune
	}{
		{"a", []rune{'0'}},
		{"b", []rune{'1', '0'}},
		{"c", []rune{'1', '1'}},
		{"abc", []rune{'0', '1', '0', '1', '1'}},
	}

	for _, tt := range tests {
		t.Run(tt.content, func(t *testing.T) {
			got := encodeWithTree(root, tt.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeWithTree(%s) = %v, want %v", tt.content, got, tt.want)
			}
		})
	}
}
