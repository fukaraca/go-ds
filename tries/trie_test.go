package tries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrie(t *testing.T) {
	qw := NewTrie()
	qw.Delete("merhaba")

	//INITIAL cases
	assert.NotEqual(t, nil, qw.Delete("merhaba"), "it is expected to return 'not existed' error")
	assert.Equal(t, false, qw.Search("merhaba"), "ıt is expected to return false")
	assert.NotEqual(t, nil, qw.Insert(""), "it is not expected to return nil for meaningless '' input ")
	assert.Equal(t, 0, qw.wordCount, "it is expected to return 0 as word count")

	//INSERT some words

	assert.Equal(t, nil, qw.Insert("hi"), "error not expected")
	assert.Equal(t, nil, qw.Insert("hello"), "error not expected")
	qw.Insert("merhaba")
	qw.Insert("selam")
	qw.Insert("hallo")
	assert.Equal(t, 5, qw.WordCount(), "it is expected to return 5 as word count")

	//insert same word
	assert.NotEqual(t, nil, qw.Insert("hi"), "it is expected to return 'word is already exist 'error")

	//SEARCH
	assert.Equal(t, true, qw.Search("hi"), "it should return true")
	assert.Equal(t, false, qw.Search("Hi"), "false return expected since it is case sensitive")

	//DELETE

	//insert new data set
	assert.Equal(t, nil, qw.Insert("oregon"), "error not expected")
	assert.Equal(t, nil, qw.Insert("oreo"), "error not expected")
	assert.Equal(t, nil, qw.Insert("or"), "error not expected")
	assert.Equal(t, nil, qw.Insert("oregonal"), "error not expected")

	//deletion of longer word shouldn't cause shorters extinct
	assert.Equal(t, nil, qw.Delete("oregon"), "error not expected")
	assert.Equal(t, false, qw.Search("oregon"), "it should return false")
	assert.Equal(t, true, qw.Search("oreo"), "it should return true")
	assert.Equal(t, true, qw.Search("or"), "it should return true")
	assert.Equal(t, true, qw.Search("oregonal"), "it should return true")

	//lets insert again
	assert.Equal(t, nil, qw.Insert("oregon"), "it is not expected to return err")

	assert.Equal(t, nil, qw.Delete("or"), "error not expected")
	assert.Equal(t, false, qw.Search("or"), "it should return false")
	assert.Equal(t, true, qw.Search("oreo"), "it should return true")
	assert.Equal(t, true, qw.Search("oregonal"), "it should return true")

	//for monochained / unique words
	assert.Equal(t, nil, qw.Insert("eyjafjallajökull"), "it is not expected to return error")
	assert.Equal(t, nil, qw.Delete("eyjafjallajökull"), "it is not expected to return error")

	//for words that has sub words like or+chestra
	assert.Equal(t, nil, qw.Insert("or"), "error not expected")
	assert.Equal(t, nil, qw.Insert("orchestra"), "error not expected")

	assert.Equal(t, nil, qw.Delete("orchestra"), "it is not expected to return error")

}
