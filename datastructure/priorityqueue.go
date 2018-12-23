package datastructure

import "container/heap"

type Comparable interface {
	Compare(other Comparable) int
}

// An Item is something we manage in a priority queue.
type Item struct {
	Value Comparable
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value.Compare(pq[j].Value) < 0
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0: n-1]
	return item
}

func (pq *PriorityQueue) Enqueue(val interface{}) {
	// Insert a new item and then modify its priority.
	item := &Item{
		Value: val.(Comparable),
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue) Dequeue() interface{}{
	item := heap.Pop(pq).(*Item)
	return item.Value
}

func NewPriorityQueue(len int) *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}
