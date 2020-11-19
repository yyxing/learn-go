package tree

import "fmt"

type Node struct {
	Value        int
	Left, Right  *Node
	channelValue chan int
}

func (node Node) Print() {
	//fmt.Println("node value: ", node.Value)
	value := <-node.channelValue
	fmt.Println("node value: ", value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil cannot set value")
		return
	}
	node.Value = value
}
