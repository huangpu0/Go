package main // 声明只有可执行的包名为 main
import (
	"fmt"  // 引用的fmt包 mian 函数里面打印文本到标准输出
	"math" // 运算等？
)

/// ----------------------------------GO 强类型语言------------------------------

/// func main() - main 是一个特殊的函数。整个程序就是从 main 函数开始运行的。main 函数必须放置在 main 包中。{ 和 } 分别表示 main 函数的开始和结束部分。
func main() {

	// day1 hello world
	fmt.Println("Hello World")
	fmt.Println("Hello World")

	// day2 变量的声明使用
	day2()

	// day3 类型
	day3()

}

//mark: ---- 变量的声明使用
func day2()  {

	var age int = 1000
	var new_age = 100 // 所命名的变量必须使用 否则报错, 系统会自动推导数据类型
	fmt.Println("my name is",age,"new age is",new_age)

	var  width, height int = 100, 50 // 声明多个变量 不赋值的话、系统默认为'0'
	print("width----------",width,"\nheight----------",height)
	width  = 2000
	height = 500
	print("\nnew width----------",width,"\nnew height----------",height)

	// 一个语句中声明不同类型的变量
	var (
		name1 string = "字符串"
		name2 int    = 1000
		name3 float32
	)
	fmt.Println("\nname1----",name1,"\nname2-----",name2,"\nname3------",name3)

	// 简短声明
	old_name, old_age := "limei", 29
	fmt.Println("\n简短声明 名字---",old_name," 年龄---",old_age)
    old_name, old_age1 := "new_limei", 100
	fmt.Println("\n一变一不变-简短声明 old-new-名字---",old_name," new-年龄---",old_age1)
	old_name, old_age = "on--new_limei", 300
	fmt.Println("\nchange--简短声明 old-new-名字---",old_name," new-年龄---",old_age1)

    // 运算
    a, b := 145.8, 543.2
    c := math.Min(a,b)
    fmt.Println("\na b 之间 min value is ", c)
    c  = math.Max(a,b)
	fmt.Println("\na b 之间 max value is ", c)

}

//mark: ---- 类型
func day3()  {



}