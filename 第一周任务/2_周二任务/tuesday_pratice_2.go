package main

import "fmt"

func main() {
	// 1. 正确地制造一个 Map
	// 需求：用 make 函数造一个字典，Key 是设备名字(string)，Value 是电量(int)
	// powerMap := make(???)
	powerMap := make(map[string]int)
	// 2. 手动往里面存数据
	powerMap["Sensor_A"] = 95
	powerMap["Sensor_B"] = 20

	fmt.Println("传感器 A 的电量是:", powerMap["Sensor_A"])

	// 3. 见证 Map 的特性：读取一个根本不存在的设备
	// 需求：打印出一个根本没存进去的 "Sensor_C" 的电量，看看程序会不会崩溃？
	// fmt.Println("未知的传感器 C 电量是:", ???)
	// 在 Go 中，访问不存在的 Key 会返回该类型的零值（int 的零值是 0），而不会像 Python 那样崩溃
	fmt.Println("未知的传感器 C 电量是:", powerMap["Sensor_C"])

	// 4. 修改数据
	// 需求：把 Sensor_B 的电量修改为 15，并打印出来看看是否生效
	// powerMap["Sensor_B"] = ???
	// fmt.Println("传感器 B 更新后的电量:", powerMap["Sensor_B"])
	// 直接重新赋值即可覆盖旧值
	powerMap["Sensor_B"] = 15
	fmt.Println("传感器 B 更新后的电量:", powerMap["Sensor_B"])

}
