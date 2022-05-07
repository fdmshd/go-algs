package graphs

func DoDfs(g [][]int) {
	n := len(g)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, g, visited)
		}

	}

}
func dfs(u int, g [][]int, visited []bool) {
	visited[u] = true
	for i := 0; i < len(g[u]); i++ {
		if !visited[g[u][i]] {
			dfs(g[u][i], g, visited)
		}
	}
}
