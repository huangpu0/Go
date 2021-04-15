package main

import "fmt"

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
