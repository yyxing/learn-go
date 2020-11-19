package main

import (
	"awesomeProject/func/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{Value: 0}
	root.Left.Right = &tree.Node{Value: 2}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = &tree.Node{Value: 4}
	channel := root.TraverseChannel()
	for node := range channel {
		fmt.Println("node value : ", node.Value)
	}
	//nodeCount := 0
	//root.TraverseFunc(func(node *tree.Node) {
	//	nodeCount++
	//})
	//fmt.Println("nodeCount: ", nodeCount)
}
