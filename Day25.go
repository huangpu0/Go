package main

import (
	"fmt"
	"sync"
)

// TODO: ---- Mutex See: https://studygolang.com/articles/12598

/// 含有竞态条件的程序
var x = 0
func increment(wg *sync.WaitGroup)  {
	x = x + 1
	wg.Done()
}

/// 使用Mutex 修复版
func new_increment(wg *sync.WaitGroup, m *sync.Mutex)  {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

/// 使用信道处理竟态条件
func increment11(wg *sync.WaitGroup, ch chan bool)  {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}


func day25()  {

	/// Mutex 定义了两个方法：Lock 和 Unlock。
	/// 所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。


	/// 含有竞态条件的程序
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
	//	go increment(&w) （未使用Mutex）
		/// 使用Mutex
		go  new_increment(&w, &m)
	}
	w.Wait()
	fmt.Println("\n final value of x", x) // 不加锁因为竞争关系最终结果可能不是1000 使用Mutex 为1000


    /// 使用信道处理竟态条件
    ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment11(&w, ch)
	}
	w.Wait()
	fmt.Println(" final value of x", x) // 为2000 变量x运行2次

	// TODO: 总体说来，当 Go 协程需要与其他协程通信时，可以使用信道。而当只允许一个协程访问临界区时，可以使用 Mutex。

}


// TODO: ----  结构体取代类 See: https://studygolang.com/articles/12630

/// Go 不支持类，而是提供了结构体。结构体中可以添加方法。结构体取代类 (新建文件夹opp/NewEmployee/NewEmployee.go)
/// 包的导入


func day26()  {


}


// TODO: ----  组合取代继承 See: https://studygolang.com/articles/12680

/// 通过嵌套结构体进行组合
type author27 struct {
	firstName string
	lastName  string
	bio       string
}

func (a author27)fullName() string {
	return fmt.Sprintf(" %s %s", a.firstName, a.lastName)
}

type post27  struct {
	 title   string
	 content string
	 author27
}

func (p post27)details()  {
	fmt.Println("\n Title: ", p.title)
	fmt.Println(" Content: ", p.content)
	fmt.Println(" Author: ", p.author27.fullName())
	fmt.Println(" Bio: ", p.bio) // 不用 'p.author27'
}

/// 结构体切片的嵌套
type website27 struct {
	posts []post27
}

func (w website27)contents27()  {
	fmt.Println("\n Content of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}
func day27()  {

	/// 通过嵌套结构体进行组合
	author1 := author27{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	post1 := post27{
		 "Inheritance in Go",
		 "Go supports composition instead of inheritance",
		 author1,
	}

	post1.details()

	/// 结构体切片的嵌套
	post2 := post27{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post27{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website27{
		posts: []post27{post1, post2, post3},
	}
	w.contents27()


}