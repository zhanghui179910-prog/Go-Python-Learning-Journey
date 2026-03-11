package main

import "fmt"

// 1. 定义电阻结构体 (抛弃 Class，使用 Struct)
type Resistor struct {
	Name  string
	Value float64 // 阻值 (欧姆)
}

// 2. 为结构体编写方法：计算功耗 P = I^2 * R
// 这里的 (r Resistor) 叫做 "值接收者"，类似于 Python 的 self，但它是值的拷贝
func (r Resistor) GetPower(current float64) float64 {
	return current * current * r.Value
}

// 3. 结构体嵌套：定义一个串联电路，里面包含一个电阻的切片 (Slice)
type SeriesCircuit struct {
	Name      string
	Resistors []Resistor
}

// 4. 计算串联总阻值
// 这里的 (c *SeriesCircuit) 叫做 "指针接收者"。
// 核心法则：如果你需要在方法里修改结构体的数据，或者结构体很大想避免拷贝浪费内存，就用 * 指针。
func (c *SeriesCircuit) TotalResistance() float64 {
	total := 0.0
	for _, r := range c.Resistors {
		total += r.Value // 串联电阻直接相加
	}
	return total
}

func main() {
	// 初始化元器件
	r1 := Resistor{Name: "R1(限流电阻)", Value: 100.0}
	r2 := Resistor{Name: "R2(分压电阻)", Value: 220.0}

	fmt.Printf("%s 在 2A 电流下的功耗: %.2f W\n", r1.Name, r1.GetPower(2.0))

	// 组装电路
	myCircuit := SeriesCircuit{
		Name:      "测试主路",
		Resistors: []Resistor{r1, r2},
	}

	fmt.Printf("[%s] 的总阻值为: %.2f 欧姆\n", myCircuit.Name, myCircuit.TotalResistance())
}
