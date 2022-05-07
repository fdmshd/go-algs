package graphs

func ShortestWay(g [][]int, u, w int) {
	d := make([]int, len(g))
	for i := 0; i < len(d); i++ {
		d[i] = 9223372036854775807
	}
	d[u] = 0

	p := TopologicalSort(g)
	for i:=0;i<len(d);i++{
		for j:
	}
}
