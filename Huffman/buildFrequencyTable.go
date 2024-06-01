package huffman

func buildTable(Content string) map[rune]int {
	var table = make(map[rune]int)

	for _, char := range Content {
		_, ok := table[char]
		if ok {
			table[char] += 1
		} else {
			table[char] = 1
		}
	}

	return table
}
