package main

import (
	"fmt"
	"sort"    //  必须引入，为了给最后乱序的切片排队
	"strings" // 为了处理字符串切割
)

// 【为什么要写这个？】
// 因为 Map 是“无序散沙”，无法排序。
// 我们需要把 Map 里的“名字”和“次数”像胶水一样粘在一起，变成一个“包裹”。
// 这个 ErrorFreq 就是我们定义的“包裹规格”。
type ErrorFreq struct {
	Name  string
	Count int
}

func main() {
	// 【数据源】原始的日志大长串
	logData := "Timeout, Disconnect, Noise, Timeout, CRC_Error, Noise, Timeout, Disconnect, Voltage_Drop, Noise, Timeout"

	// 【为什么要 Split？】
	// 将字符串切成一片一片的单词，存入切片 (Slice) 中。
	// 此时数据形态：["Timeout", "Disconnect", ...]
	errorList := strings.Split(logData, ", ")

	// ================= 1. 统计频率 (Map 阶段) =================

	// 【为什么要用 make？】
	// Go 的 Map 必须先“向内存申请地盘”才能存东西。
	freqMap := make(map[string]int)

	// 【为什么用 _？】
	// range 返回两个值：索引(0,1,2...)和单词。我们不需要数字索引，
	// 但 Go 规定定义的变量必须用，否则报错。下划线 _ 就是“我不想要，扔掉”。
	for _, errName := range errorList {
		// 【为什么这么写就能计数？】
		// freqMap[errName] 如果找不到，会默认返回 0。
		// 找到或返回 0 后再 ++，就完成了“计数并存回去”的动作。
		freqMap[errName]++
	}

	// ================= 2. 打包准备 (Slice 阶段) =================

	// 【为什么要声明这个？】
	// 准备一个空杯子，专门装我们刚才定义的 ErrorFreq “包裹”。
	var freqSlice []ErrorFreq

	// 【为什么要转移数据？】
	// Map 是乱序的，不能排序。所以我们必须把 Map 里的东西一个个掏出来，
	// 打包进 ErrorFreq 结构体，再排进 freqSlice 切片这个“队列”里。
	for k, v := range freqMap {
		// 这里的 append 会触发你观察到的 cap (容量) 变化。
		// 底层数组满了，Go 就偷偷换个大杯子。
		freqSlice = append(freqSlice, ErrorFreq{Name: k, Count: v})

		// 打印监控：亲眼看看“换杯子”的瞬间
		fmt.Printf("存入 %s, len: %d, cap: %d\n", k, len(freqSlice), cap(freqSlice))
	}

	// ================= 3. 排序 (Sort 阶段) =================

	// 【为什么怎么实现排序？】
	// sort.Slice 会在这个队列里两两对比。
	// i 和 j 是两个格子的编号。如果 i 的次数大于 j 的次数，就让 i 站前面。
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Count > freqSlice[j].Count // 降序排列
	})

	// ================= 4. 输出 (Final 阶段) =================

	fmt.Println("\n--- 周二打卡：故障高频词排行榜 ---")
	// 把排好队的切片一个个印出来
	for i, item := range freqSlice {
		fmt.Printf("排名 %d | 故障类型: %-15s | 出现次数: %d\n", i+1, item.Name, item.Count)
	}
}
