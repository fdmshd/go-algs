package substrings

func CalculeteShifts(t []rune) map[rune]int {
	res := make(map[rune]int)
	for i, val := range t {
		res[val] = len(t) - i - 1
	}
	return res
}

func BoyerMoore(str, template []rune) int {
	shifts := CalculeteShifts(template)
	m := len(template)
	n := len(str)
	i := m - 1
	j := m - 1
	for i < n {
		if str[i] == template[j] {
			i--
			j--
			if j == 0 {
				return i
			}
		} else {
			shift := shifts[str[i]]
			if shift == 0 {
				i += m
			} else {
				i += shift
			}
			j = m - 1
		}
	}
	return -1
}
