package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var mu sync.Mutex


func main() {
	ch:=make(chan int)
	wg.Add(2)
	go Process(ch)
	ch<-1		//堵塞
	go Process(ch)

	//go add()go add()

	wg.Wait()
	fmt.Println("x的值为：",x)
}

func Process(ch chan int ){
	var receive int //用receive来接受
	add()		//调用add
	time.Sleep(time.Second)

	receive=<-ch	//接收ch的值，疏通管道
	fmt.Println("成功：",receive)
}

func add() {
	for i := 0; i < 50000; i++ {
		x = x + 1

		//mu.Lock()mu.Unlock()，未使用锁

	}
	wg.Done()
}
