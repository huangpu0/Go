package main

import "fmt"

// TODO: ---- 18、接口（一） See: https://studygolang.com/articles/12266

// --------------------------------- 接口的声明与实现 ---------------------------------
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString)FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// TODO: --------------------------------- 接口的实际用途 ---------------------------------
/*
 我们编写一个简单程序，根据公司员工的个人薪资，计算公司的总支出。
 为了简单起见，我们假定支出的单位都是美元。
*/
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

// salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
 total expense is calculated by iterating though the SalaryCalculator slice and summing
 the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("\n Total Expense Per Month $%d", expense)
}

// TODO: --------------------------------- 接口内部表示 ---------------------------------
type Test interface {
	Tester()
}
type MyFloat float64

func (m MyFloat) Tester()  {
	fmt.Println("\n m value is ",m)
}

func describe(t Test)  {
	fmt.Printf("\n Interface type %T value is %v\n", t, t)
}

// TODO: --------------------------------- 空接口 ---------------------------------
func new_describe(i interface{})  {
	fmt.Printf("\n Type = %T, value = %v\n", i, i)
}

// TODO: --------------------------------- 类型断言 ---------------------------------
func assert(i interface{})  {
	/*
	 /// int 类型OK 字符串 error
	 error!!!  panic: interface conversion: interface {} is string, not int.。
	 s := i.(int) // get the underlying int value from i
	 fmt.Println("\n s value is", s)
	*/
	v, ok := i.(int)
	fmt.Println("\n value is ", v, ok)

}

// TODO: --------------------------------- 类型选择 ---------------------------------
func findType(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Printf("\n I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("\n I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("\n Unknown type\n")
	}
}

/// 还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
type Describer interface {
	Describe()
}
type NEW_Person struct {
	name string
	age  int
}

func (p NEW_Person) Describe() {
	fmt.Printf("\n %s is %d years old", p.name, p.age) // Naveen R is 25 years old
}

func new_findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("\n unknown type\n")
	}
}


// TODO: --------------------------------- 实现接口：指针接受者与值接受者 ---------------------------------
type New_Address struct {
	state   string
	country string
}
func (a *New_Address) Describe() { // 使用指针接受者实现
	fmt.Printf("\n State %s Country %s", a.state, a.country)
}


// TODO: --------------------------------- 实现多个接口 ---------------------------------
type New_SalaryCalculator interface {
	New_DisplaySalary()
}

type New_LeaveCalculator interface {
	New_CalculateLeavesLeft() int
}

type New_Employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e New_Employee)New_DisplaySalary()  {
	fmt.Printf("\n %s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e New_Employee)New_CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

// TODO: --------------------------------- 接口的嵌套 ---------------------------------
type New_EmployeeOperations interface {
	New_SalaryCalculator
	New_LeaveCalculator
}

// TODO: --------------------------------- 接口的零值 ---------------------------------
type New_Describer interface {
	New_Describer()
}

func day18_19() {

	/// 接口的声明与实现
	name := MyString("Sam Anderson")
	var  v VowelsFinder
	v = name
	fmt.Println("\n Vowels are %c", v.FindVowels()) //  [97 101 111]

	/// 接口实际用途
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

    /// 接口的内部表示
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()

	/// 空接口
	s := "Hello World"
	new_describe(s)
	i := 55
	new_describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	new_describe(strt)

	/// 类型判断
	var  new_s    interface{} = 56
	assert(new_s)
	var  new_str  interface{} = "Steven Paul"
	assert(new_str)

	/// 类型选择
	findType("Naveen")
	findType(77)
	findType(89.98)

	new_findType("Naveen")
	p := NEW_Person{
		name: "Naveen R",
		age: 25,
	}
	new_findType(p)

	/// (接口2) ------- 实现接口：指针接受者与值接受者
	var d1 Describer
	p1 := NEW_Person{"Sam",25}
	d1 = p1
	d1.Describe()
	p2 := NEW_Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := New_Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a // 这是合法的 New_Address 类型的指针实现了 Describer 接口
	d2.Describe()


	/// 实现多个接口
	e := New_Employee{
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 5000,
		pf: 200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var new_s1 New_SalaryCalculator = e
	new_s1.New_DisplaySalary()
	var new_l New_LeaveCalculator = e
	fmt.Printf("\n Leaves left =", new_l.New_CalculateLeavesLeft()) //  Leaves left =%!(EXTRA int=25)

    /// 接口的嵌套
    var empOp New_EmployeeOperations = e
    empOp.New_DisplaySalary()
    fmt.Printf("\n Leaves left =", empOp.New_CalculateLeavesLeft()) // Leaves left =%!(EXTRA int=25)


    /// 接口的零值
	var new_d1 New_Describer
	if new_d1 == nil {
		fmt.Printf("\n new_d1 is nil and has type %T value %v\n", new_d1, new_d1)
	}
	/*
	error!!!
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10a60ee]

	new_d1.New_Describer()
    */


}
