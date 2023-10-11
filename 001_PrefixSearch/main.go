package main

import (
	"bufio"
	"fmt"
	"os"
)

// größere Liste z.B. hier: https://www.mit.edu/~ecprice/wordlist.10000

func main() {
	tree := LoadWords("sample.txt")
	// tree := LoadWords("wordlist.10000.txt")

	// DumpTree(tree, 0)
	hits := SearchTree(tree, "abs")
	fmt.Println(hits)
}

func LoadWords(filename string) *WordNode {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	root := NewWordTree()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		word := reader.Text()
		chars := []rune(word)
		// fmt.Println(word, chars)
		insertInfoTree(root, chars)
	}
	return root
}

func DumpTree(tree *WordNode, indent int) {
	spaces := ""
	for i := 0; i < indent; i++ {
		spaces += ". "
	}
	for key, sub := range tree.children {
		fmt.Printf("%s %c\n", spaces, key)
		if sub != nil {
			DumpTree(sub, indent+1)
		}
	}
}

type WordNode struct {
	children    map[rune]*WordNode
	isEndOfWird bool
}

func NewWordTree() *WordNode {
	return &WordNode{
		children:    make(map[rune]*WordNode),
		isEndOfWird: false,
	}
}

func insertInfoTree(tree *WordNode, chars []rune) {
	if len(chars) > 0 {
		key := chars[0]
		if tree.children[key] == nil {
			tree.children[key] = NewWordTree()
		}
		insertInfoTree(tree.children[key], chars[1:])
	} else {
		tree.isEndOfWird = true
	}
}

func makeList(tree *WordNode, prefix string) []string {
	var list []string
	if tree.isEndOfWird {
		list = append(list, prefix)
	}
	for key, sub := range tree.children {
		subPrefix := prefix + string(key)
		list = append(list, makeList(sub, subPrefix)...)
	}
	return list
}

func searchPrefix(tree *WordNode, chars []rune, prefix string) []string {
	if len(chars) == 0 {
		return makeList(tree, prefix)
	}

	key := chars[0]
	fmt.Printf("Search into subtree for %c\n", key)
	if tree.children[key] == nil {
		// no matches
		fmt.Printf("No matches for %c\n", key)
		return nil
	}
	return searchPrefix(tree.children[key], chars[1:], prefix+string(key))
}

func SearchTree(tree *WordNode, word string) []string {
	chars := []rune(word)
	return searchPrefix(tree, chars, "")
}
