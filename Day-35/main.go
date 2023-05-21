package main

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) AddChild(value int) {
	newNode := &Node{Value: value}
	t.Root.Children = append(t.Root.Children, newNode)
}

func PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)

	for _, child := range node.Children {
		PreOrderTraversal(child)
	}
}

func main() {
	root := &Node{Value: 1}
	tree := &Tree{Root: root}

	child1 := &Node{Value: 2}
	child2 := &Node{Value: 3}
	tree.Root.Children = append(tree.Root.Children, child1, child2)

	tree.AddChild(4)
	tree.AddChild(5)

	child2.Children = append(child2.Children, &Node{Value: 6})
	child2.Children = append(child2.Children, &Node{Value: 7})

	fmt.Println("Pre-order traversal:")
	PreOrderTraversal(root)
}
