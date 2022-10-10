package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 关闭通道
	close(c)
}

func main() {
	// 通道（channel）是用来传递数据的一个数据结构
	// 可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯
	// 发送 ch <- v，接收 v := <-ch
	// 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
	c := make(chan int, 10)

	// 通过 go 开启 goroutine (轻量级线程)
	// 同一个程序中的所有 goroutine 共享同一个地址空间
	go fibonacci(cap(c), c)

	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
