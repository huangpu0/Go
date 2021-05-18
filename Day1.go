package main // 声明只有可执行的包名为 main
import (
	"fmt"  // 引用的fmt包 mian 函数里面打印文本到标准输出
	"math" // 运算等？
	"unsafe"
)

/// ----------------------------------GO 强类型语言 SDK 1.16.3 ------------------------------

/// func main() - main 是一个特殊的函数。整个程序就是从 main 函数开始运行的。main 函数必须放置在 main 包中。{ 和 } 分别表示 main 函数的开始和结束部分。
func main() {

	// day2 hello world
	fmt.Println(" Hello World")
	fmt.Println(" Hello World")

	// day3 变量的声明使用
	//day3()

	// day4 类型
	//day4()

	// day5 常量
	//day5()

	// day6 函数
	//day6()

	// day7 包
	//day7()

	// day8 if-else语句
	//day8()

	// day9 循环
	//day9()

	// day10 switch语句
	//day10()

	// day11 数组与切片
	//day11()

	// day12 可变参数函数
	//day12()

	// day13 Maps
	//day13()

    // day14 字符串
    //day14()

	// day15 指针
	//day15()

	// day16 结构体
	//day16()

	// day17 方法
	//day17()

	// day18_19 接口（1、2）
	//day18_19()

	// day21 Go协程
	//day21()

	// day22 信道
	//day22()

	// day23 缓存信道和工作池
	//day23()

	// day24 Select
	//day24()

	// day25 Mutex
	//day25()

    // day26 结构体取代类
    //day26()

	// day27 组合取代继承
	//day27()

	// day28 多态
	//day28()

	// day29 Defer
	//day29()

	// day30 错误处理
	//day30()

	// day31 自定义错误
    day31()

}

// TODO: ---- 3、变量的声明使用 See: https://studygolang.com/articles/11756
func day3() {

	var age int = 1000
	var new_age = 100 // 所命名的变量必须使用 否则报错, 系统会自动推导数据类型
	fmt.Println("\n my name is", age, "new age is", new_age)

	var width, height int = 100, 50 // 声明多个变量 不赋值的话、系统默认为'0'
	print("\n width----------", width, "\nheight----------", height)
	width = 2000
	height = 500
	print("\n new width----------", width, "\nnew height----------", height)

	// 一个语句中声明不同类型的变量
	var (
		name1 string = "字符串"
		name2 int    = 1000
		name3 float32
	)
	fmt.Println("\n name1----", name1, "\nname2-----", name2, "\nname3------", name3)

	// 简短声明
	old_name, old_age := "limei", 29
	fmt.Println("\n 简短声明 名字---", old_name, " 年龄---", old_age)
	old_name, old_age1 := "new_limei", 100
	fmt.Println("\n 一变一不变-简短声明 old-new-名字---", old_name, " new-年龄---", old_age1)
	old_name, old_age = "on--new_limei", 300
	fmt.Println("\n change--简短声明 old-new-名字---", old_name, " new-年龄---", old_age1)

	// 运算
	a, b := 145.8, 543.2
	c := math.Min(a, b)
	fmt.Println("\n a b 之间 min value is ", c)
	c = math.Max(a, b)
	fmt.Println("\n a b 之间 max value is ", c)

}

// TODO: ---- 4、类型 See: https://studygolang.com/articles/11869
func day4() {

	// 支持的基本类型 bool, string
	a := true
	b := false
	fmt.Println("\n a: ", a, " b: ", b)
	c := a && b
	fmt.Println("\n c: ", c)
	d := a || b
	fmt.Println("\n d:", d)

	// string
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("\n My name is", name)

	/* 数字类型 */

	// 有符号整型
	// int8 (-128 ~~ 127), int16 (-32768 ~~ 32767), int32 , int64, int
	var int_a int = 89
	int_b := 95
	fmt.Println("\n value of int_a is", int_a, "and int_b is ", int_b)
	// int_a 的类型及大小 %T 用于打印类型，而 %d 用于打印字节大小
	fmt.Println("\n type of int_a is %T, size of int_a is %d", int_a, unsafe.Sizeof(int_a))
	// int_b 的类型和大小
	fmt.Println("\n type of int_b is %T, size of int_b is %d", int_b, unsafe.Sizeof(int_b))

	// 无符号整型
	// unit8（0 ～ 255）, unit16 0 ～ 65535）, unit32, unit64, unit

	// float32 （32 位浮点数）, float64 （64 位浮点数）
	float_a, float_b := 5.67, 8.97
	fmt.Println("\n type of float_a %T b %T\n", float_a, float_b)
	sum := float_a + float_b
	diff := float_a - float_b
	fmt.Println("\n float sum: ", sum, " diff: ", diff)

	no1, no2 := 56, 89
	fmt.Println("\n int sum", no1+no2, " diff: ", no1-no2)

	// complex64 （实部和虚部都是 'float32' 类型的复数）, complex128 （实部和虚部都是 'float64' 类型的复数）
	c1 := complex(5, 7)
	fmt.Println("\n c1: ", c1)
	c2 := 8 + 27i
	fmt.Println("\n c2: ", c2)
	cadd := c1 + c2
	fmt.Println("\n complex sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("\n product: ", cmul)

	// byte （uint8的别名）
	// rune （int32的别名）

	// TODO: ---------------------------- 类型转换 ----------------------------
	i := 55   // int
	j := 67.8 // float64
	fmt.Println("\n i: ", i, "j: ", j)
	/*
	   new_sum := i + j // 不允许 int + float64 (mismatched types int and float64)
	   fmt.Println("\n new_sum: ",new_sum)
	*/

	new_sum := i + int(j) //j is converted to int
	fmt.Println("\n new_sum: ", new_sum)

	// i
	new_i := 10
	var new_j float64 = float64(new_i)
	fmt.Println("\n new_j: ", new_j)

}

// TODO: ---- 5、常量 See: https://studygolang.com/articles/11872
func day5() {

	const a int = 50 // 允许
	// a = 89 // 不允许重新赋值 （关键字 const）
	fmt.Println("\n a: ", a)
	var aaa = math.Sqrt(4) // 允许
	fmt.Println("\n aaa value is ", aaa)
	// const  bbb = math.Sqrt(4) // 不允许 (error: const initializer math.Sqrt(4) is not a constant)

	// TODO: -------------------- 字符串常量 --------------------
	const hello = "Hello World"
	fmt.Println("\n hello is %T value is ", hello, hello)
	var name = "Sam"
	fmt.Println("\n type %T value %v", name, name)

	// 带有类型的常量
	const typehell0 string = "Hello World"
	fmt.Println("\n type is value is ", typehell0)

	// TODO: -------------------- GO 强类型语言 分配过程中混合类型是不被允许的 --------------------
	var defaultName = "Sam" // 允许
	fmt.Println("\n defalutName is ", defaultName)
	type myString string
	var customName myString = "Sam" // 允许
	fmt.Println("\n customeName is ", customName)
	// customName = defaultName // 不允许 (error cannot use defaultName (type string) as type myString in assignment)

	// TODO: -------------------- bool类型 --------------------
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst // 允许
	fmt.Println("\n defaultBool is value ", defaultBool)
	var customeBool myBool = trueConst // 允许
	fmt.Println("\n customeBool is value ", customeBool)
	// defaultBool = customeBool // 不允许 （error cannot use customeBool (type myBool) as type bool in assignment ）

	// TODO: -------------------- 数字常量 --------------------
	const aa = 5
	var intVar int = aa
	var int32Var int32 = aa
	var float64Var float64 = aa
	var complex64Var complex64 = aa
	fmt.Println("\n intVar: ", intVar, " int32Var: ", int32Var, " float64Var: ", float64Var, " complex64Var: ", complex64Var)

	// TODO: -------------------- 数字表达式 --------------------
	var new_a = 5.9 / 8
	fmt.Println("\n new_a type %T value %v ", new_a, new_a)

}
