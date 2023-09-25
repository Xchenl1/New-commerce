package main

import (
	"E-commerce_system/snow"
	"fmt"
	"sync"
)

/**
雪花算法
	 41bit timestamp | 10 bit machineID : 5bit workerID 5bit dataCenterID ｜ 12 bit sequenceBits
最多使用69年
*/

var wg sync.WaitGroup

func main() {
	w := snow.NewWorker(5, 5)

	ch := make(chan uint64, 100)
	count := 100
	wg.Add(count)
	defer close(ch)
	//并发 count个goroutine 进行 snowFlake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			id, _ := w.NextID()
			ch <- id
		}()
	}
	wg.Wait()
	m := make(map[uint64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		fmt.Println(id)
		if ok {
			fmt.Printf("repeat id %d\n", id)
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", len(m), "snowflake ID Get successed!")
}
