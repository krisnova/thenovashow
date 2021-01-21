package libnova

type NodeBinary struct {
	Node
	RBinary
}

type NodeNeighbor struct {
	Node
	RNeighbors
}

// R [relationships]
//
// Use the "R" prefix to denote a type of data
// structure relationship

type RNeighbors struct {
	Neighbors []*NodeNeighbor
}

type RBinary struct {
	Left  *NodeBinary
	Right *NodeBinary
}

// Node
//
// The base type for all Nodes
// Note: this does not contain relationships

type Node struct {
	ID    int    // Index
	Key   string // Name
	Value interface{}
}
