package graphs

func TopologicalSort(g [][]int) []int {
	n := len(g)

	inDegree := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < len(g[i]); j++ {
			inDegree[g[i][j]] += 1
		}
	}
	var next []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			next = append(next, i)
		}
	}
	var result []int
	currentVert := 0
	for len(next) > 0 {
		currentVert = next[0]
		next = next[1:]
		result = append(result, currentVert)
		for i := 0; i < len(g[currentVert]); i++ {
			inDegree[g[currentVert][i]] -= 1
			if inDegree[g[currentVert][i]] == 0 {
				next = append(next, g[currentVert][i])
			}
		}
	}
	return result
}
