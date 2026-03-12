package main

import (
	"encoding/csv"  // 用于将数据写成逗号分隔的表格（Excel可读）
	"encoding/json" // Go语言的“翻译官”，负责 Go结构体 和 通用JSON文本 之间的互转
	"fmt"           // 负责终端的格式化输入和输出 (类似 Python 的 print)
	"os"            // 负责和操作系统打交道：读写硬盘文件、获取命令行参数、退出程序
	"strconv"       // 字符串转换包 (String Conversion)：负责把文字 "1" 变成数字 1
)

// Task 定义了我们在内存中操作的“标准便签”格式
// 这是你抛弃 Python 面向对象(Class)，拥抱 Go 结构体(Struct) 的核心
type Task struct {
	ID       int
	Name     string
	Status   string
	SubTasks []Task // 【核心知识点：结构体嵌套】切片里面装自己类型的元素，实现树状多层级
}

// memoryTasks 是程序运行时的“白板”，所有数据都在内存里
// 程序一关就会被清空，所以必须搭配下面的 fileName 存进硬盘
var memoryTasks = []Task{}

// fileName 定义了“保险柜”的名字。使用 const(常量) 防止后面手滑把文件名改错了
const fileName = "tasks.json"

// handleDelete 演示了 Go 语言中最经典的“切片删除法”
func handleDelete(id int) {
	found := false

	for i, task := range memoryTasks {
		if task.ID == id {
			// 【魔法操作：切片拼接】
			// memoryTasks[:i] 截取目标元素的前半段
			// memoryTasks[i+1:] 截取目标元素的后半段
			// ... (展开操作符) 把后半段打散，塞进前半段里，完美挤掉第 i 个元素
			memoryTasks = append(memoryTasks[:i], memoryTasks[i+1:]...)
			found = true
			fmt.Printf("任务 [%d] %s 已成功删除！\n", task.ID, task.Name)
			break // 找到了就停下，提高效率
		}
	}

	if !found {
		fmt.Printf("错误: 找不到 ID 为 %d 的任务。\n", id)
		return
	}

	saveTasks() // 内存改了，记得同步到硬盘保险柜
	handleList()
}

// handleDone 演示了如何修改切片里的数据
func handleDone(id int) {
	found := false

	for i, task := range memoryTasks {
		if task.ID == id {
			// 【避坑点】：必须通过索引 memoryTasks[i] 去修改原切片！
			// 不能写 task.Status = "已完成"，因为 task 只是当前循环里拿出来的一个“临时复印件”
			memoryTasks[i].Status = "已完成"
			found = true
			fmt.Printf("任务 [%d] %s 已标记为完成！\n", task.ID, task.Name)
			break
		}
	}

	if !found {
		fmt.Printf("错误: 找不到 ID 为 %d 的任务。\n", id)
		return
	}

	saveTasks()
	handleList()
}

// saveTasks 是“正向翻译官”，把内存数据序列化(翻译)并保存到硬盘
func saveTasks() {
	// MarshalIndent 把结构体翻译成带缩进、好看的 JSON 格式
	data, err := json.MarshalIndent(memoryTasks, "", "  ")
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	// 0644 是 Linux/Unix 的文件权限设置，代表普通的可读写文件
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
	}
}

// loadTasks 是“反向翻译官”，程序启动时从硬盘读取数据，复原到内存
func loadTasks() {
	// 尝试去读保险柜里的文件
	data, err := os.ReadFile(fileName)
	if err != nil {
		return // 如果文件不存在(比如第一次运行)，啥也不干，直接用空切片跑
	}
	// 【核心知识点：指针】这里必须传 &memoryTasks
	// 把内存真实地址告诉 Unmarshal，它才能把解析好的数据塞进切片里
	json.Unmarshal(data, &memoryTasks)
}

// handleList 负责把内存里的数据漂亮地打印在屏幕上
func handleList() {
	fmt.Println("--- 当前任务列表 ---")
	// 第一个返回值是索引(这里不需要，所以用下划线 _ 忽略)，第二个是具体元素
	for _, task := range memoryTasks {
		// 打印主任务
		fmt.Printf("[%d] %s - 状态: %s\n", task.ID, task.Name, task.Status)

		// 检查这个主任务肚子里有没有子任务
		if len(task.SubTasks) > 0 {
			for i, sub := range task.SubTasks {
				// 打印子任务，前面加点空格制造层级感
				fmt.Printf("  └─ 子任务[%d-%d] %s - 状态: %s\n", task.ID, i+1, sub.Name, sub.Status)
			}
		}
	}
}

// handleAdd 演示了基础的切片追加 (Append)
func handleAdd(name string) {
	// 简单生成一个递增的 ID
	newID := len(memoryTasks) + 1

	// 组装一个纯正的 Go 结构体数据
	newTask := Task{
		ID:     newID,
		Name:   name,
		Status: "待办",
	}
	// 追加到白板上
	memoryTasks = append(memoryTasks, newTask)

	fmt.Printf("成功添加任务: %s (ID: %d)\n", name, newID)

	saveTasks()
	handleList()
}

// handleAddSub 演示了多层级结构体的深度操作
func handleAddSub(parentID int, subName string) {
	found := false

	for i, task := range memoryTasks {
		// 先找到那个要做父亲的主任务
		if task.ID == parentID {
			newSub := Task{
				ID:     len(task.SubTasks) + 1,
				Name:   subName,
				Status: "待办",
			}

			// 精确制导：把子任务追加到 memoryTasks[i] 这个老爹的 SubTasks 肚子里
			memoryTasks[i].SubTasks = append(memoryTasks[i].SubTasks, newSub)
			found = true
			fmt.Printf("成功向任务 [%d] 添加子任务: %s\n", parentID, subName)
			break
		}
	}

	if !found {
		fmt.Printf("错误: 找不到 ID 为 %d 的主任务。\n", parentID)
		return
	}

	saveTasks()
	handleList()
}

// handleExport 是后端/DevOps高频需求：导出为 CSV 表格
func handleExport() {
	file, err := os.Create("tasks_export.csv")
	if err != nil {
		fmt.Println("创建 CSV 文件失败:", err)
		return
	}
	// 【核心知识点：延迟执行】defer 保证在函数结束退出的那一刻，一定会执行关闭文件的操作，防止内存泄漏
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // 确保内存缓冲区的数据全部被“冲刷”进硬盘

	// 写入表头，要求传入一个字符串切片
	writer.Write([]string{"层级", "任务ID", "任务名称", "状态"})

	// 把树状嵌套的数据，拍平变成一行行的二维表格记录
	for _, task := range memoryTasks {
		// 【避坑点：类型转换】CSV 写入器只认字符串 (string)
		// 所以必须用 strconv.Itoa 把数字格式的 ID 强转回字符串格式
		writer.Write([]string{"主任务", strconv.Itoa(task.ID), task.Name, task.Status})

		for i, sub := range task.SubTasks {
			// Sprintf 不会直接打印在屏幕上，而是把排版好的结果当成一个字符串返回给你
			subID := fmt.Sprintf("%d-%d", task.ID, i+1)
			writer.Write([]string{"  └ 子任务", subID, sub.Name, sub.Status})
		}
	}

	fmt.Println("🎉 成功将任务导出到 tasks_export.csv！")
}

// 程序总入口
func main() {
	// 1. 上班第一件事：读盘恢复记忆！
	loadTasks()

	// 2. 长度防御：如果老板说话大喘气（只输入了 go run main.go）
	if len(os.Args) < 2 {
		fmt.Println("用法: task-cli [command]")
		fmt.Println("可用命令: add, list, done, delete, add-sub, export")
		os.Exit(1)
	}

	// 3. 提取老板下达的主命令（比如 add, list）
	command := os.Args[1]

	// 4. 根据命令分配任务（路由分发）
	switch command {
	case "list":
		handleList()
	case "add":
		// 切片防越界保护
		if len(os.Args) < 3 {
			fmt.Println("错误: add 命令需要提供任务名称。")
			os.Exit(1)
		}
		taskName := os.Args[2]
		handleAdd(taskName)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("错误: done 命令需要提供任务 ID。")
			os.Exit(1)
		}
		idStr := os.Args[2]            // 拿到的是字符串
		id, err := strconv.Atoi(idStr) // 转成整数
		if err != nil {
			fmt.Println("错误: 任务 ID 必须是数字。")
			os.Exit(1)
		}
		handleDone(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("错误: delete 命令需要提供任务 ID。")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("错误: 任务 ID 必须是数字。")
			os.Exit(1)
		}
		handleDelete(id)

	case "add-sub":
		if len(os.Args) < 4 {
			fmt.Println("错误: add-sub 命令需要提供主任务 ID 和 子任务名称。")
			fmt.Println("示例: task-cli add-sub 1 \"子任务名\"")
			os.Exit(1)
		}
		parentID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("错误: 主任务 ID 必须是数字。")
			os.Exit(1)
		}
		subName := os.Args[3]
		handleAddSub(parentID, subName)

	case "export":
		handleExport()

	default:
		// 如果用户乱输命令，给予提示
		fmt.Printf("未知命令: %s\n", command)
	}
}
