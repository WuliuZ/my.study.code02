package main

import (
	"fmt"
	"sync"
)


func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		for i := 1; i < 101; i++ {
			ch <- 1
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1:",i)
			}
		}
		wg.Done()
	}()

	go func() {

		for i := 1; i < 101; i++ {
			<- ch
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2:",i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
//两个匿名函数中都有对管道的读写操作，读写时就会引发堵塞
//一个堵塞时另一个就疏通，然后执行if语句，if语句同时只有一个成立
//完成轮流打印



