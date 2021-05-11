package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: ---- 缓冲信道和工作池 See: https://studygolang.com/articles/12512

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

