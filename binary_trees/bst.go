package binary_trees

import (
	"fmt"
	"sync"
)

//BsTree is a binary search tree type data structure
//For brevity all fields are exported
type binarySearchTree struct {
	Root   *Node
	Length int //total node count
	lock   *sync.Mutex
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

//NewBinarySearchTree creates and returns a new BST
func NewBinarySearchTree(rootValue int) *binarySearchTree {
	root := &Node{Value: rootValue}
	return &binarySearchTree{
		Root:   root,
		Length: 1,
		lock:   &sync.Mutex{},
	}
}

//Insert basicly inserts given value as a node to binary search tree
func (b *binarySearchTree) Insert(value int) error {
	_, ok := b.Search(value)
	if ok {
		return fmt.Errorf("a node with %d value is alredy exist", value)
	}

	var locate func(*Node, *Node, int)
	locate = func(root, parent *Node, value int) {
		if root != nil {
			if root.Value > value {
				locate(root.Left, root, value)
			} else {
				locate(root.Right, root, value)
			}
		} else {

			root = &Node{
				Parent: parent,
				Left:   nil,
				Right:  nil,
				Value:  value,
			}
			if parent.Value > value {
				parent.Left = root
			} else {
				parent.Right = root
			}
			b.Length++
		}

	}
	b.lock.Lock()
	locate(b.Root, nil, value)
	b.lock.Unlock()
	return nil
}

//Search looks up for given node id and brings it with a boolean if found any, otherwise returns nil, and false
func (b *binarySearchTree) Search(value int) (*Node, bool) {

	//ret, ok := &Node{}, false
	var searcher func(*Node, int) (*Node, bool)

	searcher = func(root *Node, val int) (*Node, bool) {
		ret, ok := &Node{}, false

		switch {
		case root == nil:
			return nil, false
		case root.Value == val:
			ret, ok = root, true
			break
		case root.Value < val:
			ret, ok = searcher(root.Right, val)
		case root.Value > val:
			ret, ok = searcher(root.Left, val)
		}
		return ret, ok
	}

	ret, ok := searcher(b.Root, value)
	if ok {
		return ret, ok
	}

	return nil, false
}

//Delete simply deletes given valued node from tree and replaces accordingly if it require
func (b *binarySearchTree) Delete(value int) error {
	node, ok := b.Search(value) //this is the NODE to be deleted
	if !ok {
		return fmt.Errorf("node with %d id couldn't be found", value)
	}
	successor := &Node{}
	var findSuccessor func(*Node)
	//find successor within rightmost children of  NODE and assign it to successor variable
	findSuccessor = func(initNode *Node) {
		if initNode.Left == nil {
			successor = initNode
			return
		}
		findSuccessor(initNode.Left)
	}
	b.lock.Lock()
	if node.Left == nil && node.Right == nil { //in case of no child belongs to the node
		if node.Parent != nil {
			if node.Parent.Value > value {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		}
		node = &Node{
			Parent: nil,
			Left:   nil,
			Right:  nil,
			Value:  0,
		}
		b.Length--
		b.lock.Unlock()
	} else if node.Left != nil && node.Right == nil { //in case there is only one child of node which will be deleted
		//adopt the child to grandparent
		if node.Parent != nil {
			if node.Parent.Value > node.Value {
				node.Parent.Left, node.Left.Parent = node.Left, node.Parent
			} else {
				node.Parent.Right, node.Left.Parent = node.Left, node.Parent
			}
		} else {
			node.Left.Parent = nil
		}
		node = node.Left
		b.Length--
		b.lock.Unlock()
	} else if node.Left == nil && node.Right != nil { //in case there is only one child of node which will be deleted
		//adopt the child to grandparent
		if node.Parent != nil {
			if node.Parent.Value > node.Value {
				node.Parent.Left, node.Right.Parent = node.Right, node.Parent
			} else {
				node.Parent.Right, node.Right.Parent = node.Right, node.Parent
			}
		} else {
			node.Right.Parent = nil
		}
		node = node.Right
		b.Length--
		b.lock.Unlock()
	} else { //in case, both right and left child exist
		findSuccessor(node.Right)
		if successor == node.Right { //in case first rightmost child is the successor
			//adopt the child to the grandparent
			if node.Parent != nil {
				if node.Parent.Value > node.Value {
					node.Parent.Left, node.Right.Parent = node.Right, node.Parent

				} else {
					node.Parent.Right, node.Right.Parent = node.Right, node.Parent
				}
			} else {
				node.Right.Parent = nil
			}
			temp := node.Left
			node = node.Right
			node.Left = temp
			node.Left.Parent = node

			b.Length--
			b.lock.Unlock()
		} else { //in case successor is not first rightmost child
			node.Value = successor.Value

			//b.Length--
			b.lock.Unlock()
			_ = deleteNode(successor, b)

		}

	}
	if value == b.Root.Value {
		b.Root = node
	}
	return nil
}

//delete is helper function for recursive purpose of Delete method of the tree. A little verbosity doesn't hurt.
func deleteNode(node *Node, b *binarySearchTree) error {
	value := node.Value

	b.lock.Lock()
	if node.Left == nil && node.Right == nil { //in case of no child belongs to the node
		if node.Parent != nil {
			if node.Parent.Value > value {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		}
		node = &Node{
			Parent: nil,
			Left:   nil,
			Right:  nil,
			Value:  0,
		}
		b.Length--
		b.lock.Unlock()
	} else if node.Left == nil && node.Right != nil { //in case there is only one child of node which will be deleted
		//adopt the child to grandparent
		if node.Parent != nil {
			node.Parent.Left, node.Right.Parent = node.Right, node.Parent
		}
		node = node.Right
		b.Length--
		b.lock.Unlock()
	}

	return nil
}

//InorderTraversal function travels the tree in order and appends values to a slice
func (b *binarySearchTree) InorderTraversal() []int {
	if b.Root.Left == nil && b.Root.Right == nil {
		return []int{}
	}
	var inorderList []int
	var travers func(*Node)

	travers = func(node *Node) {

		if node.Left != nil {
			travers(node.Left)
		}
		if node.Left == nil && node.Right == nil { //base case
			inorderList = append(inorderList, node.Value)
		}
		if node.Left != nil && node.Right == nil { //second case
			inorderList = append(inorderList, node.Value)
		}
		if node.Right != nil { //third case
			inorderList = append(inorderList, node.Value)
			travers(node.Right)
		}
	}
	travers(b.Root)
	return inorderList
}

//Height function returns BST's height
func (b *binarySearchTree) Height() int {
	if b.Root.Left == nil && b.Root.Right == nil {
		return 1
	}
	height := 0
	var travers func(*Node)

	travers = func(node *Node) {

		if node.Left != nil {
			travers(node.Left)
		}
		if node.Left == nil && node.Right == nil { //base case
			ret := b.CountHeight(node.Value)
			if ret > height {
				height = ret
			}
		}
		if node.Left != nil && node.Right == nil { //second case
		}
		if node.Right != nil { //third case
			travers(node.Right)
		}
	}
	travers(b.Root)
	return height

}

//CountHeight is helper function for Height()
func (b *binarySearchTree) CountHeight(value int) int {

	counter := 0
	var searcher func(*Node, int) int

	searcher = func(root *Node, val int) int {

		switch {
		case root == nil:
			return counter
		case root.Value == val:
			break
		case root.Value < val:
			counter++
			counter = searcher(root.Right, val)
		case root.Value > val:
			counter++
			counter = searcher(root.Left, val)
		}
		return counter
	}

	searcher(b.Root, value)

	return counter
}
