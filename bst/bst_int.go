package bst

import (
	//"fmt"

	"github.com/ippoippo/go-core-lib/structures/treenode"
)

// InsertMethod represents the type of insert implementation
type InsertMethod int

const (
	// InsertMethodRecursive ...
	InsertMethodRecursive InsertMethod = iota
	// InsertMethodNonRecursive ...
	InsertMethodNonRecursive
)

// SearchMethod represents the type of search implementation
type SearchMethod int

const (
	// SearchMethodRecursive ...
	SearchMethodRecursive SearchMethod = iota
	// SearchMethodNonRecursive ...
	SearchMethodNonRecursive
)

// A IntTree is the structure for a Binary Search tree, operating on `Int` nodes
type IntTree struct {
	Root *treenode.Int
}

// NewIntTree creates a new tree, from an initial node
func NewIntTree(root *treenode.Int) *IntTree {
	return &IntTree{
		Root: root,
	}
}

// String returns full dump of the contents of the tree
func (t *IntTree) String() string {
	return t.Root.String()
}

// Insert will insert (value/label) into IntTree `t` using `method`
func (t *IntTree) Insert(value int, label string, method InsertMethod) {
	if method == InsertMethodRecursive {
		t.Root = insertIntRecursive(t.Root, value, label)
		return
	}

	// Non recursive implementation
	current := &t.Root
	if *current == nil {
		t.Root = treenode.NewIntRoot(value, label)
	} else {
		for *current != nil {
			if (*current).EqualToValue(value) {
				// Overwrite
				(*current).Value = value
				(*current).Label = label
				return
			}

			if (*current).MoreThanValue(value) { // current.Value > value
				current = &(*current).Left
			} else { // value > current.Value
				current = &(*current).Right
			}
		}
		*current = treenode.NewIntRoot(value, label)
	}
}

func insertIntRecursive(current *treenode.Int, value int, label string) *treenode.Int {
	if current == nil {
		return treenode.NewIntRoot(value, label)
	} else if current.EqualToValue(value) {
		// Overwrite
		current.Value = value
		current.Label = label
	} else if current.MoreThanValue(value) { // current.Value > value
		current.Left = insertIntRecursive(current.Left, value, label)
	} else { // current.Value < value
		current.Right = insertIntRecursive(current.Right, value, label)
	}
	return current
}

// Search will search based on value using `method`
func (t *IntTree) Search(value int, method SearchMethod) *treenode.Int {
	if method == SearchMethodRecursive {
		return searchRecursively(value, t.Root)
	}

	// Non recursive implementation
	current := t.Root
	for current != nil {
		if current.EqualToValue(value) {
			return current
		}
		if current.MoreThanValue(value) { // current.Value > value
			current = current.Left
		} else { // value > current.Value
			current = current.Right
		}
	}
	return current
}

func searchRecursively(value int, node *treenode.Int) *treenode.Int {
	if node == nil || node.EqualToValue(value) {
		return node
	}

	if node.MoreThanValue(value) {
		return searchRecursively(value, node.Left)
	}
	// else node.LessThanValue(value)
	return searchRecursively(value, node.Right)
}
