package heaps

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestNewMaxHeap(t *testing.T) {
	qw := NewMaxHeap()

	//INITIALS
	assert.NotNil(t, qw.Delete(5), "it is expected to get an error")
	ok, _ := qw.Search(5)
	assert.Equal(t, false, ok, "it is expected to return false")
	count, _ := qw.SearchForLatest(5)
	assert.Equal(t, -1, count, "it is expected to return -1 for count since heap is empty")
	assert.Panics(t, func() {
		qw.Max()
	}, "panic is expected because of heap is empty")
	assert.Panics(t, func() {
		qw.Min()
	}, "panic is expected because of heap is empty")

	//INSERT
	// 97,46,37,12,3,7,31,6,9    https://media.geeksforgeeks.org/wp-content/cdn-uploads/yes.jpg
	testData := []int{97, 46, 37, 12, 3, 7, 31, 6, 9}
	for _, datum := range testData {
		qw.Insert(datum)
	}

	assert.Equal(t, 9, qw.Length(), "it should return 9 for length")
	assert.Equal(t, 3, qw.Min(), "it should return 3 for min")
	assert.Equal(t, 97, qw.Max(), "it should return 97 for max")
	//SEARCH

	//root
	ok, ind := qw.Search(97) //must return ok and 0 for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, ind)
	//child
	ok, ind = qw.Search(3) //must return ok and 4 for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 4, ind)
	//leaf,tail
	ok, ind = qw.Search(9) //must return ok and 8 for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 8, ind)
	//not existed case
	ok, ind = qw.Search(500) //must return ok and 8 for index
	assert.Equal(t, false, ok)
	assert.Equal(t, -1, ind)

	//DELETE
	//max
	err := qw.Delete(97)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 46, qw.Max(), "since 97 was deleted from heap new Max must be 46")
	//min
	err = qw.Delete(3)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 6, qw.Min(), "since 3 was deleted from heap new Min must be 6")
	assert.NotNil(t, qw.Delete(500), "it must throw an error since key is not even exist")

	qw1 := NewMaxHeap()
	rand.Seed(time.Now().UnixMilli())
	compSlice := []int{}

	for i := 0; i < 300; i++ {
		temp := rand.Intn(350)
		if i%10 == 0 {
			qw1.Insert(25) //in case of same key existance
		}
		compSlice = append(compSlice, temp)
		qw1.Insert(temp)
	}

	assert.Equal(t, 330, qw1.Length(), "it is expected to get 100 as count")
	for _, comps := range compSlice {
		count, ind = qw1.SearchForLatest(comps)
		if count == -1 || ind == -1 {
			t.Errorf(err.Error())
		}
	}

	count, ind = qw1.SearchForLatest(500) //search same item
	if count != -1 || ind != -1 {
		t.Errorf(err.Error())
	}

}

func TestNewMinHeap(t *testing.T) {
	qw := NewMinHeap()

	//INITIALS
	assert.NotNil(t, qw.Delete(5), "it is expected to get an error")
	ok, _ := qw.Search(5)
	assert.Equal(t, false, ok, "it is expected to return false")
	assert.Panics(t, func() {
		qw.Max()
	}, "panic is expected because of heap is empty")
	assert.Panics(t, func() {
		qw.Min()
	}, "panic is expected because of heap is empty")

	//INSERT
	// 1,2,3,17,19,36,7,25,100    https://upload.wikimedia.org/wikipedia/commons/6/69/Min-heap.png
	testData := []int{1, 2, 3, 17, 19, 36, 7, 25, 100}
	for _, datum := range testData {
		qw.Insert(datum)
	}

	assert.Equal(t, 9, qw.Length(), "it should return 9 for length")
	assert.Equal(t, 1, qw.Min(), "it should return 1 for min")
	assert.Equal(t, 100, qw.Max(), "it should return 100 for max")
	//SEARCH

	//root
	ok, ind := qw.Search(100) //must return ok and  for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 8, ind)
	//child
	ok, ind = qw.Search(17) //must return ok and 4 for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, ind)
	//leaf,tail
	ok, ind = qw.Search(1) //must return ok and 8 for index
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, ind)
	//not existed case
	ok, ind = qw.Search(500) //must return ok and 8 for index
	assert.Equal(t, false, ok)
	assert.Equal(t, -1, ind)

	//DELETE
	//max
	err := qw.Delete(100)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 36, qw.Max(), "since 100 was deleted from heap new Max must be 36")
	//min
	err = qw.Delete(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 2, qw.Min(), "since 1 was deleted from heap new Min must be 2")
	assert.NotNil(t, qw.Delete(500), "it must throw an error since key is not even exist")
	qw1 := NewMaxHeap()
	rand.Seed(time.Now().UnixMilli())
	compSlice := []int{}

	for i := 0; i < 100; i++ {
		temp := rand.Intn(200)
		compSlice = append(compSlice, temp)
		qw1.Insert(temp)
	}
	assert.Equal(t, 100, qw1.Length(), "it is expected to get 100 as count")
	for _, comps := range compSlice {
		ok, ind = qw1.Search(comps)
		if ok == false || ind == -1 {
			t.Errorf(err.Error(), comps)
		}
	}
}
