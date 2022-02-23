package queues

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	qw := New()
	assert.Equal(t, nil, qw.First(), "initial length must be 0")
	assert.Equal(t, nil, qw.Last(), "initial length must be 0")
	assert.Equal(t, true, qw.IsEmpty(), "initial length must be 0")
	qw.Enqueue("1")
	qw.Enqueue("2")
	qw.Enqueue("3")
	assert.Equal(t, "3", qw.Last(), "last elemet should be 3")
	assert.Equal(t, "1", qw.First(), "first element should 1")
	assert.Equal(t, 3, qw.Len(), "length of the queue mst be 3 now")
	qw.Dequeue()
	qw.Dequeue()
	qw.Dequeue()
	assert.Equal(t, 0, qw.Len(), "len must be 0 now")
}
