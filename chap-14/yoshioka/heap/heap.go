package heap

type Pair struct {
	Dist int
	Idx  int
}

type Heap []Pair

func (heap *Heap) Push(x Pair) {
	*heap = append(*heap, x)

	i := len(*heap) - 1
	for i > 0 {
		p := (i - 1) / 2 //top number
		if (*heap)[p].Dist >= x.Dist {
			break
		}
		(*heap)[i] = (*heap)[p]
		i = p
	}
	(*heap)[i].Dist = x.Dist
}

func (heap *Heap) Top() Pair {
	if len(*heap) > 0 {
		return (*heap)[0]
	} else {
		return Pair{-1, -1}
	}
}

func (heap *Heap) Pop() {
	if len(*heap) == 0 {
		return
	}
	x := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]

	i := 0
	for i*2+1 < len(*heap) {
		child1 := i*2 + 1
		child2 := i*2 + 2
		if (child2 < len(*heap)) && ((*heap)[child2].Dist > (*heap)[child1].Dist) {
			child1 = child2
		}
		if (*heap)[child1].Dist <= x.Dist {
			break
		}
		(*heap)[i] = (*heap)[child1]
		i = child1
	}
	(*heap)[i] = x
}
