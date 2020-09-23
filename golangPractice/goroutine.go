package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func goroutineTest() {

	//os.Stdout標準化輸出到console
	logger := log.New(os.Stdout, "", 0)

	//在goroutines進行同步
	//wg集合goroutine
	var wg sync.WaitGroup

	//add goroutine 1
	wg.Add(1)
	go func() {
		//defer 延遲執行於return前
		defer wg.Done()
		logger.Println("Print from goroutine 1")
	}()

	//add goroutine 2
	wg.Add(1)
	go func() {
		//defer 延遲執行於return前
		defer wg.Done()
		logger.Println("Print from goroutine 2")
	}()

	//add goroutine 2
	wg.Add(1)
	go func() {
		//defer 延遲執行於return前
		defer wg.Done()
		logger.Println("Print from goroutine 3")
	}()

	logger.Println("Print from main")

	//等goroutines
	wg.Wait()

}

// channel是同步的，要等到資料到齊
func channelTest() {
	//建立通道
	message := make(chan string)

	// init a goroutine
	go func() {
		//寫入通道
		message <- "Hello from channel"
	}()

	//讀出通道
	msg := <-message
	fmt.Println(msg)
}

// buffer channel是非同步的，不用等到資料到齊
func bufferedChannelTest() {
	logger := log.New(os.Stdout, "", 0)

	var wg sync.WaitGroup

	// Make a buffered channel.
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Println("Print from goroutine", <-ch)
		}()
	}

	logger.Println("Print from main")
	wg.Wait()

}

//指定channel方向
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

var m *sync.Mutex

func mutexTest() {
	m = new(sync.Mutex)
	go read(1)
	go read(2)
	time.Sleep(time.Second) // 让goroutine有足够的时间执行完
}

func read(i int) {
	fmt.Println(i, "begin lock")
	m.Lock()
	fmt.Println(i, "in lock")
	m.Unlock()
	fmt.Println(i, "unlock")
}
