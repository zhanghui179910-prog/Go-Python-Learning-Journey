package main

import "fmt"

type Resistor struct {
	Name  string
	Value float64
}

func (r Resistor) GetPower(current float64) float64 {
	return current * current * r.Value
}

type SeriesCircuit struct {
	Name      string
	Resistors []Resistor
}

func (c *SeriesCircuit) TotalResistance() float64 {
	total := 0.0
	for _, r := range c.Resistors {
		total += r.Value
	}
	return total
}

func main() {

	r1 := Resistor{Name: "R1(限流电阻)", Value: 100.0}
	r2 := Resistor{Name: "R2(分压电阻)", Value: 220.0}

	fmt.Printf("%s 在2A电流下的功耗: %.2f w\n", r1.Name, r1.GetPower(2.0))

	myCircuit := SeriesCircuit{
		Name:      "测试主路",
		Resistors: []Resistor{r1, r2},
	}

	fmt.Printf("[%s] 的总阻值为: %.2f 欧姆\n", myCircuit.Name, myCircuit.TotalResistance())
}
