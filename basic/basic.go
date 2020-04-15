package basic

import "fmt"

func Basic() {
	declare()
}

func declare() {
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
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)
}
