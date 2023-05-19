package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}

	queue.Enqueue("Jimmy")
	queue.Enqueue("Tom")
	queue.Enqueue("Alex")
	fmt.Println("length:", queue.Size())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println("length", queue.Size()) // 1
}
