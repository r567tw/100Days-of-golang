package main

// heap 是一種二元樹結構, 適合作排序、優先權等應用
import (
	"container/heap"
	"fmt"
)

// 定义堆中的元素结构体
type Item struct {
	value    string // 元素值
	priority int    // 优先级
}

// 定义堆类型
type MinHeap []Item

// 实现 heap.Interface 接口的 Len 方法
func (h MinHeap) Len() int {
	return len(h)
}

// 实现 heap.Interface 接口的 Less 方法
func (h MinHeap) Less(i, j int) bool {
	return h[i].priority < h[j].priority
}

// 实现 heap.Interface 接口的 Swap 方法
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现 heap.Interface 接口的 Push 方法
func (h *MinHeap) Push(x interface{}) {
	item := x.(Item)
	*h = append(*h, item)
}

// 实现 heap.Interface 接口的 Pop 方法
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func main() {
	h := &MinHeap{}

	// 添加元素到堆中
	heap.Push(h, Item{value: "A", priority: 3})
	heap.Push(h, Item{value: "B", priority: 2})
	heap.Push(h, Item{value: "C", priority: 4})
	heap.Push(h, Item{value: "D", priority: 1})

	// 输出堆中的元素
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.value, item.priority)
	}
}
