package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func BFS(graph [][]int, start int) {
	visited := make([]bool, len(graph))
	queue := Queue{}
	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		fmt.Printf("%d ", node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	graph := [][]int{
		{1, 2},    // Node 0 is connected to nodes 1 and 2
		{0, 3, 4}, // Node 1 is connected to nodes 0, 3, and 4
		{0, 5},    // Node 2 is connected to nodes 0 and 5
		{1},       // Node 3 is connected to node 1
		{1},       // Node 4 is connected to node 1
		{2},       // Node 5 is connected to node 2
	}

	startNode := 0
	fmt.Println("BFS traversal starting from node", startNode)
	BFS(graph, startNode)
}
