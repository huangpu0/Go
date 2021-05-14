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
