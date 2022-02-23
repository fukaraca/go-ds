package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	obj := New()
	assert.Equal(t, nil, obj.Pop(), "must return nil for now")
	assert.Equal(t, nil, obj.Top(), "must return nil for now")
	obj.Push(new(interface{}))
	assert.Equal(t, 1, obj.Len(), "stack length must be 1")
	obj.stack[0] = 5
	assert.Equal(t, 5, obj.Top(), "top of the stack must b 5")
	assert.Equal(t, 5, obj.Pop(), "popped element must be 5")
	assert.Equal(t, true, obj.IsEmpty(), "stack must be empty now")
}
