package main

import "fmt"

// Node is a node of tree
type Node struct {
	key      string
	children map[rune]*Node
}

// NewTrie is creating a root node
func NewTrie() *Node {
	return &Node{
		key:      "",
		children: make(map[rune]*Node),
	}
}

// Insert is inserting a word to tree
func (n *Node) Insert(word string) {
	runes := []rune(word)
	curNode := n
	for _, r := range runes {
		if nextNode, exist := curNode.children[r]; exist {
			curNode = nextNode
		} else {
			curNode.children[r] = &Node{
				key:      string(r),
				children: make(map[rune]*Node),
			}
		}
	}
}

// Search is searching a word from a tree
func (n *Node) Search(word string) bool {
	if len(n.key) == 0 && len(n.children) == 0 {
		return false
	}

	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, exist := curNode.children[r]; exist {
			curNode = nextNode
		} else {
			return false
		}
	}

	return true
}

func main() {
	t := NewTrie()

	t.Insert("w")
	t.Insert("wo")
	t.Insert("wor")
	t.Insert("worl")
	t.Insert("world")

	fmt.Println(t)
	fmt.Println(t.Search("world"))
}
