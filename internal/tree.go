package internal

// PrefixTree is an interface for checking if a string is a prefix of a valid word.
// It is used for pruning the search tree if no valid words could fit in a location.
type PrefixTree interface {
	AddWord(word string)
	Check(word string) bool
}

// TreeNode is an implementation of PrefixTree.
type TreeNode struct {
	Children map[byte]*TreeNode
}

// NewTreeNode is a constructor for a TreeNode.
func NewTreeNode() *TreeNode {
	t := new(TreeNode)
	t.Children = make(map[byte]*TreeNode)
	return t
}

// AddWord adds a new word to the prefix tree.
func (t *TreeNode) AddWord(word string) {
	if word == "" {
		return
	}

	if t.Children[word[0]] == nil {
		t.Children[word[0]] = NewTreeNode()
	}

	t.Children[word[0]].AddWord(word[1:])
}

// Check checks if the given string is a prefix of any valid word.
func (t *TreeNode) Check(word string) bool {
	if word == "" {
		return true
	}

	child, ok := t.Children[word[0]]
	if !ok {
		return false
	}

	return child.Check(word[1:])
}
