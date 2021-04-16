package main

import (
	"fmt"
	"unicode/utf8"
)

// TODO: ----  可变参数函数 See: https://studygolang.com/articles/12173
func day12() {

	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	// 给可变参数传入切片
	nums := []int{89, 90, 95}
	// error find(89, nums)
	find(89, nums...)

	// 不直观错误
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println("\n welcome value is ", welcome) // 如果你认为它输出 [Go world] 。恭喜你！你已经理解了可变参数函数和切片

}

func find(num int, nums ...int) {
	fmt.Println("\n type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println("\n ", num, " found at index ", i, " in ", nums)
			found = true
		}
	}
	if !found {
		fmt.Println("\n not found in ", nums)
	}
	fmt.Println("\n")

}

// 不直观错误
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println("\n s value is ", s)
}



// TODO: ----  Maps See: https://studygolang.com/articles/12251
func day13()  {

	// diction? key:value?
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("\n map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"]  = 9000
	fmt.Println("\n personSalary map contents: ", personSalary) //  map[jamie:15000 mike:9000 steve:12000]

	// 添加元素
	personSalary = map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("\n personSalary map contents: ", personSalary) // map[jamie:15000 mike:9000 steve:12000]

	// 获取map中的元素
	employee := "jamie"
	fmt.Println("\n Salary of ", employee, " value is ", personSalary[employee]) // Salary of jamie is 15000
	fmt.Println("\n Salary of joe is ", personSalary["joe"]) // Salary of joe is 0

	// 判断map是否含有key
	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("\n Salary of", newEmp, "vaule is ", value)
	}else {
		fmt.Println("\n newEmp vaule is ", newEmp, "not found")
	}
    // 遍历key value
	for key, value := range personSalary {
		fmt.Println("\n personSalary[%s] = %d\n", key, value)
	}

	// 删除map中的元素
	fmt.Println("\n map before deletion ", personSalary) // map[jamie:15000 mike:9000 steve:12000]
    delete(personSalary, "steve")
    fmt.Println("\n map after deletion ", personSalary) // map[jamie:15000 mike:9000]

    // 获取map的长度
    fmt.Println("\n length is ", len(personSalary)) //  length is  2


    // ------------------------------------- Map -------------------------------------
    /*
      和 slices 类似，map 也是引用类型。
      当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。
      因此，改变其中一个变量，就会影响到另一变量。
    */
	personSalary["steve"] = 12000
	fmt.Println("\n Original person salary ", personSalary)
	newPersonSalay := personSalary
	newPersonSalay["mike"] = 18000
	fmt.Println("\n Person salary changed", personSalary)


    // Map 的相等性
    map1 := map[string]int{
    	"one": 1,
    	"two": 2,
	}

	map2 := map1
	fmt.Println("\n map1 value is ", map1, " map2 vaule is ", map2)
	// error: -- invalid operation: map1 == map2 (map can only be compared to nil)
	/*
	if map1 == map2 {
		fmt.Println("\n true true true true")
	}
    */

}


// TODO: ---- 字符串 See: https://studygolang.com/articles/12261
func day14()  {

	name := "Hello World"
	fmt.Println("\n name value is ", name)
	printBytes(name)
	fmt.Println("\n")
	printChars(name)
	fmt.Println("\n")

	// rune 它也是 int32 的别称
	printRuneChars(name)
	fmt.Println("\n")

	// 字符串的for range 循环
	printCharsAndBytes(name)

	// 用字节切片构造字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0XA9}
	str := string(byteSlice)
	fmt.Println("\n str value is ", str)
	byteSlice = []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println("\n str value is ", str)

	// 用 rune 切片构造字符串
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str = string(runeSlice)
    fmt.Println("\n str value is ", str)

    // 字符串长度
    word1 := "Hello,World"
    length(word1)
	word2 := "Pets"
	length(word2)

	// 字符串是不可变的
	h := "hello"
	// fmt.Println("\n mutate value is ", mutate(h)) cannot assign to s[0] (strings are immutable)
	fmt.Println("\n mutate value is ", new_mutate([]rune(h)))

}

// !!! Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。
/*
func mutate(s string) string  {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	return s
}
*/
func new_mutate(s []rune) string  {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	return string(s)
}

// 字符串长度
func length(s string)  {
	fmt.Println("\n length of %s is %d\n", s, utf8.RuneCountInString(s))
}

// 单独获取字符串每一个字节
func printBytes(s string)  {
	for i := 0; i < len(s); i ++ {
		fmt.Println("\n %x", s[i])
	}
}

func printChars(s string)  {
	for i := 0; i < len(s); i ++ {
		fmt.Println("\n %c", s[i])
	}
}

// rune
func printRuneChars(s string)  {
	runes := []rune(s)
	for i := 0; i < len(runes); i ++ {
		fmt.Println("\n %c", s[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("\n %c starts at byte %d\n", rune, index)
	}
}