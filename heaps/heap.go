package heaps

import (
	"fmt"
	"sync"
)

type maxHeap struct {
	array  []int
	length int
	lock   *sync.Mutex
}

type minHeap struct {
	array  []int
	length int
	lock   *sync.Mutex
}

//NewMaxHeap returns a new maxHeap
func NewMaxHeap() *maxHeap {
	return &maxHeap{lock: &sync.Mutex{}}
}

//NewMinHeap returns a new maxHeap
func NewMinHeap() *minHeap {
	return &minHeap{lock: &sync.Mutex{}}
}

//Insert inserts given key in accordance with the heap properties
func (h *maxHeap) Insert(key int) {
	h.lock.Lock()
	h.array = append(h.array, key)
	h.length++
	child := h.length - 1
	//insert at the end,
	//and keep going up as long as key>parent
	for key > h.array[parent(child)] {
		par := parent(child)
		h.array[par], h.array[child] = h.array[child], h.array[par]
		child = par
	}
	h.lock.Unlock()
}

//Insert inserts given key in accordance with the heap properties
func (h *minHeap) Insert(key int) {
	h.lock.Lock()
	h.array = append(h.array, key)
	h.length++
	child := h.length - 1
	//insert at the end,
	//and keep going up as long as key<parent
	for key < h.array[parent(child)] {
		par := parent(child)
		h.array[par], h.array[child] = h.array[child], h.array[par]
		child = par
	}
	h.lock.Unlock()
}

//Delete removes given key if exist and heapify rest
func (h *maxHeap) Delete(key int) error {
	if h.length == 0 {
		return fmt.Errorf("heap is empty already")
	}
	ok, ind := h.Search(key)
	if !ok {
		return fmt.Errorf("key %d could not found", key)
	}
	h.lock.Lock()
	h.array[ind] = h.array[h.length-1]
	h.array = h.array[:h.length-1]
	h.length--
	//heapify
	h.array = BuildMaxHeap(h.array...).array
	h.lock.Unlock()
	return nil
}

//Delete removes given key if exist and heapify rest
func (h *minHeap) Delete(key int) error {
	if h.length == 0 {
		return fmt.Errorf("heap is empty already")
	}
	ok, ind := h.Search(key)
	if !ok {
		return fmt.Errorf("key %d could not found", key)
	}
	h.lock.Lock()
	h.array[ind] = h.array[h.length-1]
	h.array = h.array[:h.length-1]
	h.length--
	//heapify
	h.array = BuildMinHeap(h.array...).array
	h.lock.Unlock()
	return nil
}

//Search looks up for given key and returns boolean result and keys index if exist
func (h *maxHeap) Search(key int) (bool, int) {
	if h.length == 0 {
		return false, -1
	}
	for i := 0; i < h.length; i++ {
		if h.array[i] == key {
			return true, i
		}
	}
	return false, -1
}

//Search looks up for given key and returns boolean result and keys index if exist
func (h *minHeap) Search(key int) (bool, int) {
	if h.length == 0 {
		return false, -1
	}
	for i := 0; i < h.length; i++ {
		if h.array[i] == key {
			return true, i
		}
	}
	return false, -1
}

//SearchForLatest returns total count of given key and its encountered latest index if exist.
//In case of absence, returns -1,-1. This is useful especially for larger heaps
func (h *maxHeap) SearchForLatest(key int) (int, int) {
	if h.length == 0 {
		return -1, -1
	}
	count := 0
	found := -1
	var lookUp func(int) (bool, int)
	lookUp = func(i int) (bool, int) {
		if key > h.array[i] { //return if search key is bigger than subRoot
			return false, -1
		}
		if h.array[i] == key {
			count++
			if i >= found {
				found = i
			}
		}
		l, r := leftChild(i), rightChild(i)
		if h.length > l {
			lookUp(l)
		}
		if h.length > r {
			lookUp(r)
		}
		return false, -1
	}
	lookUp(0)
	if found >= 0 {
		return count, found
	}
	return -1, -1
}

//BuildMaxHeap returns a new heap created with given parameters
func BuildMaxHeap(vars ...int) *maxHeap {
	temp := NewMaxHeap()
	for _, i := range vars {
		temp.Insert(i)
	}
	return temp
}

//BuildMinHeap returns a new heap created with given parameters
func BuildMinHeap(vars ...int) *minHeap {
	temp := NewMinHeap()
	for _, i := range vars {
		temp.Insert(i)
	}
	return temp
}

//parent returns parent of given index
func parent(ind int) int {
	return (ind - 1) / 2
}

//leftChild returns left child of given node index
func leftChild(ind int) int {
	return 2*ind + 1
}

//rightChild return right child of given node index
func rightChild(ind int) int {
	return 2*ind + 2
}

//Length returns total key count of heap
func (h *maxHeap) Length() int {
	return h.length
}

//Length returns total key count of heap
func (h *minHeap) Length() int {
	return h.length
}

//Max return maximum valued key of heap
func (h *maxHeap) Max() int {
	if h.length > 0 {
		return h.array[0]
	}
	panic("heap is empty")
}

//Min returns minimum valued key of heap
func (h *maxHeap) Min() int {
	if h.length > 0 {
		min := h.array[0]
		for _, i := range h.array {
			if i < min {
				min = i
			}
		}
		return min
	}
	panic("heap is empty")
}

//Min returns minimum valued key of heap
func (h *minHeap) Min() int {
	if h.length > 0 {
		return h.array[0]
	}
	panic("heap is empty")
}

//Max returns maximum valued key of heap
func (h *minHeap) Max() int {
	if h.length > 0 {
		max := 0
		for _, i := range h.array {
			if i > max {
				max = i
			}
		}
		return max
	}
	panic("heap is empty")
}
