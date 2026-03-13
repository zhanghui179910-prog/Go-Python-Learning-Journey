📝 Go 语言第二周 (周一) 知识点总结笔记：

## 1. 并发编程核心概念

### 1.1 Goroutine（协程）
- **定义**：Go 语言中的轻量级线程，由 Go 运行时管理
- **启动方式**：使用 `go` 关键字 + 函数调用，如 `go func() { ... }()`
- **特点**：创建成本低（约 2KB 初始栈），调度由 Go 运行时负责，非操作系统线程

### 1.2 sync.WaitGroup（等待组）
- **作用**：协调多个 Goroutine 的同步执行，等待所有 Goroutine 完成任务
- **三个核心方法**：
  1. `Add(delta int)`：增加等待的 Goroutine 数量（计数器 +delta）
  2. `Done()`：标记一个 Goroutine 完成任务（计数器 -1）
  3. `Wait()`：阻塞主程序，直到计数器归零
- **使用模式**：`Add()` 必须在启动 Goroutine 前调用，`Done()` 通常在 Goroutine 中使用 `defer` 确保执行

## 2. 文件与路径操作

### 2.1 目录创建
- `os.MkdirAll(path string, perm os.FileMode)`：递归创建目录
- **权限参数**：0755（Unix 标准权限，所有者可读写执行，其他用户只读执行）

### 2.2 文件读写
- `os.WriteFile(filename string, data []byte, perm os.FileMode)`：一次性写入文件
- `os.ReadFile(filename string) ([]byte, error)`：一次性读取整个文件
- **字符串转换**：`[]byte(content)` 将字符串转为字节切片

### 2.3 路径处理
- `filepath.Join(elem ...string)`：跨平台路径拼接，自动处理不同操作系统的路径分隔符
- **重要**：避免手动拼接路径（如 `tempDir + "/file.txt"`），防止 Windows/Linux 兼容性问题

## 3. 时间测量与性能评估

### 3.1 时间记录
- `time.Now()`：获取当前时间点
- `time.Since(t time.Time)`：计算从时间点 t 到现在的耗时
- **使用场景**：性能基准测试，并发任务耗时统计

## 4. 匿名函数与闭包

### 4.1 Goroutine 中的参数传递
- **问题**：直接在 Goroutine 中使用循环变量 `i` 会导致数据竞争（所有 Goroutine 可能读取到相同的值）
- **解决方案**：将循环变量作为参数传递给匿名函数
  ```go
  go func(id int) {
      // 使用 id 而非 i
  }(i)  // 传递当前 i 的副本
  ```

### 4.2 defer 关键字
- **作用**：延迟执行，在函数返回前执行清理操作
- **在并发中的应用**：确保 `wg.Done()` 无论函数成功或失败都会执行
  ```go
  defer wg.Done()  // 保证计数器减一
  ```

## 5. 并发编程中的关键难点

### 5.1 WaitGroup 的正确使用顺序
```go
// ❌ 错误示例：Add() 在 go 语句之后
go func() {
    wg.Add(1)  // 可能导致主程序在 Add 之前就 Wait()
    defer wg.Done()
    // ...
}()

// ✅ 正确示例：Add() 在 go 语句之前
wg.Add(1)  // 先增加计数器
go func() {
    defer wg.Done()
    // ...
}()
```

### 5.2 Goroutine 的变量捕获问题
```go
// ❌ 错误示例：所有 Goroutine 可能都读取到 i=100
for i := 0; i < 100; i++ {
    go func() {
        fmt.Println(i)  // 危险的闭包！
    }()
}

// ✅ 正确示例：传递参数副本
for i := 0; i < 100; i++ {
    go func(id int) {
        fmt.Println(id)  // 安全的副本
    }(i)
}
```

### 5.3 资源清理的确定性
- **问题**：并发程序可能提前退出，导致资源泄漏
- **解决方案**：使用 `defer` 确保清理操作
  ```go
  defer os.RemoveAll(tempDir)  // 确保临时目录被删除
  ```

## 6. 最佳实践总结

1. **提前规划并发数量**：在循环开始前确定需要启动的 Goroutine 总数
2. **使用带缓冲的通道**：如果 Goroutine 间需要通信，考虑使用通道（Channel）
3. **限制并发度**：避免无限制创建 Goroutine（使用工作池模式）
4. **错误处理**：Goroutine 内部的错误应该通过通道传回主程序
5. **性能监控**：使用 `time` 包测量实际性能提升

## 7. 今日实战项目回顾

### 项目目标
创建 100 个测试文件，使用 100 个 Goroutine 并发读取，测量并发 vs 串行性能差异

### 核心代码逻辑
1. **准备阶段**：串行创建 100 个文本文件
2. **并发阶段**：使用 WaitGroup 协调 100 个 Goroutine 并发读取
3. **收尾阶段**：等待所有 Goroutine 完成，清理临时文件

### 学习收获
- 理解了 Go 并发模型的基本原理
- 掌握了 sync.WaitGroup 的同步机制
- 学会了避免 Goroutine 中的常见陷阱
- 实践了文件操作和路径处理的跨平台写法

---

**注**：并发编程是 Go 的核心优势，也是学习的难点。今天的练习为后续学习通道（Channel）、互斥锁（Mutex）等高级并发特性打下了坚实基础。