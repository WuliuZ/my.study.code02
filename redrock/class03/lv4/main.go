package main

import (
	"fmt"
	"time"
)

func main() {
	over := make(chan bool,1)		//声明无缓存的管道over，改为缓存为1
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		//子协程没有主协程运行快，
		// 导致子协程还没有执行就又开始了下一轮for循环
		time.Sleep(time.Second/100)
		if i == 9 {
			over <- true	//死锁,导致堵塞可以加个缓存
		}
	}
	<-over
	fmt.Println("over!!!")
}
