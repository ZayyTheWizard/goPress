package huffman

import (
	"fmt"
	"sort"
)

func Encode(Content string) (string, error) {

	// Build Frequency Table
	table := buildTable(Content)

	// Sort table Decending order

	keys := make([]rune, 0, len(table))

	for key := range table {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return table[keys[i]] < table[keys[j]]
	})

	fmt.Println(keys)
	for r := range keys {
		fmt.Printf("%d", r)
	}

	// Build Tree
	buildHuffmanTree(Content)

	// Encode using Tree

	// Return encode string
	return "", nil
}
