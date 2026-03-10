package main

import (
	"fmt"
)

func main() {

	rssiValues := []int{-45, -68, -85, -95, -50, -72, -40, -88, -60, -100}

	var sum int = 0

	for i, rssi := range rssiValues {

		sum += rssi

		var status string
		switch {
		case rssi >= -50:
			status = "信号极佳"
		case rssi >= -70:
			status = "信号良好"
		case rssi >= -90:
			status = "信号边缘"
		default:
			status = "存在丢包风险"
		}

		fmt.Printf("终端 %d | 信号强度：%4d dBm | 评估结果：%s\n", i+1, rssi, status)

	}

	average := float64(sum) / float64(len(rssiValues))
	fmt.Println("-------------------------------")
	fmt.Printf("共检测 %d 个终端，平均 RSSI 值为：%.2f dBm\n", len(rssiValues), average)
}
