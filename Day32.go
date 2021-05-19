package main

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"time"
)

// TODO: ---- 32、panic 和 recover See: https://studygolang.com/articles/12785

/*
 当程序发生异常时，无法继续运行。在这种情况下，我们会使用 panic 来终止程序
 当程序发生 panic 时，使用 recover 可以重新获得对该程序的控制
*/

/*
 panic 有两个合理的用例。
 发生了一个不能恢复的错误，此时程序不能继续运行。 一个例子就是 web 服务器无法绑定所要求的端口。在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。
 发生了一个编程上的错误。 假如我们有一个接收指针参数的方法，而其他人使用 nil 作为参数调用了它。在这种情况下，我们可以使用 panic，因为这是一个编程错误：用 nil 参数调用了一个只能接收合法指针的方法。
*/

/// panic 示例
func fullName(firstName *string, lastName *string) {
	//defer fmt.Println("\n deferred call in fullName") /// 发生 panic 时的 defer
	defer recoverName32() /// recover
	if firstName == nil {
		panic(" runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic(" runtime error: last name cannot be nil")
	}
	fmt.Printf("\n %s %s\n", *firstName, *lastName)
	fmt.Println(" returned normally from fullName")
}

/// recover
func recoverName32()  {
	if r := recover(); r != nil {
		fmt.Println(" Recover from ", r)
	}
}

/// panic，recover 和 Go 协程
/*
  只有在相同的 Go 协程中调用 recover 才管用。
  recover 不能恢复一个不同协程的 panic。
*/
func recovery32() {
	if r := recover(); r != nil {
		fmt.Println(" recovered:", r)
	}
}

func a() {
	defer recovery32()
	fmt.Println(" Inside A")
	b() // go b() (修改为 b()，就可以恢复 panic 了，因为 panic 发生在与 recover 相同的协程里)
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println(" Inside B")
	panic(" oh! B panicked")
}

/// 运行时 panic
func r32() {
	if r := recover(); r != nil {
		fmt.Println(" Recovered", r)
		/// 恢复后获得堆栈跟踪
		debug.PrintStack()
	}
}
func a32() {
	defer r32() // 修改恢复运行时 panic
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println(" normally returned from a")
}

func day32()  {

	/// panic 示例\  recover
	defer fmt.Println(" deferred call in main") /// 发生 panic 时的 defer
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println(" returned normally from main")

	/// panic，recover 和 Go 协程
	a()
	fmt.Println(" normally returned from main")

	/// 运行时 panic
	a32()
	fmt.Println(" normally returned from main")

}

// TODO: ---- 33、函数是一等公民（头等函数） See: https://studygolang.com/articles/12789

/// 什么是头等（第一类）函数？
/*
 支持头等函数（First Class Function）的编程语言，可以把函数赋值给变量，
 也可以把函数作为其它函数的参数或者返回值
*/


/// 用户自定义的函数类型
type add33 func(a int, b int) int

/// 高阶函数 把函数作为参数，传递给其它函数
func simple(a func(a, b int) int) {
	fmt.Println( " ",a(60, 7))
}

/// 在其它函数中返回函数
func simple33() func(a, b int) int  {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

/// 闭包
func appendStr() func(string) string {
	t := " Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

/// 头等函数的实际用途
type student33 struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter33(s []student33, f func(student33) bool) []student33 {
	var r []student33
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

/// 我们把这种对集合中的每个元素进行操作的函数称为 map 函数
func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func day33()  {

	/// 函数赋值给一个变量
	a := func() {
		fmt.Println("\n hello world first class function")
	}
	a()
	fmt.Printf(" %T", a) // func()

	/// 不用赋值给变量
	func() {
		fmt.Println("\n new hello world first class function")
	}()

	/// 向匿名函数传递参数
	func(n string){
		fmt.Println(" Welcome", n)
	}("Gophers") // Welcome Gophers

	/// 用户自定义的函数类型
	var a33 add33 = func(a int, b int) int {
		return a + b
	}
    s := a33(5, 6)
    fmt.Println(" Sum", s) // 11

    /// 把函数作为参数，传递给其它函数
    f := func(a, b int) int {
    	return a + b
	}
	simple(f) // 67

	/// 在其它函数中返回函数
	sss := simple33()
	fmt.Println(" ",sss(60, 7)) // 67

	/// 闭包
    new_a := 5
    func() {
    	fmt.Println(" new_a value is ", new_a)
	}()
	a11 := appendStr()
	b11 := appendStr()
	fmt.Println(a11(" World"))
	fmt.Println(b11(" Everyone"))

	fmt.Println(a11(" Gopher"))
	fmt.Println(b11(" !"))

    /// 头等函数的实际用途
    s1 := student33{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student33{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
    new_s := []student33{s1, s2}
    new_f := filter33(new_s, func(s student33) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
    fmt.Println(" ", new_f)

    /// map 函数
	a55 := []int{5, 6, 7, 8, 9}
	r55 := iMap(a55, func(n int) int {
		return n * 5
	})
	fmt.Println(" ", r55)

}

// TODO: ---- 34、反射 See: https://studygolang.com/articles/13178

/// 什么是反射

/*
  反射就是程序能够在运行时检查变量和值，求出它们的类型。
*/
type order struct {
	ordId      int
	customerId int
}

func createQuery(o order) string {
	i := fmt.Sprintf("\n insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

/// 优化通用  [reflect.Kind Kind 和 Type 的类型可能看起来很相似 (Kind 表示该类型的特定类别)]
func createQueryUniversal(q interface{}) {
	t := reflect.TypeOf(q)
	//v := reflect.ValueOf(q)
	v := t.Kind()
	fmt.Println(" Type ", t)
	fmt.Println(" Value ", v)
}

/// NumField() 和 Field() 方法
func createQuery34(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println(" Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf(" Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}

/// 完整的程序
type employeeAll struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQueryAll(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf(" insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf(" %s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf(" %s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf(" %s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf(" %s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println(" Unsupported type")
				return
			}
		}
		query = fmt.Sprintf(" %s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println(" unsupported type")
}

func day34()  {

	/// SQL 插入查询。
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(" ",createQuery(o))

	/// reflect.Kind
	createQueryUniversal(o)

	/// NumField() 和 Field() 方法
	createQuery34(o)

	/// Int() 和 String() 方法 (可以帮助我们分别取出 reflect.Value 作为 int64 和 string。)
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

    /// 完整程序
	e := employeeAll{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQueryAll(e)
	i := 90
	createQueryAll(i)

}