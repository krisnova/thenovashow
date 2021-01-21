package libnova

// NewPopulatedBTree Will return a populated a btree
func NewPopulatedBTree() *BinarySearchTree {
	btree := NewBinarySearchTree("Nóva")
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "seventy seven",
			Value: 77,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "sixty nine",
			Value: 69,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "four twenty",
			Value: 420,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "seven",
			Value: 7,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "eight",
			Value: 8,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "eighty eight",
			Value: 88,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "fourteen",
			Value: 14,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "fifty six",
			Value: 56,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "six six six",
			Value: 666,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "twenty seven",
			Value: 27,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "fourty two",
			Value: 42,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "four zero four",
			Value: 404,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "minus one",
			Value: -1,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "one thousand three hundred thirty seven",
			Value: 1337,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "fourteen fourty eight",
			Value: 1448,
		},
	})
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "seventeen eighty nine",
			Value: 1789,
		},
	})
	return btree
}

// NewPopulatedBTree Will return a populated a btree
// This implementation used to work, but is now broken due to how we have 3 fields
// Key, Value, and Data
// This is in fact a BROKEN tree
func NewInvalidPopulatedBTree() *BinarySearchTree {
	btree := NewBinarySearchTree("Nóva")
	btree.Insert(&NodeBinary{
		Node: Node{
			Key:   "bad key",
			Value: 10,
		},
		RBinary: RBinary{
			Left: &NodeBinary{
				Node: Node{
					Key:   "Greater!",
					Value: 11,
				},
			},
			Right: &NodeBinary{
				Node: Node{
					Key:   "Less than!",
					Value: 9,
				},
			},
		},
	})
	return btree
}
