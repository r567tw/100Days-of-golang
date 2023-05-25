package main

import "fmt"

type Node struct {
	key      int
	children []*Node
}

func NewNode(key int) *Node {
	return &Node{
		key:      key,
		children: []*Node{},
	}
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func DFS(node *Node,visited []Node) {
	fmt.Printf("%v ", node.key)
	visited = append(visited,*node)
	
	for _, child := range node.children {
		processed := false
		for _,visit := range visited{
			if visit.key == child.key {
				processed = true
				break
			}
		}
		if !processed {
			DFS(child,visited)
		}
	}
}

func main() {
	var visited []Node
	visited = make([]Node,0)
	// 建立樹結構
	root := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)

	root.AddChild(node2)
	root.AddChild(node3)
	node2.AddChild(node4)
	node3.AddChild(node5)

	fmt.Println("DFS Result:")
	DFS(root,visited)
	fmt.Println("")
}
