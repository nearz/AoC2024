package day16

type StateHeap []state

func (h StateHeap) Len() int      { return len(h) }
func (h StateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h StateHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h *StateHeap) Push(x any) { *h = append(*h, x.(state)) }
func (h *StateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TrackHeap []track

func (h TrackHeap) Len() int      { return len(h) }
func (h TrackHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h TrackHeap) Less(i, j int) bool {
	return h[i].s.cost < h[j].s.cost
}
func (h *TrackHeap) Push(x any) { *h = append(*h, x.(track)) }
func (h *TrackHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
