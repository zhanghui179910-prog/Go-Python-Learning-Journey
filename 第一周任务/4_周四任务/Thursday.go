package main

import (
	"encoding/json"
	"fmt"
	"os" // 【新增】用于处理真正的文件读写
	"strings"
)

// ==========================================
// 1. 定义接口 (Interface)
// ==========================================
type DataExporter interface {
	Export(data any) error
}

// ==========================================
// 2. 定义结构体 (Structs)
// ==========================================
type JSONExporter struct {
	FileName string
}

type CSVExporter struct {
	FileName string
}

// ==========================================
// 3. 实现接口方法 (Duck Typing)
// ==========================================

func (j JSONExporter) Export(data any) error {
	fmt.Printf("准备将数据导出到 JSON 文件: %s\n", j.FileName)

	// 【修改】使用 MarshalIndent 让生成的 JSON 带有缩进，文件看着更清晰
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %v", err)
	}

	// 【核心新增：真实文件写入】
	// 使用 os.WriteFile 将字节流写入硬盘。0644 是常见的文件读写权限。
	err = os.WriteFile(j.FileName, bytes, 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("✅ 成功生成并写入 JSON 内容到文件: %s\n", j.FileName)
	return nil
}

func (c CSVExporter) Export(data any) error {
	// 【核心新增：触发 panic】
	// 假设如果文件名为空，这是一个绝对不允许发生的致命错误
	if c.FileName == "" {
		panic("致命错误：CSV 导出器的文件名不能为空！")
	}

	fmt.Printf("准备将数据导出到 CSV 文件: %s\n", c.FileName)

	switch v := data.(type) {
	case map[string]string:
		fmt.Printf("✅ 成功处理 CSV 数据，包含 %d 条记录\n", len(v))
		return nil
	default:
		return fmt.Errorf("CSV 导出失败: 不支持的数据类型 %T", v)
	}
}

// ==========================================
// 4. 见证奇迹的时刻 (Main)
// ==========================================
func main() {
	sampleData := map[string]string{
		"User1": "Alice",
		"User2": "Bob",
	}

	// 实例化
	jsonExp := JSONExporter{FileName: "users.json"}
	// 【注意这里】：我故意把 CSV 的文件名设为空，为了触发下面的 panic！
	csvExp := CSVExporter{FileName: ""}

	exporters := []DataExporter{jsonExp, csvExp}

	for _, exp := range exporters {
		// 【核心新增：使用匿名函数包裹 defer 和 recover】
		// 为什么要加一层 func() {}？
		// 因为如果直接写在 main 里，发生 panic 时整个 main 函数就退出了。
		// 写在这个匿名的包裹函数里，即使当次循环崩溃了，拦截之后，for 循环还能继续执行下一个！
		func() {
			defer func() {
				if r := recover(); r != nil {
					// 成功拦截恐慌，程序免于崩溃
					fmt.Printf("🚑 触发了 panic 被成功拦截！抢救信息: %v\n", r)
				}
			}()

			// 执行具体的导出逻辑
			err := exp.Export(sampleData)

			// 常规的 error 处理
			if err != nil {
				fmt.Println("❌ 发生常规错误:", err)
			}
		}() // 记得加上 () 调用这个匿名函数

		fmt.Println(strings.Repeat("-", 30))
	}

	// 证明程序没有彻底死掉
	fmt.Println("🎉 主程序循环结束，顺利通关！")
}
