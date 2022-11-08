package sorts

import "log"

// broken
func LSD(arr []int) {
	n := len(arr)
	m := countDigs(arr[0])
	newArr := make([]int, n)
	for j := m; j >= 0; j-- {
		p := make([]int8, 10)
		for i := 0; i < n; i++ {
			pn := digit(arr[i], j, m)
			if pn < 0 {
				log.Printf("arr[%d]=%d, j:%d,m:%d", i, arr[i], j, m)
			}
			p[pn]++
		}
		for k := 1; k < len(p); k++ {
			p[k] += p[k-1]
		}
		for k := len(p) - 1; k > 0; k-- {
			p[k] = p[k-1]
		}
		p[0] = 0
		for i := 0; i < n; i++ {
			d := digit(arr[i], j, m)
			newArr[p[d]] = arr[i]
			p[d]++
		}
		copy(arr, newArr)
	}
}

func digit(numb, dig, n int) int {
	for i := n; i > dig; i-- {
		numb /= 10
	}
	return numb % 10
}

func countDigs(num int) int {
	digs := 0
	for num != 0 {
		digs++
		num /= 10
	}
	return digs - 1
}
