package main

import (
	"fmt"
)


// TODO: ---- 6、函数 See: https://studygolang.com/articles/11892
func day6()  {


	functionname()

	// 单值返回
	price, no := 90, 6
	totalPrice := calulateBill(price, no)
	fmt.Println("\n Total price is", totalPrice)

	// 多值返回
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Println("\n Area %f Perimeter %f", area, perimeter)
	new_area, new_perimeter := new_rectProps(10.8, 5.6)
	fmt.Println("\n NEW_Area %f NEW_Perimeter %f", new_area, new_perimeter)

	// 空白符 '_' 可以表示任何类型的任何值
	area1, _ := rectProps(10.8, 5.6) // 返回值周长被丢掉
	fmt.Println("\n Area1 %f new is ", area1)

}

// mark: ------ 无返回值
func functionname()  {
	// 不需输入参数、且无返回值
}

// mark: ------ 单返回值
func calulateBill(price int, no int) int  {
	var totalPrice = price + no // 商品总价 = 商品单价 * 数量
	return totalPrice // 返回总价
}

// mark: ------ 多返回值
func rectProps(length, width float64) (float64, float64)  {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func new_rectProps(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}


// TODO: ---- 7、包 See: https://studygolang.com/articles/11893
/*
   到目前为止，我们看到的 Go 程序都只有一个文件，文件里包含一个 main 函数和几个其他的函数。
   在实际中，这种把所有源代码编写在一个文件的方法并不好用。
   以这种方式编写，代码的重用和维护都会很困难。而包（Package）解决了这样的问题。
 */

func day7()  {

	var rectLen, rectWdth float64 = 6, 7
	fmt.Println("Geometrical shape properties", rectLen, rectWdth)
	/*Area function of rectangle package used*/
	//fmt.Println("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWdth))

}