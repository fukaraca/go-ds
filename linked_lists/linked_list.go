package linked_lists

import (
	"errors"
	"sync"
)

// doubly LinkedList
type LinkedList struct {
	Head *Node
	Tail *Node
	len  int
	lock *sync.Mutex
}

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}

//NewLinkedList returns a new linked list
func New() *LinkedList {
	nev := LinkedList{}
	nev.lock = &sync.Mutex{}
	return &nev
}

//InsertAtTail inserts a node to tail with the given data and returns the node
func (l *LinkedList) InsertAtTail(data interface{}) {
	if l.len == 0 {
		l.InsertAtHead(data)
		return
	}
	//if it is not first node in the list
	l.lock.Lock()
	neo := &Node{
		Data:     data,
		Next:     nil,
		Previous: l.Tail,
	}
	olderTail := l.Tail
	olderTail.Next = neo
	l.Tail = neo
	l.len++
	l.lock.Unlock()

}

//InsertAtHead inserts a node to head with given data
func (l *LinkedList) InsertAtHead(data interface{}) {
	if l.len == 0 {
		l.lock.Lock()
		neo := &Node{
			Data:     data,
			Next:     nil,
			Previous: nil,
		}
		l.Head = neo
		l.Tail = neo
		l.len++
		l.lock.Unlock()
		return
	}
	//if it is not first node
	l.lock.Lock()
	neo := &Node{
		Data:     data,
		Next:     l.Head,
		Previous: nil,
	}
	olderHead := l.Head
	olderHead.Previous = neo
	l.Head = neo
	l.len++
	l.lock.Unlock()
}

//Delete given item from the list
func (l *LinkedList) Delete(nood *Node) {
	if l.len != 0 {
		l.lock.Lock()
		defer l.lock.Unlock()
		//for the sake of for loops functionality, we need to be sure head is not node to be deleted
		if l.Head == nood {
			l.Head = l.Head.Next
			l.Head.Previous = nil
			l.len--
			return
		}
		temp := l.Head
		for temp.Next != nil {
			if temp == nood {
				temp.Previous.Next, temp.Next.Previous = temp.Next, temp.Previous
				l.len--
				return
			}
			temp = temp.Next
		}
		if l.Tail == nood {
			l.Tail = l.Tail.Previous
			l.Tail.Next = nil
			l.len--
			return
		}

	}

}

//DeleteAtHead deletes first element from the list
func (l *LinkedList) DeleteAtHead() {
	if l.len > 0 {
		l.lock.Lock()
		if l.len > 1 {
			l.Head.Next.Previous = nil
		}
		l.Head = l.Head.Next
		l.len--
		l.lock.Unlock()
	}
}

//Search returns the given element index from the list if found any. If  there is no match, it returns -1
func (l *LinkedList) Search(elem *Node) (int, error) {
	if l.len == 0 {
		return -1, errors.New("list is empty")
	}
	temp := l.Head
	for ind := 0; temp != nil; ind++ {
		if temp == elem {
			return ind, nil
		}
		temp = temp.Next
	}
	return -1, errors.New("there is no match")
}

//Get brings ind indexed element from list and error if occurs any
func (l *LinkedList) Get(ind int) (*Node, error) {
	if l.len == 0 {
		return nil, errors.New("list is empty")
	}
	if ind >= l.len {
		return nil, errors.New("index out of range")
	}
	temp := l.Head
	for i := 0; temp != nil; i++ {
		if i == ind {
			break
		}
		temp = temp.Next
	}
	return temp, nil
}

//IsEmpty returns true if list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

//Len returns node count of the list
func (l *LinkedList) Len() *int {
	return &l.len //pointer is for hash_table compatibility
}
