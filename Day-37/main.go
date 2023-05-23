package main

import "fmt"

// 图结构
type Graph struct {
	vertices []*Vertex
}

// 顶点结构
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// 添加顶点
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// 添加边
func (g *Graph) AddEdge(from, to int) {
	// 检查顶点是否存在
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		// 检查边是否已存在
		if contains(fromVertex.adjacent, to) {
			err := fmt.Errorf("Edge (%v --> %v) already exists", from, to)
			fmt.Println(err.Error())
		} else {
			fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		}
	}
}

// 获取顶点
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

// 检查切片中是否包含特定元素
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

// 打印图结构
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %v:", v.key)
		for _, a := range v.adjacent {
			fmt.Printf(" %v", a.key)
		}
		fmt.Println()
	}
}

func main() {
	graph := Graph{}
	// 添加顶点
	for i := 0; i < 6; i++ {
		graph.AddVertex(i)
	}
	// 添加边
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)

	// 打印图结构
	graph.Print()
}
