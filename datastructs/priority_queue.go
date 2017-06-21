// Package datastructs defines a reference implementation of
// a priority queue as shown on the official
// documentation of Go but modified to suit our purposes
//
// https://golang.org/pkg/container/heap/#example__priorityQueue
//
// The only difference is that the value of each item in the queue is a map
// of file readers to a returned integer. This will help us keep track from
// which file we should read next
package datastructs

import (
	"bufio"
	"container/heap"
)

// Item is a 'node' in the priority queue
type Item struct {
	Value    map[*bufio.Reader]int // The value of the item; arbitrary.
	Priority int                   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Len gives the length of the queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less makes a comparison of two elements in the queue
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps two elements of a priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item in the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes an item from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value map[*bufio.Reader]int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
