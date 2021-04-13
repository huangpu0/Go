package main

import (
	"fmt"
)

// TODO: ----  if-else 语句 See: https://studygolang.com/articles/11902
func day8()  {

	num := 10
	if num % 2 == 0 { //checks if number is even
		fmt.Println("\n the number is even")
	} else {
        fmt.Println("\n the number is odd")
	}

	if num := 10; num %2 == 0 { //checks if number is even
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


// TODO: ----  循环 See: https://studygolang.com/articles/11924
func day9()  {

	// break
	for i := 1; i <= 10; i ++ {
		if i > 5 {
			break // loop is terminated if i > 5
		}
		fmt.Println("\n ", i)

	}
    fmt.Println("\n line after for loop")

    // continue
	for i := 1; i <= 10; i++ {
		if i %2 == 0 {
			continue //
		}
		fmt.Println("\n ", i)

	}
	fmt.Println("\n continue for loop")

	// 分号被省略，并且只有条件存在
	new_i := 0
	for new_i <= 10 {  //semicolons are ommitted and only condition is present
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


// TODO: ----  switch 语句 See: https://studygolang.com/articles/11957
func day10()  {

	finger := 4
	switch finger  {
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
	switch  {
	case num >= 0 && num <= 50:
        fmt.Println("\n num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("\n num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("\n num is greater than 100")
	}

	// Fallthrough 语句 (失败之后的处理？)
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