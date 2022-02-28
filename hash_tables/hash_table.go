package hash_tables

import (
	"fmt"
	"go-ds/linked_lists"
	"sync"
)

// Go maps are complex implementation of hash tables.
//This package provides only a functional but basic hash table data structure with its methods.
// Collisions are handled with chaining

type hashTable struct {
	length  int                        // total word count
	size    int                        // total count of buckets
	buckets []*linked_lists.LinkedList // github.com/fukaraca/go-ds/linked_list
	lock    *sync.Mutex
}

//NewHashTable creates a new hash table with given size and returns its pointer
func NewHashTable(size int) *hashTable {
	table := []*linked_lists.LinkedList{}

	for i := 0; i < size; i++ {
		table = append(table, linked_lists.New())
	}

	return &hashTable{
		buckets: table,
		size:    size,
		lock:    &sync.Mutex{}}

}

//Insert adds a new element to hash table and returns if exists
func (h *hashTable) Insert(word string) error {
	if word == "" {
		return fmt.Errorf("missing string input")
	}

	if ok, _ := h.Search(word); ok {
		return fmt.Errorf("the word '%s' is already exist", word)
	}
	h.lock.Lock()
	ind := h.hash(word)
	h.buckets[ind].InsertAtTail(word)
	h.length++
	h.lock.Unlock()
	return nil
}

//Search looks up for given word and returns boolean result and error if any exist
func (h *hashTable) Search(word string) (bool, error) {
	if word == "" {
		return false, fmt.Errorf("missing string input")
	}
	head := h.buckets[h.hash(word)].Head
	for head != nil {
		if head.Data.(string) == word {
			return true, nil
		}
		head = head.Next
	}

	return false, fmt.Errorf("word '%s' couldn't be found", word)
}

func (h *hashTable) Delete(word string) error {
	if word == "" {
		return fmt.Errorf("missing string input")
	}
	ind := h.hash(word)
	if *h.buckets[ind].Len() == 0 {
		return fmt.Errorf("word '%s' couldn't be found", word)
	}

	head := h.buckets[ind].Head
	if head.Data.(string) == word {

		h.lock.Lock()
		if *h.buckets[ind].Len() == 1 {
			h.buckets[ind].Head = nil
			h.length--
			*h.buckets[ind].Len()--
			h.lock.Unlock()
			return nil
		}
		h.buckets[ind].Head = h.buckets[ind].Head.Next
		h.buckets[ind].Head.Previous = nil
		h.length--
		*h.buckets[ind].Len()--

		h.lock.Unlock()
		return nil
	}
	for head.Next != nil {
		if head.Data.(string) == word {
			h.lock.Lock()
			head.Previous.Next, head.Next.Previous = head.Next, head.Previous
			h.length--
			*h.buckets[ind].Len()--
			h.lock.Unlock()
			return nil
		}
		head = head.Next
	}
	if h.buckets[ind].Tail.Data == word {
		h.lock.Lock()
		h.buckets[ind].Tail = h.buckets[ind].Tail.Previous
		h.buckets[ind].Tail.Next = nil
		h.length--
		*h.buckets[ind].Len()--
		h.lock.Unlock()
		return nil
	}

	return fmt.Errorf("word '%s' couldn't be found", word)
}

//hash is simple hashing function that returns index according to its size.
func (h *hashTable) hash(word string) int {
	total := 0
	for _, l := range word {
		total += int(l)
	}
	return total % h.size
}

//Length returns total count of words that was exist in hash table
func (h *hashTable) Length() int {
	return h.length
}
