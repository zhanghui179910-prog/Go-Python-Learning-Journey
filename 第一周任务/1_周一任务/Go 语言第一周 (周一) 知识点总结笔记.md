📝 Go 语言第一周 (周一) 知识点总结笔记：

1. Go 程序的基础骨架
入口限制：所有的可执行 Go 程序，都必须包含 package main，并且逻辑的起点必须是 func main()。

包管理：使用 import 导入标准库（如 "fmt" 用于格式化输入输出）。Go 对此非常严格，导入了但没使用的包会直接导致编译报错。

2. 变量声明的“两副面孔”
简短声明 (:=)：如 rssiValues := []int{...}。

特点：自动推断类型，代码简洁。

限制：只能用在函数内部，不能用于定义全局变量。

标准声明 (var)：如 var sum int = 0。

特点：显式指定类型 (int)，常用于需要预先分配变量但稍后赋值的场景，或者定义包级别的全局变量。

3. 核心数据结构初探：切片 (Slice)
写法：[]int{-45, -68}。注意方括号内是空的，这在 Go 里叫切片 (Slice)，它是动态伸缩的。如果写成 [10]int{...} 有固定长度，那叫数组 (Array)。Go 语言中绝大多数场景使用的是切片。

4. 遍历神器：for...range
语法：for index, value := range collection { ... }

对应 Python：非常像 Python 里的 for index, value in enumerate(list):。

作用：同时获取当前元素的索引 (i) 和具体的值 (rssi)。如果不需要索引，可以使用空白标识符 _ 忽略它（例如 for _, rssi := range ...），否则声明了不使用也会编译报错。

5. 更优雅的条件分支：switch
无表达式 Switch：Go 的 switch 后面可以不跟具体变量，直接在 case 后面写条件判断（如 case rssi >= -50:）。这完美替代了多重 if-else if 结构。

自带 Break：与 C/C++ 不同，Go 的 case 执行完毕后会自动跳出，不需要手动写 break，避免了逻辑漏洞。

6. 极其严格的“类型转换”
强类型铁律：Go 不会做任何隐式的类型转换。两个 int 类型相除，结果永远是截断后的 int。

显式转换：为了计算精确的平均值，必须用 float64(变量) 强制把整数转换为 64位浮点数，如 float64(sum) / float64(len(rssiValues))。

7. 格式化输出 (fmt.Printf)
%d：占位输出整数。

%4d：输出宽度为 4 的整数（方便排版对齐）。

%s：占位输出字符串。

%.2f：输出保留两位小数的浮点数。

\n：手动换行（Println 自带换行，但 Printf 需要手动加）。

Println 自动换行
Printf 没有

:= 
= 
var int xx = 
xxx := []int {}
for x, x := xxx{}
switch{}