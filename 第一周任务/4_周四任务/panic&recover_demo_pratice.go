package main

import "fmt"

// 模拟一个极其危险的操作
func dangerousOperation() {
	// 【核心机制】：defer 是 Go 语言里的“延迟执行”。
	// 无论这个函数是正常结束，还是因为 panic 崩溃，defer 里的代码都一定会最后执行。
	defer func() {
		// 尝试使用 recover() 捕获可能发生的 panic
		if r := recover(); r != nil {
			fmt.Println("🚑 成功拦截到恐慌 (panic)，程序免于崩溃！错误原因是:", r)
		}
	}()

	fmt.Println("⚡ 危险操作开始...")

	// 突然遭遇致命打击，触发 panic！
	// 就像 Python 里的 raise Exception("数据库连接彻底断开！")
	panic("数据库连接彻底断开！")

	// 这一行代码永远不会被执行，因为上面已经 panic 了
	fmt.Println("✅ 危险操作顺利结束")
}

func main() {
	fmt.Println("🚀 主程序启动")

	dangerousOperation()

	// 如果没有 recover 拦截，程序会在 dangerousOperation 里直接挂掉，这句话就打印不出来了。
	// 但因为我们拦截成功了，主程序依然可以从容不迫地继续往下走。
	fmt.Println("🎉 主程序正常结束，没有受到影响。")
}
