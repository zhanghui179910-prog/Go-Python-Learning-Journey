package main

import (
	"fmt"
	"sync"
	"time" // 引入 time 包用于模拟耗时
)

func main() {
	// 创建一个有缓冲的 Channel，容量为 5
	numbers := make(chan int, 5)

	// 使用 WaitGroup 等待协程结束
	var wg sync.WaitGroup

	// 1. 生产者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 生成 1-15 的数字发入通道 [cite: 2]
		for i := 1; i <= 15; i++ {
			fmt.Printf("厨师做好了第 %d 道菜，准备放上保温架...\n", i)
			numbers <- i // 如果保温架满了，代码会卡在这里
			fmt.Printf(">>> 成功！第 %d 道菜已放上保温架\n", i)
		}
		// 发送方负责关闭通道 [cite: 2]
		close(numbers)
		fmt.Println("=== 厨师发完所有菜，宣告打烊 ===")
	}()

	// 2. 消费者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 使用 range 遍历通道 [cite: 2]
		for num := range numbers {
			fmt.Printf("---- 服务员端走数字: %d, 计算平方: %d\n", num, num*num)

			// 故意让服务员每处理一个数据就睡 1 秒，模拟耗时操作
			time.Sleep(1 * time.Second)
		}
	}()

	// 阻塞主协程，直到生产者和消费者都完成工作
	wg.Wait()
	fmt.Println("任务执行完毕，主程序安全退出。")
}
