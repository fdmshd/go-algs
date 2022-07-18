package substrings

func FindPrefix(s []rune) []int {
	j, i := 0, 1
	pref := make([]int, len(s))
	for i < len(s) {
		if s[i] != s[j] {
			if j == 0 {
				pref[i] = 0
				i++
			} else {
				j = pref[j-1]
			}
		} else {
			pref[i] = j + 1
			i++
			j++
		}
	}
	return pref
}

func KMP(s, t []rune) int {
	p := FindPrefix(t)
	i, j := 0, 0
	n := len(s)
	m := len(t)
	for i < n {
		if s[i] == t[j] {
			i++
			j++
			if j == m {
				return i - m
			}
		} else {
			if j > 0 {
				j = p[j] - 1
			} else {
				i++
			}
		}

	}
	return -1
}
