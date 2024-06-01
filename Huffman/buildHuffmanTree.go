package huffman

func buildHuffmanTree(pairs PairList, size int) *Tree {
	root := make([]*Tree, 0, size)

	for i := 0; i < size; i += 1 {
		root = append(root, NewNode(pairs[i].Key, pairs[i].Value))
	}

	// Build Huffman tree
	for size > 1 {
		// Find the two nodes with the lowest frequencies
		minIndex1, minIndex2 := 0, 1
		if root[minIndex1].Frequency > root[minIndex2].Frequency {
			temp := minIndex1
			minIndex1 = minIndex2
			minIndex2 = temp
		}
		for i := 2; i < size; i++ {
			if root[i].Frequency < root[minIndex1].Frequency {
				minIndex2 = minIndex1
				minIndex1 = i
			} else if root[i].Frequency < root[minIndex2].Frequency {
				minIndex2 = i
			}
		}

		// Create a new internal node by merging the two nodes with lowest frequencies
		aNode := NewNode('$', root[minIndex1].Frequency+root[minIndex2].Frequency)
		aNode.left = root[minIndex1]
		aNode.right = root[minIndex2]

		// Remove the two nodes with lowest frequencies from the array and add the new node
		root[minIndex1] = aNode
		root[minIndex2] = root[size-1]
		size--
	}

	return root[0]
}
