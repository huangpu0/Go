package main

import "fmt"

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

