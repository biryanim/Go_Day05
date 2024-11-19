package heap

import (
	"container/heap"
	"errors"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (ph PresentHeap) Len() int { return len(ph) }
func (ph PresentHeap) Less(i, j int) bool {
	if ph[i].Value == ph[j].Value {
		return ph[i].Size < ph[j].Size
	}
	return ph[i].Value > ph[j].Value
}
func (ph PresentHeap) Swap(i, j int) { ph[i], ph[j] = ph[j], ph[i] }

func (ph *PresentHeap) Push(x any) {
	*ph = append(*ph, x.(Present))
}

func (ph *PresentHeap) Pop() any {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

func getNCoolestPresents(presents []Present, n int) ([]Present, error) {
	if n < 0 || n > len(presents) {
		return nil, errors.New("n is out of range")
	}
	h := PresentHeap(presents)
	heap.Init(&h)

	coolestPresents := make([]Present, 0, n)
	for i := 0; i < n; i++ {
		coolestPresents = append(coolestPresents, heap.Pop(&h).(Present))
	}
	return coolestPresents, nil
}

func IsValid(presents []Present) bool {
	for _, present := range presents {
		if present.Size <= 0 {
			return false
		}
	}
	return true
}

func grabPresents(presents []Present, capacity int) ([]Present, error) {
	if capacity < 0 {
		return nil, errors.New("capacity is less than 0")
	}
	if !IsValid(presents) {
		return nil, errors.New("invalid input")
	}
	n := len(presents)
	backpack := make([][]int, n+1)
	for i := range backpack {
		backpack[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if presents[i-1].Size <= j {
				backpack[i][j] = max(backpack[i-1][j], presents[i-1].Value+backpack[i-1][j-presents[i-1].Size])
			} else {
				backpack[i][j] = backpack[i-1][j]
			}
		}
	}

	w := capacity
	res := make([]Present, 0, n)
	for i := n; i > 0 && w > 0; i-- {
		if backpack[i][w] != backpack[i-1][w] {
			res = append(res, presents[i-1])
			w -= presents[i-1].Size
		}
	}

	return res, nil
}

/*
[0 0 	0 	 0	  0   ]
[0 0 	0 	 2000 2000]
[0 0 	0 	 2000 3000]
[0 1500 1500 2000 3500]
*/
