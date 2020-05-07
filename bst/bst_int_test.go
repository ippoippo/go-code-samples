package bst

import (
	//"fmt"
	"testing"

	"github.com/ippoippo/go-core-lib/structures/treenode"
)

func TestInsert(t *testing.T) {
	var nilNode *treenode.Int
	l1 := []int{
		1,
	}

	l2134 := []int{2, 1, 3, 4}

	l2143 := []int{2, 1, 4, 3}

	l1234 := []int{1, 2, 3, 4}

	tests := []struct {
		vals     []int // List of values to insert into tree
		method   InsertMethod
		expected string // Expected result
		desc     string // Description of testcase
	}{
		// Recursive
		{l1, InsertMethodRecursive, "(1)", "One node - Recursive"},
		{l2134, InsertMethodRecursive, "((1) 2 (3 (4)))", "2,1,3,4 - Recursive"},
		{l2143, InsertMethodRecursive, "((1) 2 ((3) 4))", "2,1,4,3 - Recursive"},
		{l1234, InsertMethodRecursive, "(1 (2 (3 (4))))", "1,2,3,4 - Recursive"},

		// Iterative
		{l1, InsertMethodNonRecursive, "(1)", "One node - Iterative"},
		{l2134, InsertMethodNonRecursive, "((1) 2 (3 (4)))", "2,1,3,4 - Iterative"},
		{l2143, InsertMethodNonRecursive, "((1) 2 ((3) 4))", "2,1,4,3 - Iterative"},
		{l1234, InsertMethodNonRecursive, "(1 (2 (3 (4))))", "1,2,3,4 - Iterative"},
	}

	for _, test := range tests {
		tree := NewIntTree(nilNode)

		// Insert as per sequence defined
		for _, val := range test.vals {
			tree.Insert(val, "", test.method)
		}

		// Check the result
		actual := tree.String()
		if actual != test.expected {
			t.Errorf("[%s]:String() = Actual:[%q], Expected:[%q]", test.desc, actual, test.expected)
		}
	}
}

func TestSearch(t *testing.T) {
	var nilNode *treenode.Int
	l1 := []int{
		1,
	}

	l2143 := []int{2, 1, 4, 3}

	l1234 := []int{1, 2, 3, 4}

	tests := []struct {
		vals     []int // List of values to insert into tree
		param    int
		method   SearchMethod
		expected string // Expected result
		desc     string // Description of testcase
	}{
		// Recursive
		{l1, 1, SearchMethodRecursive, "(1)", "One node - Recursive - Search OK"},
		{l1, 2, SearchMethodRecursive, "()", "One node - Recursive - Search not found"},
		{l2143, 3, SearchMethodRecursive, "(3)", "2,1,4,3 - Recursive - Search OK"},
		{l2143, 5, SearchMethodRecursive, "()", "2,1,4,3 - Recursive - Search not found"},
		{l1234, 3, SearchMethodRecursive, "(3 (4))", "1,2,3,4 - Recursive - Search OK"},
		{l1234, 5, SearchMethodRecursive, "()", "1,2,3,4 - Recursive - Search not found"},

		// Iterative
		{l1, 1, SearchMethodNonRecursive, "(1)", "One node - Recursive - Search OK"},
		{l1, 2, SearchMethodNonRecursive, "()", "One node - Recursive - Search not found"},
		{l2143, 3, SearchMethodNonRecursive, "(3)", "2,1,4,3 - Recursive - Search OK"},
		{l2143, 5, SearchMethodNonRecursive, "()", "2,1,4,3 - Recursive - Search not found"},
		{l1234, 3, SearchMethodNonRecursive, "(3 (4))", "1,2,3,4 - Recursive - Search OK"},
		{l1234, 5, SearchMethodNonRecursive, "()", "1,2,3,4 - Recursive - Search not found"},
	}

	for _, test := range tests {
		tree := NewIntTree(nilNode)

		// Insert as per sequence defined
		for _, val := range test.vals {
			tree.Insert(val, "", InsertMethodRecursive)
		}

		// Check the result
		actual := tree.Search(test.param, test.method).String()
		if actual != test.expected {
			t.Errorf("[%s]:String() = Actual:[%q], Expected:[%q]", test.desc, actual, test.expected)
		}
	}
}
