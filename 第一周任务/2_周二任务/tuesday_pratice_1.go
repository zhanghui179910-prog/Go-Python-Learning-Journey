package main

import "fmt"

func main() {
	// 1. 声明一个纯正的 数组 (Array)
	// 注意看：方括号里有数字 3，代表它的长度死死焊在了 3
	var arr [3]int = [3]int{10, 20, 30}

	// 2. 声明一个纯正的 切片 (Slice)
	// 注意看：方括号里是空的，代表它是弹簧皮筋，随便拉伸
	var sli []int = []int{10, 20, 30}

	fmt.Println("刚开始的数组:", arr)
	fmt.Println("刚开始的切片:", sli)

	// ================= 破坏性试验 =================

	// 试图给切片加一个新数据 (40)
	sli = append(sli, 40)
	fmt.Println("追加后的切片:", sli)

	//

	// 试图给数组加一个新数据 (40)
	// ⚠️ 编译器会在这里立刻把你拦下！
	arr = append(arr, 40)
	fmt.Println("追加后的数组:", arr)
}
