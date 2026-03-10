package main

import (
	"fmt"
	"sort"
	"strings"
)

type ErrorFrea struct {
	Name  string
	Count int
}

func main() {

	logData := "Timeout, Disconnect, Noise, Timeout, CRC_Error, Noise, Timeout, Disconnect, Voltage_Drop, Noise, Timeout"

	errorList := strings.Split(logData, ",")

	freMap := make(map[string]int)

	// 循环阅读erroList 中的数据并检查有没有errname无责 errName ++

	for _, errName := range errorList {

		freMap[errName]++

	}

	// 声明定义变量，遍历循环fremap 将 其中的K，v读取出来，并存储到freqslice中。

	var freqSlice []ErrorFrea

	for k, v := range freMap {

		freqSlice = append(freqSlice, ErrorFrea{Name: k, Count: v})

		fmt.Printf("存入 %s, len: %d, cap: %d\n", k, len(freqSlice), cap(freqSlice))

	}

	// 对其进行排序。后面的func不是很懂。

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].Count > freqSlice[j].Count
	})

	fmt.Println("\n--- 周二打卡：故障高频词排行榜 ---")

	for i, item := range freqSlice {
		fmt.Printf("排名 %d | 故障类型: %-15s | 出现次数: %d\n", i+1, item.Name, item.Count)
	}
}
