package unionfind

import "fmt"

type UnionFind struct {
	par []int
	siz []int
}

func (uf *UnionFind) Ini(n int) {
	uf.par = make([]int, n)
	uf.siz = make([]int, n)

	for i := 0; i < n; i++ {
		uf.par[i] = -1
		uf.siz[i] = 1
	}
}

func (uf *UnionFind) Root(x int) int {
	if uf.par[x] == -1 {
		return x
	} else {
		uf.par[x] = uf.Root(uf.par[x])
		return uf.par[x]
	}
}

func (uf *UnionFind) Issame(x, y int) bool {
	return uf.Root(x) == uf.Root((y))
}

func (uf *UnionFind) Unite(x, y int) bool {
	x = uf.Root(x)
	y = uf.Root(y)

	if x == y {
		return false
	}

	if uf.siz[x] < uf.siz[y] {
		x, y = y, x //swap
	}

	uf.par[y] = x
	uf.siz[x] += uf.siz[y]
	return true
}

func (uf *UnionFind) Size(x int) int {
	return uf.siz[uf.Root(x)]
}

func main() {
	var uf UnionFind
	uf.Ini(7)

	uf.Unite(1, 2)
	uf.Unite(2, 3)
	uf.Unite(5, 6)

	fmt.Println(uf.Issame(1, 3))
	fmt.Println(uf.Issame(2, 5))

	uf.Unite(1, 6)
	fmt.Println(uf.Issame(2, 5))
}
