package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 这里的步骤是创建了一个DateExport的接口，其实我不是很理解，这里的Export是自定义的吗，
// 虽然说只要拥有一共Export 能够接受any 类型参数并返回 error的方法就是鸭子类型
// 我不是很理解鸭子类型是上面意思
// any是空接口是上面意思呢
type DateExport interface {
	Export(data any) error
}

// 这里就是定义结构体，能够理解，文件名存储的形式是字符串
type JSONExporter struct {
	FileName string
}

// 这个也是一样
type CSVExporter struct {
	FileName string
}

// 这里就是定义一共函数自定义变量是j，使用的是JSONExport结构体的变量，定义的方法名叫Export
// 这个Export是我这里定义了 定义接口的DataExport才可以使用是吗
// 然后这里的返回的数据类型是error类型的数据就和float64一样是吗
func (j JSONExporter) Export(data any) error {

	fmt.Printf("准备将数据导出到 JSON 文件: %s\n", j.FileName)

	// 这一段的意思是定义了两个变量用来存储数据转成json格式的字节流是吗
	bytes, err := json.Marshal(data)

	// nil是什么呢好像没有定义这个变量,这个是关键字吗
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %v", err)
	}

	fmt.Printf("成功生成 JSON 内容: %s\n\n", string(bytes))
	return nil

}

func (c CSVExporter) Export(data any) error {

	fmt.Printf("准备将数据导出到 CSV 文件: %s\n", c.FileName)

	// 其实我不是很理解为什么是这个写法
	// switch V := data.(type)
	// 这个意思是声明一个switch 类型的变量V 获取data.(type)中的数据,用v的数据来进行case判断吗
	// 还有一个疑惑是data.(type)是什么意思呢不理解
	switch v := data.(type) {
	// 这里的意思是 v 只要读取的数据和 map的一样就执行吗
	case map[string]string:
		fmt.Printf("成功处理 CSV 数据，包含 %d 条记录\n\n", len(v))
		return nil
	default:
		return fmt.Errorf("CSV 导出失败: 不支持的数据类型 %T", v)
	}

}

func main() {
	sampleData := map[string]string{
		"User1": "Alice",
		"User2": "Bob",
	}

	jsonExp := JSONExporter{FileName: "users.json"}
	csvExp := CSVExporter{FileName: "users.csv"}

	// 主程序中这里的意思是用这个变量通过接口切片获取实例化之后的数据吗.
	// jsonExp , csvExp 通过接口和函数的判断返回的值传入这个切片中然后复制给exporters是吗
	exporters := []DateExport{jsonExp, csvExp}

	// 这里for填写的是_,是为了多创建出来的变量没有使用以防出现的问题吗
	for _, exp := range exporters {

		// 这一步不是很理解
		// 变量exp.Export(sampleData)获取的是什么,脉络我有点模糊.
		err := exp.Export(sampleData)

		if err != nil {
			fmt.Println("❌ 发生错误:", err)
		} else {
			fmt.Println("✅ 导出成功！")
		}

		fmt.Println(strings.Repeat("-", 30))
	}
}
