package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: ---- 23、缓冲信道和工作池 See: https://studygolang.com/articles/12512

func write(ch chan int)  {

	for i := 0; i < 5; i ++ {
		ch <- i
		fmt.Println("\n successfully wrote", i , "to ch")
	}
	close(ch)
}


/// WaitGroup
func process( i int, wg *sync.WaitGroup)  {
	fmt.Println("\n started Groutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf(" Goroutine %d ended\n", i)
	wg.Done()
}

/// 工作池的实现
var jobs    = make(chan Job, 10)
var results = make(chan Result, 10)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	 job         Job
	 sumofdigits int
}

func digits23(number int) int  {
	sum := 0
	no  := number
	for no != 0 {
		digit := no % 10
		sum   += digit
		no   /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup)  {
	for job := range jobs {
		output := Result{job, digits23(job.randomno)}
		results <- output
	}
	wg.Done()
}

// Go 协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
// 作业分配给工作者
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf(" Job id %d, input random no %d , sum of digits23 %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}


func day23()  {

	/// 创建一个缓冲信道
	ch := make(chan string, 2)
	ch <- "naveen"
    ch <- "paul"
  //  ch <- "steve" 加入会产生死锁
    fmt.Println("\n ", <- ch)
    fmt.Println("\n ", <- ch)

    /// 在向缓冲信道写入数据时，什么时候会发生阻塞。
	new_ch := make(chan int, 2)
    go write(new_ch)
	time.Sleep(2 * time.Second)
	for v := range new_ch {
		fmt.Println(" read value ", v, "from new_ch")
		time.Sleep(2 * time.Second)
	}

	/// 长度 vs 容量
	/*
	 缓冲信道的容量是指信道可以存储的值的数量。我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
	 缓冲信道的长度是指信道中当前排队的元素个数。
	*/
	ch_str := make(chan string, 3)
	ch_str <- "naveen"
	ch_str <- "paul"
	fmt.Println("\n capacity is", cap(ch_str)) // 3
	fmt.Println(" length is", len(ch_str)) // 2
	fmt.Println(" read value", <-ch_str) // naveen
	fmt.Println(" new length is", len(ch_str)) // 1


	/// WaitGroup
    no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go  process(i, &wg)
	}
    wg.Wait()
	fmt.Println(" All go routines finished executing")


	/// 工作池的实现
    startTime := time.Now()
    noOfJobs  := 100
    go allocate(noOfJobs)
    done := make(chan bool)
    go result(done)
    noOfWorkers := 10
    createWorkerPool(noOfWorkers)
    <-done
    endTime := time.Now()
    diff := endTime.Sub(startTime)
    fmt.Println(" total time taken ", diff.Seconds(), "seconds")


}


// TODO: ---- 24、Select See: https://studygolang.com/articles/12522

/*
  select 语句用于在多个发送/接收信道操作中进行选择。
  select 语句会一直阻塞，直到发送/接收操作准备就绪。
  如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。
  该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作
*/
func server1(ch chan string)  {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string)  {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}


/// select 的应用
func process_select(ch chan string)  {
	time.Sleep( 10500 * time.Millisecond)
	ch <- "process successful"
}

/// 随机选取
func server3(ch chan string)  {
	ch <- "from server3"
}

func server4(ch chan string)  {
	ch <- "from server4"
}

func day24()  {

	/// 什么是 select？
	output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println("\n s1 value is", s1)
	case s2 := <-output2:
		fmt.Println("\n s2 value is", s2)
	}
    // print s2 value is from server2

    /// select 的应用
    ch := make(chan string)
    go process_select(ch)
	for  {
		/// 大约'10.5s'之后输入process successful 不在执行下面的代码
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			 fmt.Println(" received value: ", v)
			 return //
			 default:
			 fmt.Println(" no value recevied")
		}
	}


	// TODO:  参考下面输出 注释掉for循环

	/// 死锁与默认情况
	/*
	error!!! fatal error: all goroutines are asleep - deadlock!
	+ 默认情况不会发生死锁 如果 select 只含有值为 nil 的信道 也会走默认（var new_ch chan string）
	*/
	new_ch := make(chan string)
	select {
	case <-new_ch:
	default:
		fmt.Println(" defalut case exected")
	}


	/// 随机选取
	output3 := make(chan string)
	output4 := make(chan string)
	go server3(output3)
	go server4(output4)
	time.Sleep(1 * time.Second)
	select {
	case s3 := <-output3:
		fmt.Println("\n s3 value is", s3)
	case s4 := <-output4:
		fmt.Println("\n s4 value is", s4)
	}

    /// 空slect
    /*
     error!! fatal error: all goroutines are asleep - deadlock!
     select {
		}
    */



}