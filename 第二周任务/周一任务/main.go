package main

import (
	"fmt"
	"os"
	"path/filepath" // 专门处理路径，解决 Windows(\) 和 Linux(/) 斜杠不通用的问题
	"sync"          // 核心包：提供同步原语，比如 WaitGroup
	"time"          // 核心包：处理时间和计时
)

func main() {
	// --- 第一阶段：环境准备 (串行执行) ---
	tempDir := "./test_files"
	os.MkdirAll(tempDir, 0755) // 创建文件夹，0755 是 Unix 系统的标准权限

	fmt.Println("开始生成 100 个测试文件...")
	for i := 0; i < 100; i++ {
		// Sprintf 的 'f' 代表 format。它不打印，而是返回一个拼接好的字符串
		fileName := filepath.Join(tempDir, fmt.Sprintf("file_%d.txt", i))
		content := fmt.Sprintf("这是第 %d 个文件的内容", i)
		// 将内容写成文件。[]byte(content) 是将字符串转成机器能识别的二进制字节流
		os.WriteFile(fileName, []byte(content), 0644)
	}
	fmt.Println("文件生成完毕！\n--- 开始并发读取 ---")

	// --- 第二阶段：并发核心控制 ---

	// wg 就像一个“倒计时计数器”。
	// 它的作用是告诉主程序：还有几个分身没回来，先别结束！
	var wg sync.WaitGroup

	startTime := time.Now() // 记录此刻时间

	for i := 0; i < 100; i++ {
		// 【关键 1】：Add(1) 必须在 go 关键字之前调用。
		// 这相当于在登记簿上写下：“又派出一个分身，目前总计任务数 +1”。
		wg.Add(1)

		// 【关键 2】：go 关键字会立即启动一个新的协程 (Goroutine) 。
		// 它不会等待函数执行完，而是直接让循环跳到下一次 i++。
		go func(id int) {
			// 【关键 3】：defer 会在函数执行结束（无论是成功还是报错）时最后执行。
			// 调用 Done() 相当于在登记簿上把任务数 -1。
			defer wg.Done()

			fileName := filepath.Join(tempDir, fmt.Sprintf("file_%d.txt", id))

			// 模拟耗时读取
			data, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Printf("读取失败: %v\n", err)
				return
			}

			// 只打印部分结果，避免刷屏
			if id < 5 || id == 99 {
				fmt.Printf("协程 %d 读取内容: %s\n", id, string(data))
			}
		}(i) // 此处的 (i) 是把当前的循环变量 i 传进匿名函数，复制给参数 id
	}

	// --- 第三阶段：收尾等待 ---

	// 【关键 4】：Wait() 会死死守在这里，直到 wg 的计数器归零。
	// 如果不写这一行，主程序执行完循环就会瞬间退出，你可能什么都看不到。
	wg.Wait()

	fmt.Printf("--- 并发读取完毕 ---\n总共耗时: %v\n", time.Since(startTime))

	// 清理现场
	os.RemoveAll(tempDir)
	fmt.Println("测试文件已清理。")
}
