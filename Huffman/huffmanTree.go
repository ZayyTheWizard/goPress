package huffman

import "fmt"

type Tree struct {
	Letter    rune
	Frequency int
	right     *Tree
	left      *Tree
}

func NewNode(Letter rune, freq int) *Tree {
	return &Tree{
		Frequency: freq,
		Letter:    Letter,
		right:     nil,
		left:      nil,
	}
}

func printTree(t *Tree) {
	if t == nil {
		return
	}

	printTree(t.left)
	fmt.Printf("%c, %v\n", t.Letter, t.Frequency)
	printTree(t.right)
}
