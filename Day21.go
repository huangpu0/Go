package main

import (
	"fmt"
	"time"
)

// TODO: ---- Go协程 See: https://studygolang.com/articles/12342

/// 启动一个Go协程
func hello()  {
	fmt.Println("\n Hello world goroutine")
}


/// 启动多个Go协程
func numbers()  {
	for i := 1; i <= 5; i ++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf(" %d", i)
	}
}

func alphabets()  {
	for i := 'a'; i <= 'e'; i ++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf(" %c", i)
	}
}

func day21()  {

    // Go 协程是什么
	/*
	   Go协程是与其他函数或方法一起并发运行的函数或方法。Go协程可以看作轻量级线程。与线程相比。创建一个Go协程的成本很小。
	   因此在 Go 应用中，常常会看到有数以千计的 Go 协程并发地运行。
	*/

    /// 启动一个Go协程
	go hello()
    /*
     不加 '''time.Sleep(1 * time.Second)''' 直接输出 '''fmt.Println("\n main function")'''
     该程序只会输出文本 main function。我们启动的 Go 协程究竟出现了什么问题？要理解这一切，我们需要理解两个 Go 协程的主要性质。
	 启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
     如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	 现在你应该能够理解，为何我们的 Go 协程没有运行了吧。在第 11 行调用了 go hello() 之后，程序控制没有等待 hello 协程结束，立即返回到了代码下一行，打印 main function。接着由于没有其他可执行的代码，Go 主协程终止，于是 hello 协程就没有机会运行了。
    */
    time.Sleep(1 * time.Second)
	fmt.Println("\n main function")

	/// 启动多个Go协程
	go numbers()
    go alphabets()
	//执行后输出 1 a 2 3 b 4 c 5 d e
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\n mian terminated")


}


// TODO: ---- 信道 See: https://studygolang.com/articles/12402

/// 用信道重写上述代码
func new_hello(done chan bool) {
	fmt.Println("\n Hello world goroutine")
	/// hello 协程里加入休眠函数，以便更好地理解阻塞的概念。
	time.Sleep(4 * time.Second)
	fmt.Println("\n hello go routine awake and going to write to done")
	done <- true
}

/// 信道的另一个示例
func calcSquares(number int, squareop chan int) {
	sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit
	//	number /= 10
	//}
	//squareop <- sum

	/// 代码复用
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit * digit
	//	number /= 10
	//}
	//cubeop <- sum

	/// 代码复用
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

/// 单向信道
func sendData(sendch chan <- int)  {
	sendch <- 10
}

/// 关闭信道和使用 for range 遍历信道
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

/// 我们可以使用 for range 循环，重写信道的另一个示例这一节里面的代码，提高代码的可重用性。
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func day22()  {

    /*
	什么是信道？ （chan T 表示'T'类型的信道）
	信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收。
	*/
	var a chan int
	if a == nil {
		fmt.Println("\n channel a is nil, going to define it") // channel a is nil, going to define it
		a = make(chan int)
		fmt.Printf("\n Type of a is %T", a) // Type of a is chan int
	}

	/// 通过信道进行发送和接送
	/*
	data := <- a // 读写信道 a
	a <- data    // 写入信道 a
    */

	/// 信道的代码示例
	done := make(chan bool)
	fmt.Println("\n Main going to call hello go goroutine")
	go new_hello(done)
	<-done
	fmt.Println("\n main function")


    /// 信道的另一个示例
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println(" Final output", squares + cubes)

	/// 死锁
	/*
	 error!!!
	 fatal error: all goroutines are asleep - deadlock!
	 ch := make(chan int)
	 ch <- 5
	*/

	/// 单向信道
	/*
	我们创建了唯送（Send Only）信道 sendch。chan<- int 定义了唯送信道，因为箭头指向了 chan。在第 12 行，我们试图通过唯送信道接收数据，于是编译器报错：
	main.go:11: invalid operation: <-sendch (receive from send-only type chan<- int)

	 sendch := make(chan <- int)
	 go sendData(sendch)
	 fmt.Println("\n "<-sendch)
	*/
    cha1 := make(chan int)
    go sendData(cha1)
	fmt.Println("\n ", <-cha1)


    /// 关闭信道和使用 for range 遍历信道
	ch := make(chan int)
	go producer(ch)
	//for {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("\n Received ", v, ok) //  Received  '0 ~~ 10' true
	//}

	/// for range 循环用于在一个信道关闭之前，从信道接收数据。
	/// 接下来我们使用 for range 循环重写上面的代码。
	for v := range ch {
		fmt.Println("\n Received ",v) //  Received  '0 ~~ 10'
	}



}