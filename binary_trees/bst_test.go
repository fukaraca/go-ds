package binary_trees

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

//Coverage 98.4
func TestNewBinarySearchTree(t *testing.T) {
	qw := NewBinarySearchTree(5)
	assert.NotEqual(t, nil, qw.Insert(5), "it is expected to throw an error since 5 is already exist")
	_, ok := qw.Search(5)
	assert.Equal(t, true, ok, "it is expected to return true for 5")
	_, ok = qw.Search(6)
	assert.NotEqual(t, true, ok, "it is not expected to return true for 6")
	assert.NotEqual(t, nil, qw.Delete(6), "it is not expected to return nil err for non-existed node deletion")
	assert.Equal(t, nil, qw.Delete(5), "it is expected to return nil err")
	assert.Equal(t, []int{}, qw.InorderTraversal(), "it is expected to return for empty int slice")

	//create new BST
	qw1 := NewBinarySearchTree(50)
	//left branch
	_ = qw1.Insert(30)
	_ = qw1.Insert(35)
	_ = qw1.Insert(20)
	_ = qw1.Insert(10)
	_ = qw1.Insert(40)
	_ = qw1.Insert(45)
	_ = qw1.Insert(48)
	_ = qw1.Insert(46)
	_ = qw1.Insert(49)
	_ = qw1.Insert(12)
	_ = qw1.Insert(44)
	_ = qw1.Insert(43)
	_ = qw1.Insert(42)
	_ = qw1.Insert(41)
	//right branch
	_ = qw1.Insert(70)
	_ = qw1.Insert(60)
	_ = qw1.Insert(80)
	_ = qw1.Insert(90)
	_ = qw1.Insert(55)
	_ = qw1.Insert(52)
	_ = qw1.Insert(53)
	_ = qw1.Insert(85)
	_ = qw1.Insert(83)
	_ = qw1.Insert(84)
	_ = qw1.Insert(94)
	_ = qw1.Insert(99)

	_ = `                50                             
            30                                70              
         20     35                          60     80           
      10           40                    55           90        
         12           45              52           85     94     
                  44     48              53     83           99  
               43     46     49                    84           
            42                                               
         41                                                  
                                                            
` //visual was gotten from Visualize function
	expected := []int{10, 12, 20, 30, 35, 40, 41, 42, 43, 44, 45, 46, 48, 49, 50, 52, 53, 55, 60, 70,
		80, 83, 84, 85, 90, 94, 99}
	gOt := qw1.InorderTraversal()
	assert.Equal(t, expected, gOt, "it is expected to get: gOt list for inorder traversal")
	//height
	assert.Equal(t, 8, qw1.Height(), " it is expected to return 8 for height of the tree")
	//delete
	err := qw1.Delete(90)
	assert.Equal(t, nil, err, "it is expected to return nil err")
	err = qw1.Delete(qw1.Root.Value) //delete root
	assert.Equal(t, nil, err, "it is expected to return nil err")
	assert.Equal(t, 52, qw1.Root.Value, "new root value is 52 since former root vas 50")
	expected = []int{10, 12, 20, 30, 35, 40, 41, 42, 43, 44, 45, 46, 48, 49, 52, 53, 55, 60,
		70, 80, 83, 84, 85, 94, 99}
	gOt = qw1.InorderTraversal()
	assert.Equal(t, expected, gOt, "it is expected to get : gOt list...")

	for _, nodeVal := range expected {
		err = qw1.Delete(nodeVal)
		if err != nil {
			t.Errorf("expected nil but got %v for node %d", err, nodeVal)
		}
	}
	assert.Equal(t, true, qw1.Root.Right == nil && qw1.Root.Left == nil, "it is expected to get an empty BTS")

	//
	//create random nodes.
	qw2 := NewBinarySearchTree(2500000)
	rand.Seed(time.Now().UnixMilli())

	for i := 0; i < 99999; i++ {
		err = qw2.Insert(rand.Intn(5000000)) //we dont expect err for insertion since we checked with %100 coverage
		if err != nil {
			//we hit same number again.
			i-- //ignore first insertion luck
		}
	}
	orderedList := qw2.InorderTraversal()
	assert.Equal(t, 100000, len(orderedList), "we expect 100k count of node in the newly created list")
	//make list randomized
	tempMap := make(map[int]int)
	for i, item := range orderedList {
		tempMap[i] = item
	}

	//delete randomly to increase test case scenarios
	for _, val := range tempMap {
		err = qw2.Delete(val)
		if err != nil {
			t.Errorf("expected nil but got %v for node %d", err, val)
		}
	}
	assert.Equal(t, true, qw2.Root.Right == nil && qw2.Root.Left == nil, "it is expected to get an empty BTS")

}
