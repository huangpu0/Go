package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
)

// TODO: ---- 35、读取文件 See: https://studygolang.com/articles/14669

func day35()  {

	///1、将整个文件读取到内存 (绝对路径)
	data1, err1 := ioutil.ReadFile("/Users/itrader-dev/Desktop/Go/File/test.txt")
	if err1 != nil {
		fmt.Println("\n File reading error ", err1)
		return
	}
	fmt.Println(" Contents of file:", string(data1))

	/// 2、使用命令行标记来传递文件路径
	fptr := flag.String("fpath", "ios.txt", "file path to read from")
	flag.Parse()
	fmt.Println(" value of fpath is", *fptr)
	data2, err2 := ioutil.ReadFile(*fptr)
	if err2 != nil {
		fmt.Println(" File reading error", err2)
		//return
	}
	fmt.Println(" Contents of file:", string(data2))

	/// 3. 将文件绑定在二进制文件中
	/*
	  虽然从命令行获取文件路径的方法很好，但还有一种更好的解决方法。
	  如果我们能够将文本文件捆绑在二进制文件，岂不是很棒？这就是我们下面要做的事情.
	  有很多包可以帮助我们实现。我们会使用 packr，因为它很简单，并且我在项目中使用它时，没有出现任何问题。
	  第一步就是安装 packr 包。
	  在命令提示符中输入下面命令，安装 packr 包。 (go get -u github.com/gobuffalo/packr/...)
   */


	/// 分块读取文件
	f, err := os.Open(string(data1))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println(" Error reading file:", err)
			break
		}
		fmt.Println(" ",string(b))
	}

}


// TODO: ---- 36、写入文件 See: https://studygolang.com/articles/19475

func day36()  {

	/// 将字符串写入文件
	f, err := os.Create("new_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(" ",l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	/// 将字节写入文件 指定目录(写入bytes  内容 hello bytes)
	f1, err1 := os.Create("/Users/itrader-dev/Desktop/Go/File/bytes")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err1 := f1.Write(d2)
	if err1 != nil {
		fmt.Println(err1)
		f1.Close()
		return
	}
	fmt.Println(" ",n2, "bytes written successfully")
	err1 = f1.Close()
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	/// 将字符串一行一行的写入文件
	f22, err22 := os.Create("lines")
	if err22 != nil {
		fmt.Println(" ", err22)
		f22.Close()
		return
	}

	d22 := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}
	for _, v := range d22 {
		fmt.Fprintln( f22, v)
		if err22 != nil {
			fmt.Println(" error ", err22)
			return
		}
	}
	err22 = f22.Close()
	if err22 != nil {
		fmt.Println(" err22 ", err22)
		return
	}
	fmt.Println(" lines file written successfully")

	/// 追加到文件
	f22_1, err22_1 := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err22_1 != nil {
		fmt.Println(" 1err22_1 ", err22_1)
		return
	}
	newLine := "File handling is easy."
	_, err22_1 = fmt.Fprintln(f22_1, newLine)
	if err22_1 != nil {
		fmt.Println(" 2err22_1 ", err22_1)
		f22_1.Close()
		return
	}
	err22_1 = f22_1.Close()
	if err22_1 != nil {
		fmt.Println(" 3err22_1 ", err22_1)
		return
	}
	fmt.Println(" 22_1 file appended successfully")

	/// 并发写文件
    data := make(chan int)
    done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce36(data, &wg)
	}
	go consume36(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println(" File written successfully")
	} else {
		fmt.Println(" File writing failed")
	}

}

/// 并发写文件
func produce36(data chan int, wg *sync.WaitGroup)  {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}
func consume36(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(" ", err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln( f, d)
		if err != nil {
			fmt.Println(" 11err ", err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(" 22err ", err)
		done <- false
		return
	}
	done <- true

}