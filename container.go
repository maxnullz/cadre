package ngxnet

import (
	"container/heap"
)

type item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap.
}

type priorityQueue struct {
	s  []*item
	m  map[int]*item
	cf func(l, r *item) bool
}

func minHeapComp(l, r *item) bool { return l.priority < r.priority }
func maxHeapComp(l, r *item) bool { return l.priority > r.priority }

func (r *priorityQueue) Len() int           { return len(r.s) }
func (r *priorityQueue) Less(i, j int) bool { return r.cf(r.s[i], r.s[j]) }
func (r *priorityQueue) Swap(i, j int) {
	r.s[i], r.s[j] = r.s[j], r.s[i]
	r.s[i].index = i
	r.s[j].index = j
}

func (r *priorityQueue) Push(x interface{}) {
	it := x.(*item)
	if _, ok := r.m[it.value]; !ok {
		n := len(r.s)
		it.index = n
		r.s = append(r.s, it)
		r.m[it.value] = it
	} else {
		LogWarn("heap can't insert repeated value:%v", it.value)
	}
}

func (r *priorityQueue) Pop() interface{} {
	n := len(r.s)
	it := r.s[n-1]
	r.s = r.s[0 : n-1]
	delete(r.m, it.value)
	return it
}

type Heap struct {
	p *priorityQueue
}

func (r *Heap) Push(priority, value int) {
	heap.Push(r.p, &item{priority: priority, value: value})
}

func (r *Heap) Pop() int {
	return heap.Pop(r.p).(*item).value
}

func (r *Heap) Update(value, priority int) {
	if it, ok := r.p.m[value]; ok {
		it.priority = priority
		heap.Fix(r.p, it.index)
	}
}

func (r *Heap) Len() int {
	return len(r.p.m)
}

func NewMinHeap() *Heap {
	h := &Heap{p: &priorityQueue{m: map[int]*item{}, cf: minHeapComp}}
	heap.Init(h.p)
	return h
}

func NewMaxHeap() *Heap {
	h := &Heap{p: &priorityQueue{m: map[int]*item{}, cf: maxHeapComp}}
	heap.Init(h.p)
	return h
}