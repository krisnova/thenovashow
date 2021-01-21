package libnova

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
)

// BinarySearchTree is used to store
// data with effecient lookups
type BinarySearchTree struct {
	Name      string
	Root      *NodeBinary
	NodeCount int
	Mutex     sync.Mutex
}

// CalculateValue is used to calculate a unique integer
// value given a pre populated "Key" in a node
func (b *NodeBinary) CalculateValue() int {
	// Calculate MD5 sum and form an int result
	h := md5.New()
	h.Write([]byte(b.Key))
	md5 := hex.EncodeToString(h.Sum(nil))
	// Now that we have a idempotent hash, we can calculate an integer value
	bigInt := big.NewInt(0)
	bigInt.SetString(md5, 16) // Probably get away with base 2 or base 16
	i64 := bigInt.Int64()
	n := int(i64)
	// n is the int result of our arithmetic
	if n < 0 {
		n = n * -1 // Take the absolute value of n
	}
	b.ID = n
	return n
}

// NewBinarySearchTree will initialize a new binary search tree
func NewBinarySearchTree(name string) *BinarySearchTree {
	btree := &BinarySearchTree{
		Name: name,
	}
	return btree
}

// Insert is used to insert a new Node
// in the binary search tree
func (b *BinarySearchTree) Insert(node *NodeBinary) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	node.CalculateValue()
	if b.NodeCount < 1 {
		b.NodeCount++
		b.Root = node
		return
	}
	b.NodeCount++
	recursiveInsertNode(b.Root, node)
}

func recursiveInsertNode(root, node *NodeBinary) {
	// Replace
	if node.ID <= root.ID {
		if root.Left == nil {
			root.Left = node
		} else {
			recursiveInsertNode(root.Left, node)
		}
		// Right
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			recursiveInsertNode(root.Right, node)
		}
	}
}

// Search is used to search for a Node in the tree given it's key
func (b *BinarySearchTree) Search(key string) *NodeBinary {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	search := &NodeBinary{
		Node: Node{
			Key: key,
		},
	}
	search.CalculateValue()
	return recursiveSearch(search, b.Root)
}

func recursiveSearch(search, node *NodeBinary) *NodeBinary {
	if node == nil {
		return nil
	}
	if search.ID == node.ID && search.Key == node.Key {
		return node
	}
	if search.ID <= node.ID {
		// Left
		return recursiveSearch(search, node.Left)
	}
	if search.ID >= node.ID {
		// Right
		return recursiveSearch(search, node.Right)
	}
	return nil
}

// String will JSON marashall the tree
func (b *BinarySearchTree) String() (string, error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	jbytes, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("unable to marshal: %v", err)
	}
	return string(jbytes), nil
}

// IsBST can be used to validate that the BST
// is in fact a binary search tree
func (b *BinarySearchTree) IsBST() bool {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	//return iterativeCheck(b.Root)
	return recursiveCheck(b.Root)
}

func iterativeCheck(node *Node) bool {
	// Faster
	checking := true
	//processing := node
	for checking {
		// logic to operate on processing
		// left and right are greater than =
		// update processing
		// break
	}

	return true
}

func recursiveCheck(node *NodeBinary) bool {
	result := true
	if node.Left != nil {
		// We have a left Node
		// Left is always less than or equal
		if node.Left.ID > node.ID {
			return false
		}
		result = recursiveCheck(node.Left)
	}
	if node.Right != nil {
		// We have a right Node
		// Right is always greater than or equal
		if node.Right.ID < node.ID {
			return false
		}
		result = recursiveCheck(node.Right)
	}
	// Add more rules here
	return result
}
