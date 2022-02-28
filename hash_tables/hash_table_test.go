package hash_tables

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"unicode"
)

func TestNewHashTable(t *testing.T) {
	qw := NewHashTable(10)

	//INITIAL
	ok, err := qw.Search("")
	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, true, ok)

	err = qw.Delete("")
	assert.NotEqual(t, nil, err)

	err = qw.Insert("")
	assert.NotEqual(t, nil, err)

	assert.Equal(t, 0, qw.Length())

	//INSERT
	// there other three until light paper... these 6 words produce same index=6 with the hash function for size 10
	//these words will be inserted into bucket indexed at 6
	assert.Equal(t, nil, qw.Insert("there"), "it is not expected to return an error")
	qw.Insert("other")
	qw.Insert("three")
	qw.Insert("until")
	qw.Insert("light")
	qw.Insert("paper")
	assert.NotEqual(t, nil, qw.Insert("there"), "it is expected to return an error for double inserting")

	ok, err = qw.Search("there") //check head for bucket 6
	assert.Equal(t, true, ok, "it is not expected to see false return")
	assert.Equal(t, nil, err, "it is not expected to get an error")

	ok, err = qw.Search("paper") //check tail
	assert.Equal(t, true, ok, "it is not expected to see false return")
	assert.Equal(t, nil, err, "it is not expected to get an error")

	assert.Equal(t, 6, qw.Length(), "it is expected to get 6 as length")

	ok, err = qw.Search("notExist")
	assert.NotEqual(t, nil, err, "an error expected")
	assert.Equal(t, false, ok, "key 'notExist' shouldn't be exist")

	//DELETE
	//delete from tail, head and amid

	//tail
	assert.Equal(t, nil, qw.Delete("paper"), "it is not expected to get an error")
	ok, err = qw.Search("paper") //check tail
	assert.Equal(t, false, ok, "it is expected to see false return")
	assert.NotEqual(t, nil, err, "it is  expected to get an error")

	//middle
	assert.Equal(t, nil, qw.Delete("three"), "it is not expected to get an error")
	ok, err = qw.Search("three") //check mid
	assert.Equal(t, false, ok, "it is expected to see false return")
	assert.NotEqual(t, nil, err, "it is  expected to get an error")

	//head
	assert.Equal(t, nil, qw.Delete("there"), "it is not expected to get an error")
	ok, err = qw.Search("there") //check mid
	assert.Equal(t, false, ok, "it is expected to see false return")
	assert.NotEqual(t, nil, err, "it is  expected to get an error")

	//delete non existed
	err = qw.Delete("notExisted")
	assert.NotEqual(t, nil, err, "an error must be thrown since the key is not exist")

	//delete others for next test case
	assert.Equal(t, nil, qw.Delete("light"))
	assert.Equal(t, nil, qw.Delete("other"))
	assert.Equal(t, nil, qw.Delete("until"))
	assert.Equal(t, 0, qw.length, "word length in the hash table must be zero now")

	//insert whole 5757 words from the wordle list
	list := Wordlist()

	assert.Equal(t, 5757, len(list), "it is expected 5757 words in the list")

	for _, word := range list {
		err = qw.Insert(word)
		if err != nil {
			t.Error(err)
		}
	}

	assert.Equal(t, 5757, qw.Length(), "it is expected to get 5757 words in the hash table")
	assert.NotEqual(t, nil, qw.Delete("notExistedKey"))

	//search all words and delete after
	for _, word := range list {
		ok, err = qw.Search(word)
		if err != nil {
			t.Error(err, word)
		}
		if !ok {
			t.Error(err, word)
		}

		err = qw.Delete(word)
		if err != nil {
			t.Error(err, word)
		}

	}

}

//5757 pieces 5 letters of words from wordle game
//check https://github.com/fukaraca/go-wordle-cli
func Wordlist() []string {
	file, err := os.OpenFile("wordList.txt", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("file couldn't be read")
		return nil
	}
	read, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	word := []byte{}
	words := []string{}
	for _, b := range read {
		if unicode.IsLetter(rune(b)) {
			word = append(word, b)
			if len(word) == 5 {
				words = append(words, string(word))
				word = []byte{}
			}
		}
	}

	return words
}
