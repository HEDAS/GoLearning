package basic

import (
	"fmt"
)

var pi float32 = 3.14159

func Basic() {
	simple()
	variable()
}

func simple() {
	fmt.Println("==================== simple ====================")
	/*
		声明：
		Go语言和许多编程语言不同，它在声明变量时将变量的类型放在变量的名称之后。
		这样做的好处就是可以避免像C语言中那样含糊不清的声明形式，例如：int* a1, b1; 。其中只有 a1 是指针而 b1 不是。
	*/
	//var a, b *int // 将它们都声明为指针类型

	/*
		Go语言的基本类型有：
		bool
		string
		int、int8、int16、int32、int64
		uint、uint8、uint16、uint32、uint64、uintptr
		byte // uint8 的别名
		rune // int32 的别名 代表一个 Unicode 码
		float32、float64
		complex64、complex128
	*/

	/*
		当一个变量被声明之后，系统自动赋予它该类型的零值：
		int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。
		所有的内存在 Go 中都是经过初始化的。
	*/

	/*
		变量的命名规则遵循骆驼命名法

		声明：
		1.
		var 变量名 变量类型
		2.
		var (
		    a int
		    b string
		    c []float32
		    d func() bool
		    e struct {
		        x int
		    }
		)
		3.
		名字 := 表达式
		简短模式（short variable declaration）有以下限制：
		定义变量，同时显式初始化。
		不能提供数据类型。
		只能用在函数内部。

		简短变量声明被广泛用于大部分的局部变量的声明和初始化。
		var 形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。
	*/

	/*
		微软的 VC 编译器会将未初始化的栈空间以 16 进制的 0xCC 填充，而未初始化的堆空间使用 0xCD 填充，
		而 0xCCCC 和 0xCDCD 在中文的 GB2312 编码中刚好对应“烫”和“屯”字。
		因此，如果一个字符串没有结束符\0，直接输出的内存数据转换为字符串就刚好对应“烫烫烫”和“屯屯屯”。

		初始化
		var 变量名 类型 = 表达式
		var hp int = 100

		将 int 省略后，编译器会尝试根据等号右边的表达式推导 hp 变量的类型。
		var hp = 100
		等号右边的部分在编译原理里被称做右值（rvalue）。

		由于Go语言和C语言一样，编译器会尽量提高精确度，以避免计算中的精度损失。
		Go语言编译器会将小数推导为 float64
	*/
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17                     // 如果不指定类型，Go语言编译器会将 damageRate 类型推导为 float64
	var damage = float32(attack-defence) * damageRate // 使用 float32() 将结果转换为 float32 类型
	fmt.Println(damage)

	/*
		在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
	*/
	//conn, err:=net.Dial("tcp", "127.0.0.1:8080")
	//conn1, err:=net.Dial("tcp", "127.0.0.1:8080")

	/*
		变量交换
		在计算机刚发明时，内存非常“精贵”。计算机“大牛”发明了一些算法来避免使用中间变量
		var a int = 100
		var b int = 200
		a = a ^ b
		b = b ^ a
		a = a ^ b
		fmt.Println(a, b)
		这样的算法很多，但是都有一定的数值范围和类型要求。

		多重赋值在Go语言的错误处理和函数返回值中会大量地使用。
	*/
	var testIntSlice IntSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(testIntSlice.Less(3, 1))
	fmt.Println(testIntSlice.Len())
	testIntSlice.Swap(3, 1)
	fmt.Println(testIntSlice)

	/*
		匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。
		任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用

		匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
		在 Lua 等编程语言里，匿名变量也被叫做哑元变量。
	*/
	a, _ := getData()
	fmt.Println(a)
}

func variable() {
	fmt.Println("==================== variable ====================")
	/*
		函数内定义的变量称为局部变量
		函数外定义的变量称为全局变量
		函数定义中的变量称为形式参数
	*/

	/*
		全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，
		当然，不包含这个全局变量的源文件需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。

		全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。

		Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
	*/
	fmt.Printf("pi = %f\n", pi) // 全局的
	var pi float32 = 3.0
	fmt.Printf("pi = %f\n", pi) // 局部的

	/*
		在定义函数时函数名后面括号中的变量叫做形式参数（简称形参）。
		形式参数只在函数调用时才会生效，函数调用结束后就会被销毁，
		在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值。
	*/
}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getData() (int, int) {
	return 100, 200
}
