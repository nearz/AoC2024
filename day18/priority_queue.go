package day18

type StateHeap []nodeState

func (h StateHeap) Len() int      { return len(h) }
func (h StateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h StateHeap) Less(i, j int) bool {
	return h[i].steps < h[j].steps
}
func (h *StateHeap) Push(x any) { *h = append(*h, x.(nodeState)) }
func (h *StateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
