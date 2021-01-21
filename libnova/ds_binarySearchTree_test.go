package libnova

import "testing"

// TestBSTHappy will check to make sure
// we have a happy binary search tree
func TestBSTHappy(t *testing.T) {
	tree := NewPopulatedBTree()
	if !tree.IsBST() {
		t.Errorf("Failure: isnotbst")
	}
}

func TestBSTInvalidHappy(t *testing.T) {
	tree := NewInvalidPopulatedBTree()
	if tree.IsBST() {
		t.Errorf("Failure: isbst")
	}
}

func TestBSTInsertHappy(t *testing.T) {
	btree := NewPopulatedBTree()
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "A large number",
			Value: 11234521341,
		},
	})
	if !btree.IsBST() {
		t.Errorf("Insert created invalid binary search tree")
	}
}
