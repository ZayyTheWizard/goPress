package huffman

import (
	"fmt"
	"sort"
)

// Pair holds a key-value pair.
type Pair struct {
	Key   rune
	Value int
}

// A slice of pairs that implements sort.Interface to sort by value.
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// Function to convert a map to a sorted PairList.
func sortMapByValue(m map[rune]int) PairList {
	// Create a PairList from the map.
	pl := make(PairList, 0, len(m))
	for k, v := range m {
		pl = append(pl, Pair{k, v})
	}

	// Sort the PairList by value.
	sort.Sort(pl)
	return pl
}

func Encode(Content string) (string, error) {

	// Build Frequency Table
	table := buildTable(Content)

	// Sort table
	sortedTable := sortMapByValue(table)

	// Build Tree
	huffmanTree := buildHuffmanTree(sortedTable, len(sortedTable))

	// Encode using Tree
	encoded := encodeWithTree(huffmanTree, Content)
	fmt.Printf("%c", encoded)
	// Searilize Tree

	// Return Encoded string and searilized Tree Concatinated
	return "", nil
}
