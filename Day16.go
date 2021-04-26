package main

import (
	"fmt"
)

// TODO: ---- 结构体 See: https://studygolang.com/articles/12263

// 声明
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// 匿名结构体
var employee struct {
	firstName, lastName string
	age int
}

// 匿名字段
type Person struct {
	string
	int
}

// 嵌套结构体
type Address struct {
	city, state string
}
type NEWPerson struct {
	name string
	age int
	address Address
}
// 提升字段
type NEWPerson1 struct {
	name string
	age int
	Address
}

// 结构体相等性 (如果结构体包含不可比较的字段，则结构体变量也不可比较。)
type image struct {
	data map[int]int
}

func day16()  {

	// -------------------------------- 创建命名的结构体 --------------------------------
	// creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age: 25,
		salary: 500,
		lastName: "Anderson",
	}

	fmt.Println("\n emp1 value is ", emp1) //  {Sam Anderson 25 500}

	// creating structure using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("\n emp2 value is ", emp2) //  {Thomas Paul 29 800}

	// -------------------------------- 创建匿名的结构体 --------------------------------
    emp3 := struct {
		firstName, lastName string
		age, salary int
	}{
    	firstName: "Andreah",
    	lastName: "Nikola",
    	age: 31,
    	salary: 5000,
	}
	fmt.Println("\n emp3 value is ", emp3) // {Andreah Nikola 31 5000}

	// -------------------------------- 结构体的零值 --------------------------------
	var emp4 Employee // zero valued structure
    fmt.Println("\n emp4 value is ", emp4) // {  0 0}
    // 也可单独为某一字段赋值
    emp5 := Employee{
    	firstName: "John",
    	lastName: "Paul",
	}
    fmt.Println("\n emp5 value is ", emp5) //  {John Paul 0 0}


	// -------------------------------- 访问结构体的字段 --------------------------------
    fmt.Println("\n emp5 firstName value is ", emp5.firstName) // John
    // 单独为某一字段赋值
    emp5.firstName = "new_john"
	fmt.Println("\n emp5 new_firstName value is ", emp5.firstName) // new_john


	// -------------------------------- 结构体的指针 --------------------------------
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("\n emp8 firstName vaule is ", (*emp8).firstName) // Sam
	fmt.Println("\n emp8 age value is ", (*emp8).age) // 55


	// -------------------------------- 匿名字段 --------------------------------
	p := Person{"Naveen", 50}
    fmt.Println("\n p value is ", p) // {Naveen 50}
    // 虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型
    p.string = "new_Naveen"
    fmt.Println("\n p new string value is ", p.string) //  new_Naveen


	// -------------------------------- 嵌套结构体、提升字段 --------------------------------
	var new_p NEWPerson
	new_p.name = "Naveen"
	new_p.age = 50
	new_p.address = Address {
		city: "Chicago",
		state: "Illinois",
	}
    fmt.Println("\n new_p name value is ", new_p.name) //  Naveen
	fmt.Println("\n new_p city value is ", new_p.address.city) // Chicago
	var new_p1 NEWPerson1
	new_p1.name = "Naveen"
	new_p1.age = 50
	new_p1.Address = Address {
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("\n new_p1 name value is ", new_p1.name) // Naveen
	fmt.Println("\n new_p1 city value is ", new_p1.city) // Chicago


	// TODO: -------------------------------- 导出结构体和字段 --------------------------------


	// -------------------------------- 结构体相等性 --------------------------------
    name1 := Address{"Sahnghai","state1"}
	name2 := Address{"Sahnghai","state1"}
	if name1 == name2 {
		fmt.Println("name1 == name2 ")
	}else {
		fmt.Println("name1 != name2 ")
	}
	/* (如果结构体包含不可比较的字段，则结构体变量也不可比较。)
	!!! error: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)。

	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}
    */

}



// TODO: ---- 方法 See: https://studygolang.com/articles/12264
func day17()  {

	emp1 := Employee{
		firstName: "Sam",
		age: 35,
		salary: 500,
		lastName: "Anderson",
	}
	displaySalary(emp1)

	/*
	那么什么时候使用指针接收器，什么时候使用值接收器？
	一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

	指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。
	考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
	在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

	在其他的所有情况，值接收器都可以被使用。
	*/
	emp1.changeName("Michael")
	fmt.Printf("\n Employee firstName after change: %s", emp1.firstName) // Sam (未改变)

	emp1.changeAge(51)
	fmt.Printf("\n Employee age after change: %d", emp1.age) // 51 (改变)


	// ----------------- 在方法中使用值接收器 与 在函数中使用值参数 -----------------
	newValue(emp1)
	new_emp1 := &emp1
    new_emp1.newValue()


	// ----------------- 在非结构体上的方法 -----------------
	num1 := myInt(5)
	num2 := myInt(10)
	sum  := num1.add(num2)
	fmt.Println("\n Sum is", sum)


}

func displaySalary(e Employee)  {
	fmt.Println("\n Salary of %s is %s%d", e.firstName, e.lastName, e.salary)
}

/*
 使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
	e.firstName = newName
	fmt.Println("\n firstName value is ", e.firstName)
}

/*
 使用指针接收器的方法。
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func newValue(e Employee)  {
	fmt.Println("\n newvalue is ", (e.age * e.salary))
}
func (e Employee)newValue()  {
	fmt.Println("\n newvalue is ", (e.age * e.salary))
}

// 在非结构体上的方法
type myInt int

func (a myInt)add(b myInt) myInt {
	return a + b
}
