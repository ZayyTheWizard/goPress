package huffman

func encodeWithTree(t *Tree, content string) []rune {
	ans := make([]rune, 0)

	var backtrack func(rune, []rune, *Tree) (bool, []rune)
	backtrack = func(whatIneed rune, temp []rune, t *Tree) (bool, []rune) {
		if t == nil {
			return false, temp
		}

		if t.Letter == whatIneed {
			return true, temp
		}

		temp = append(temp, '0')
		found, newTemp := backtrack(whatIneed, temp, t.left)
		if found {
			return true, newTemp
		}
		temp = temp[:len(temp)-1]

		temp = append(temp, '1')
		found, newTemp = backtrack(whatIneed, temp, t.right)
		if found {
			return true, newTemp
		}
		temp = temp[:len(temp)-1]

		return false, temp
	}

	for _, a := range content {
		temp := make([]rune, 0)
		_, current := backtrack(a, temp, t)
		ans = append(ans, current...)

	}

	return ans
}
