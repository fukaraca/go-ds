package tries

import (
	"fmt"
	"sync"
)

type node struct {
	children map[uint8]*node
	isEnding bool
}

type trie struct {
	root      *node
	lock      *sync.Mutex
	wordCount int
}

//NewTrie creates a new Trie and returns it
func NewTrie() *trie {
	tempChildren := make(map[uint8]*node)
	tempNode := &node{
		children: tempChildren,
		isEnding: false,
	}
	return &trie{
		root:      tempNode,
		lock:      &sync.Mutex{},
		wordCount: 0,
	}
}

//Insert adds the word to Trie and returns error if exist already
func (t *trie) Insert(word string) error {
	if word == "" {
		return fmt.Errorf("missing input")
	}
	root := t.root
	t.lock.Lock()
	defer t.lock.Unlock()
	for i := 0; i < len(word); i++ {
		//letter exist
		if child, ok := root.children[word[i]]; ok {
			root = child
		} else {
			//letter is not exist
			newChildren := make(map[uint8]*node)
			newNode := &node{
				children: newChildren,
				isEnding: false,
			}
			root.children[word[i]] = newNode
			root = root.children[word[i]]
		}
	}
	if root.isEnding {
		return fmt.Errorf("given word %s is already exist", word)
	}

	t.wordCount++
	root.isEnding = true
	return nil
}

//Search looks-up for given word and returns result and an error
func (t *trie) Search(word string) bool {
	root := t.root

	for i := 0; i < len(word); i++ {
		//letter exist
		if child, ok := root.children[word[i]]; ok {
			root = child
		} else {
			//letter is not exist
			return false
		}
	}
	if !root.isEnding {
		return false
	}
	return true
}

//WordCount simply returns the word count
func (t *trie) WordCount() int {
	return t.wordCount
}

//Delete deletes the word and return error if any
func (t *trie) Delete(word string) error {
	if !t.Search(word) {
		return fmt.Errorf("given word %s is not even exist", word)
	}

	isUnique := true
	root := t.root
	subWord := &node{}

	t.lock.Lock()
	defer t.lock.Unlock()

	for i := 0; i < len(word); i++ {
		//check if unique
		if root.children[word[i]].isEnding && i != len(word)-1 {
			isUnique = false
			//note where was the closest neighbor
			subWord = root.children[word[i]]
		}
		root = root.children[word[i]]
	}
	t.wordCount--
	//if chain still stores data, just unmark the leaf
	if len(root.children) > 0 {
		root.isEnding = false
		return nil
	}
	//if chain is not used by any other word then THE word, just orphan from first letter
	if isUnique {
		t.root.children[word[0]] = nil
		return nil
	}
	//last case, orphan from closest neighbor
	subWord.children = nil

	return nil
}
