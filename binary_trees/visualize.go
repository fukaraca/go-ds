package binary_trees

import (
	"fmt"
	"github.com/fatih/color"
)

//Visualize basicly displays the tree for testing purpose. There are issues related with concatenation for inner branches. It' is useful when inserted carefully.
func (b *binarySearchTree) Visualize(size int) string {
	var returnee string
	var travers func(*Node)
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	traveler := &struct { //traveler simply travels with InorderTraversal
		r     int
		c     int
		found bool
	}{}
	traveler.r = 0 //starting point of the root(head)
	traveler.c = size / 2

	travers = func(node *Node) {
		initialR := traveler.r
		initialC := traveler.c

		if node.Left != nil {
			traveler.c--
			traveler.r++
			travers(node.Left)
		}
		if node.Left == nil && node.Right == nil { //base case
			grid[traveler.r][traveler.c] = node.Value
			traveler.found = true

		}
		if node.Left != nil && node.Right == nil { //second case
			if traveler.found {
				traveler.r = initialR
				traveler.c = initialC
				traveler.found = false
			}
			grid[traveler.r][traveler.c] = node.Value
			traveler.found = true
		}
		if node.Right != nil { //third case
			if traveler.found {
				traveler.r = initialR
				traveler.c = initialC
				traveler.found = false
			}
			grid[traveler.r][traveler.c] = node.Value
			traveler.c++
			traveler.r++

			travers(node.Right)
		}
		//return old position
		traveler.r = initialR
		traveler.c = initialC
	}

	//repositioning starting point to left child
	traveler.r++
	traveler.c -= size / 2 //this is for averting concatenation of inner nodes.
	travers(b.Root.Left)
	//reposition to root
	traveler.r = 0
	traveler.c = size / 2
	grid[traveler.r][traveler.c] = b.Root.Value
	//reposition to right child
	traveler.r++
	traveler.c += size / 2 //this is for averting concatenation of inner nodes
	travers(b.Root.Right)
	//visualize the grid
	for i := range grid {
		for _, el := range grid[i] {
			if el > 0 {
				returnee += color.BlueString("%02d  ", el)
				//returnee += fmt.Sprintf("%02d  ", el)
			} else {
				returnee += color.BlackString("%s  ", " ")
				//returnee += fmt.Sprintf("%02s  ", " ")
			}

		}
		returnee += fmt.Sprintf("\n")
	}
	return returnee
}
