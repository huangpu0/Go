package main

import (
	fmt "fmt"
)

// TODO: ----  8、if-else 语句 See: https://studygolang.com/articles/11902
func day8() {

	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("\n the number is even")
	} else {
		fmt.Println("\n the number is odd")
	}

	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println("\n new the number is even")
	} else {
		fmt.Println("\n new the number is odd")
	}

	num = 99
	if num <= 50 {
		fmt.Println("\n number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("\n number is between 51 and 100")
	} else { //else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。
		fmt.Println("\n number is greater than 100")
	}

}

// TODO: ----  9、循环 See: https://studygolang.com/articles/11924
func day9() {

	// break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // loop is terminated if i > 5
		}
		fmt.Println("\n ", i)

	}
	fmt.Println("\n line after for loop")

	// continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //
		}
		fmt.Println("\n ", i)

	}
	fmt.Println("\n continue for loop")

	// 分号被省略，并且只有条件存在
	new_i := 0
	for new_i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Println("\n new_i: ", new_i)
		new_i += 2
	}

	// 多重初始化和递增
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("\n %d * %d = %d \n", no, i, no*i)
	}

	// 程序就会一直打印Hello World不会停止。
	/*
		for {
			fmt.Println("Hello World")
		}
	*/

}

// TODO: ----  10、switch 语句 See: https://studygolang.com/articles/11957
func day10() {

	finger := 4
	switch finger {
	case 1:
		fmt.Println("\n Thumb")
	case 2:
		fmt.Println("\n Index")
	case 3:
		fmt.Println("\n Middle")
	case 4:
		fmt.Println("\n Ring")
	case 5:
		fmt.Println("\n Pinky")
	default:
		fmt.Println("\n incorrect finger number")
	}

	// 多表达式判断
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("\n vowel")
	default:
		fmt.Println("\n not a vowel")
	}

	// 无表达式的 switch
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("\n num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("\n num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("\n num is greater than 100")
	}

	// TODO: Fallthrough 语句 (失败之后的处理？)
	/*
	   在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。
	   使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	   让我们写一个程序来理解 fallthrough。
	   我们的程序将检查输入的数字是否小于 50、100 或 200。例如我们输入 75，程序将输出75 is lesser than 100 和 75 is lesser than 200。我们用 fallthrough 来实现了这个功能。
	*/
	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("\n %d is lesser than 50", num)
		fallthrough
	case num < 100:
		fmt.Printf("\n %d is lesser than 100", num)
		fallthrough
	case num < 200:
		fmt.Printf("\n %d is lesser than 200", num)
	}

}

func number() int {
	num := 15 * 5
	return num
}

// TODO: ----  11、数组和切片 See: https://studygolang.com/articles/12121
func day11() {

	var a [3]int // int array with length 3
	fmt.Println("\n ", a)
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println("\n new value", a) // [12,78,50]

	b := [3]int{12, 78, 50}               // short hand declaration to creat array
	fmt.Println("\n new other value ", b) // [12 78 50]

	a = [3]int{12}
	fmt.Println("\n other value ", a) // [12,0,0]

	b = [...]int{12, 78, 50}          // ... makes the complier determine the length
	fmt.Println("\n b new value ", b) // [12 78 50]

	/*
	 not possible since [3]int and [5]int are distinct types
	 var  c [5]int
	 b = c
	*/
	new_a := [...]string{"USA", "China", "India", "Germany", "France"}
	new_b := new_a // new_a copy of new_a is assigned to new_b
	new_b[0] = "Singapore"
	fmt.Println("\n new_a is ", new_a) // [USA China India Germany France]
	fmt.Println("\n new_b is ", new_b) // [Singapore China India Germany France]

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("\n befor passing to function ", num) // [5 6 7 8 8]
	changeLocal(num)                                  // num is passed by value
	// 数组 num 实际上是通过值传递给函数 changeLocal，数组不会因为函数调用而改变。这个程序将输出
	fmt.Println("\n after passing to function ", num, "lenth is ", len(num)) // [5 6 7 8 8] 长度为 5

	// TODO: ---------------------------------------------- 使用range迭代数组 ----------------------------------------------
	c := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
		fmt.Printf("\n %d the element of is %.2f \n\n", i, c[i]) // 0 67.70
	}
	// 通过使用 for 循环的 range 方法来遍历数组。range 返回索引和该索引处的值
	new_c := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for new_i, new_v := range new_c { // range returns both the index and value
		fmt.Printf("\n %d the element of is %.2f \n\n", new_i, new_v) // 0 67.70
		sum += new_v
	}
	fmt.Println("\n sum of all elements of new_c ", sum)
	// 你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行。
	for _, v := range new_c { // ignores index
		fmt.Println("\n ", v)
	}

	// TODO: ---------------------------------------------- 多维数组 ----------------------------------------------
	new_str := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(new_str)

	var new_str1 [3][2]string
	new_str1[0][0] = "apple"
	new_str1[0][1] = "samsung"
	new_str1[1][0] = "microsoft"
	new_str1[1][1] = "google"
	new_str1[2][0] = "AT&T"
	new_str1[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(new_str1)

	// TODO: ---------------------------------------------- 切片 ----------------------------------------------

	// 创建一个切片 --- 带有 T 类型元素的切片由 []T 表示 本身不拥有数据，只是对现有数组的引用
	qie_a := [5]int{76, 77, 78, 79, 80}
	var qie_b []int = qie_a[1:4]               // creat a slice from qie_a[1] to qie_a[3]
	fmt.Println("\n qie_b value is", qie_b)    // [77 78 79]
	new_cc := []int{6, 7, 8}                   // creates and array and returns a slice reference
	fmt.Println("\n new_cc value is ", new_cc) //  [6 7 8]

	/// 切片的修改 (切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。)
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("\n array before", darr) // [57 89 90 82 100 78 67 69 59]
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("\n array after ", darr) //  [57 89 91 83 101 78 67 69 59]
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creat a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("\n array before change 1", numa) // [78 79 80]
	nums1[0] = 100
	fmt.Println("\n array after modification to slice nums1", numa) // 100 79 80]
	nums2[1] = 101
	fmt.Println("\n array after modification to slice nums2", numa) // [100 101 80]

	// 切片的长度和容量
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Println("\n length of slice %d capacity %d", len(fruitslice), cap(fruitslice))                  // length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                                           // re-slicing furitslice till its capacity
	fmt.Println("\n after re-slicing length is ", len(fruitslice), "and capacity is ", cap(fruitslice)) // after re-slicing length is 6 and capacity is 6

	// 使用make创建一个切片
	i := make([]int, 5, 5)
	fmt.Println("\n i value is ", i)

	// 追加切片元素
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("\n cars:", cars, "has old length ", len(cars), "and capacity ", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("\n cars:", cars, "has new length ", len(cars), "and capacity ", cap(cars)) // capacity of cars is doubled to 6
	var names []string                                                                      // zero value of a slice is nil
	if names == nil {
		fmt.Println("\n slice is nil going to append")
		names = append(names, "John", "Sebastion", "Vinay")
		fmt.Println("\n names contents: ", names) // [John Sebastion Vinay]
	}
	veggies := []string{"potatoes", "tomatoes", "brinial"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("\n food: ", food) //  [potatoes tomatoes brinial oranges apples]

	// 切片的函数传递
	nos := []int{8, 7, 6}
	fmt.Println("\n slice before function call", nos) // [8 7 6]
	subtactOne(nos)
	fmt.Println("\n slice after  function call", nos) // [6 5 4]

	// 多维切片
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	/*
	  C C++
	  JavaScript
	  Go Rust
	*/
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Println("\n %s", v2)
		}
		fmt.Println("\n ")
	}

	// 内存优化
	countriesNeeded := countries()
	fmt.Println("\n countriesNeeded value is ", countriesNeeded)

}

// 内存优化
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	needeCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(needeCountries))
	copy(countriesCpy, needeCountries) // copies neededCountries to countriesCpy
	return countriesCpy
}

// 切片的函数传递 mark：结构体类型
type slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

// 切片的函数传递
func subtactOne(number []int) {
	for i := range number {
		number[i] -= 2
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("\n isside function ", num) // [55 6 7 8 8]
}

// 多维数组
func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println("\n %s", v2)
		}
		fmt.Println("\n")
	}
}
