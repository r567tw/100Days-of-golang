package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done() // 告知 WaitGroup 當前 Goroutine 完成

	c.mutex.Lock()         // 取得互斥
	defer c.mutex.Unlock() // 確保在函式結束時釋放互斥鎖

	c.value++
}

func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.value
}

func main() {
	counter := Counter{}

	// 等待所有 Goroutine 完成
	wg := sync.WaitGroup{}
	// 啟動多個 Goroutine 來增加計數器的值
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter.Increment(&wg)
		}()
	}

	// 等待所有 Goroutine 完成
	wg.Wait()

	// 取得最終的計數器值
	fmt.Println(counter.GetValue()) // 預期輸出: 10
}

// Ref: https://ithelp.ithome.com.tw/articles/10239538
