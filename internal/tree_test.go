package internal

import "testing"

func TestTree(t *testing.T) {
	tree := NewTreeNode()
	tree.AddWord("testing")

	if !tree.Check("test") {
		t.Error("Expected tree to contain prefix \" test\".")
	}

	if !tree.Check("") {
		t.Error("Expected tree to contain empty string.")
	}

	if tree.Check("foo") {
		t.Error("Expected tree not to contain prefix \"foo\".")
	}
}
