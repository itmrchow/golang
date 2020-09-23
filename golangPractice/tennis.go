package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func playTennis() {
	//建立無緩衝的channel
	court := make(chan int)

	//等待兩個goroutine
	wg.Add(2)

	//啟動
	go player("Nick", court)
	go player("Jeff", court)

	//發球
	court <- 1

	wg.Wait()
}

//player 模擬打球
func player(name string, court chan int) {
	//函式退出時呼叫Done來通知main函式工作完成
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			//通道關閉代表結束
			fmt.Printf("Player %s Won\n", name)
			return
		}

		//選擇隨機數，來判斷是否沒接到
		n := rand.Intn(100)
		if n%5 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//關閉通道
			close(court)
			return
		}

		//顯示擊球
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball

	}

}
