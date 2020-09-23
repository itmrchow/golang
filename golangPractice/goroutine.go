package main

import (
	"log"
	"os"
	"sync"
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