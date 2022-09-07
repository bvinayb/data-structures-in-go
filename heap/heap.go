package heap

//Get Child by simple formula (Parent_Index X 2 + 1 = Left_ChildIndex) e.g. give me left child for index 19 ~~ [19] X 2 + 1 = [39]
type MaxHeap struct {
	arrays []int
}

func (h *MaxHeap) Insert(key int) {
	h.arrays = append(h.arrays, key)
	h.MaxHeapify(len(h.arrays) - 1)
}

//func (h *MaxHeap) Extract() int {
//	extracted := h.arrays[0]
//}

func (h *MaxHeap) MaxHeapify(index int) {
	for h.arrays[parent(index)] > h.arrays[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	//swapping values of i1 to i2 and i2 to i1
	h.arrays[i1], h.arrays[i2] = h.arrays[i2], h.arrays[i1]
}
