package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"sync"
)

// TODO: ---- 28、多态 See: https://studygolang.com/articles/12598

/// 使用接口实现多态
type Income28 interface {
	calculate28() int
	source28()    string
}

type FixedBilling28 struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial28 struct {
	projectName string
	noOfHours  int
	hourlyRate int
}

func (fb FixedBilling28) calculate28() int {
	return fb.biddedAmount
}

func (fb FixedBilling28) source28() string {
	return fb.projectName
}

func (tm TimeAndMaterial28) calculate28() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial28) source28() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income28) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("\n Income28 From %s = $%d\n", income.source28(), income.calculate28())
		netincome += income.calculate28()
	}
	fmt.Printf("\n Net income28 of organisation = $%d", netincome)
}


/// 新增收益流
type Advertisement28 struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement28) calculate28() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement28) source28() string {
	return a.adName
}

func day28()  {

	/// 使用接口实现多态
	project1 := FixedBilling28{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling28{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial28{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	/// incomeStreams := []Income28{project1, project2, project3}
	/// calculateNetIncome(incomeStreams)

    /// 新增收益流
	bannerAd := Advertisement28{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd  := Advertisement28{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income28{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)

}

// TODO: ---- 29、Defer See: https://studygolang.com/articles/12719

/*
  defer 语句的用途是：含有 defer 语句的函数，
  会在该函数将要返回之前，调用另一个函数。
*/

/// 什么是 defer？
func finished29()  {
	fmt.Println(" Finished finding largest") // 3
}

func largest29(nums []int)  {
	defer finished29()
	fmt.Println("\n Started finding largest") // 1
	max := nums[0]
	for _, v := range nums{
		if v  > max {
			max = v
		}
	}
	fmt.Println(" Largest number in", nums, "is", max) // 2
}

/// 延迟方法
type person29 struct {
	firstName string
	lastName  string
}

func (p person29)fullName()  {
	fmt.Printf("\n %s %s", p.firstName, p.lastName) // 2-2
}

/// 实参取值
func peintA29(a int)  {
	fmt.Println("\n value of a in deferred function", a) // 3-2
}

/// defer 的实际应用
type rect29 struct {
	length int
	width  int
}

func (r rect29) area29(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf(" rect %v's length should be greater than zero\n", r) // 5-3
		// wg.Done() 重写优化
		return
	}
	if r.width < 0 {
		fmt.Printf(" rect %v's width should be greater than zero\n", r) // 5-2
		// wg.Done() 重写优化
		return
	}
	area := r.length * r.width
	fmt.Printf("\n rect %v's area %d\n", r, area) // 5-1
	// wg.Done() 重写优化
}


func day29()  {

	/// 什么是 defer？
	nums := []int{78, 109, 2, 563, 300}
	largest29(nums)

	/// 延迟方法
	p := person29{
		firstName: "John",
		lastName: "Smith",
	}
	defer p.fullName()
	fmt.Println("\n Welcome ") // 2-1

	/// 实参取值
	a := 5
	defer peintA29(a)
	a = 10
	fmt.Println("\n value of a before deferred function call", a) // 3-1

	/// defer 栈
	name := "Naveen"
	fmt.Printf("\n Orignal String: %s\n", string(name)) // 4-1
	fmt.Printf("\n Reversed String: ") // 4-2
	for _, v := range []rune(name) {
		defer fmt.Printf(" %c", v) // 4-3 (neevaN)
	}

	/// defer 的实际应用
	var wg sync.WaitGroup
	r1 := rect29{-67, 89}
	r2 := rect29{5, -67}
	r3 := rect29{8, 9}
	rects := []rect29{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area29(&wg)
	}
	wg.Wait()
	fmt.Println(" All go routines finished executing") // 5-4
	
}



// TODO: ---- 30、错误处理 See: https://studygolang.com/articles/12724

func day30()  {

	/// 示例错误
	f0, err0 := os.Open("/test.txt")
	if err0 != nil {
		fmt.Println("", err0)
		return
	}
	fmt.Println(f0.Name(), "opened successfully")

	/// 1. 断言底层结构体类型，使用结构体字段获取更多信息
	f, err1 := os.Open("/test.txt")
	if err1, ok := err1.(*os.PathError); ok {
		fmt.Println("\n File at path ", err1.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), " opened successfully")

	/// 2. 断言底层结构体类型，调用方法获取更多信息
	addr, err2 := net.LookupHost("golangbot123.com")
	if err2, ok := err2.(*net.DNSError); ok {
		if err2.Timeout() {
			fmt.Println(" operation timed out")
		} else if err2.Temporary() {
			fmt.Println(" temporary error")
		} else {
			fmt.Println(" generic error: ", err2) // 错误输出
		}
		return
	}
	fmt.Println(addr)

	/// 3. 直接比较
	files, err3 := filepath.Glob("[")
	if err3 != nil && err3 == filepath.ErrBadPattern {
		fmt.Println(	" ",err3)
		return
	}
	fmt.Println("matched files", files)

	/// 不可忽略错误
	files4, _ := filepath.Glob("[")
	fmt.Println(" matched files", files4)

}



// TODO: ---- 31、自定义错误 See: https://studygolang.com/articles/12784

/// 创建一个计算圆半径的简单程序，如果半径为负，它会返回一个错误。
func circleArea31(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New(" Area calculation failed, radius is less than zero")
		// 添加更多错误信息
		return 0, fmt.Errorf(" Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

/// 使用结构体类型和字段提供错误的更多信息
type areaError31 struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError31) Error() string {
	panic("implement me")
}

func (e *areaError31) Error31() string {
	return e.err
}

func (e *areaError31) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError31) widthNegative() bool {
	return e.width < 0
}

func rectArea31(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError31{err, length, width}
	}
	return length * width, nil
}



func day31()  {

	/*
	  使用 New 函数创建自定义错误
	  使用 Error 添加更多错误信息
	  使用结构体类型和字段，提供更多错误信息
	  使用结构体类型和方法，提供更多错误信息
	*/

	radius := -20.0
	area, err := circleArea31(radius)
	if err != nil {
		fmt.Println("\n error is ",err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)


	/// 使用结构体类型和字段提供错误的更多信息
	length, width := -5.0, -9.0
	area1, err1 := rectArea31(length, width)
	if err1 != nil {
		if err1, ok := err1.(*areaError31); ok {
			if err1.lengthNegative() {
				fmt.Printf("\n error: length %0.2f is less than zero\n", err1.length)
			}
			if err1.widthNegative() {
				fmt.Printf("\n error: width %0.2f is less than zero\n", err1.width)
			}
			return
		}
		fmt.Println(" ", err1)
		return
	}
	fmt.Println("\n area of rect", area1)

}


























