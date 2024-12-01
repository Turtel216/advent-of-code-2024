package utils

// TrieNode represents a node in the trie
type TrieNode struct {
	Children   map[rune]*TrieNode
	DigitValue int
}

// NewTrieNode creates and returns a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children:   make(map[rune]*TrieNode),
		DigitValue: -1,
	}
}

// Trie represents the trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new empty Trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// NewTrieWithWords creates a Trie and populates it with predefined words and values
func NewTrieWithWords(wordsValues map[string]int) *Trie {
	trie := NewTrie()
	for word, value := range wordsValues {
		trie.Insert(word, value)
	}
	return trie
}

// Insert adds a word and its corresponding value into the Trie
func (t *Trie) Insert(word string, value int) {
	currNode := t.root
	for _, char := range word {
		if _, exists := currNode.Children[char]; !exists {
			currNode.Children[char] = NewTrieNode()
		}
		currNode = currNode.Children[char]
	}
	currNode.DigitValue = value
}

// GetRootNode returns the root node of the Trie
func (t *Trie) GetRootNode() *TrieNode {
	return t.root
}

// GetNextNode returns the child node corresponding to a given character
func (t *Trie) GetNextNode(currNode *TrieNode, char rune) *TrieNode {
	return currNode.Children[char]
}
