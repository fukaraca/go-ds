package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	qw := New() //naming was chosen for brevity

	gotInt, _ := qw.Search(new(Node))
	if gotInt >= 0 {
		t.Errorf("wanted -1 but get %d", gotInt)
	}
	assert.Equal(t, true, qw.IsEmpty(), "it must return true at this moment")
	_, err := qw.Get(5)
	assert.NotEqual(t, nil, err, "it is expected return nil at this moment")
	qw.InsertAtHead("1")
	qw.InsertAtHead("2")
	qw.InsertAtHead("3")
	qw.InsertAtHead("4")
	qw.InsertAtHead("5")
	qw.InsertAtTail("13")
	qw.InsertAtTail("14")
	assert.Equal(t, 7, *qw.Len(), "it is expected to return 7")
	gotNode, _ := qw.Get(4)
	assert.Equal(t, "1", gotNode.Data.(string), "it is expected to return '1'")
	_, err = qw.Get(100)
	assert.NotEqual(t, nil, err, "it is expected to return nil due to index out of range error")
	gotNode, _ = qw.Get(3)
	gotInt, _ = qw.Search(gotNode)
	assert.Equal(t, 3, gotInt, "it is expected to get 3 for gotNode param")
	qw.Delete(gotNode)
	gotInt, _ = qw.Search(gotNode)
	assert.Equal(t, -1, gotInt, "it is expected to return -1 due to recent Delete method call")
	wantHead := qw.Head
	qw.DeleteAtHead()
	assert.NotEqual(t, wantHead, qw.Head, "it is not expected to see older head due to recent deletion")
	wantHead = qw.Head
	qw.Delete(wantHead)
	assert.Equal(t, "3", qw.Head.Data.(string), "it is expected to get string 3 data for head ")
	wantTail := qw.Tail
	qw.Delete(wantTail)
	assert.Equal(t, "13", qw.Tail.Data.(string), "it is expected to get string 13 data for tail")
	//to achieve %100 coverage
	for *qw.Len() > 0 {
		qw.DeleteAtHead()
	}
	qw.InsertAtTail("insert to tail")
	assert.Equal(t, "insert to tail", qw.Head.Data.(string), "it is expected to get same string as head and tail is same at the beginning")

}
