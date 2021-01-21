package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/supernovas/libnova"
)

func main() {
	tree := libnova.NewBinarySearchTree("Example Tree")
	tree.Insert(&libnova.NodeBinary{
		Node: libnova.Node{
			Key:   "Example Key",
			Value: "Some amazing value",
		},
	})
	//fmt.Println(tree.String())
	result := tree.Search("Example Key")
	if result == nil {
		fmt.Println("Error: nil result")
		os.Exit(1)
	}
	fmt.Println(result.Value)
}
