package main

import (
	"fmt"
	"unsafe"
)

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	i int
	f float64
	b bool
	s string
)

func main() {
	// 注释
	/* 注释 */
	const hello string = "hello"
	var num, num2 int = 123, 456

	// 值交换, 两个变量的类型必须是相同
	num, num2 = num2, num

	// := 是一个声明语句，类型由编译器推断
	num3 := 789

	// *指针变量 &返回变量存储地址
	var ptr *int = &num

	fmt.Println(fmt.Sprintf("%s, World! %d %d %d", hello, num, num2, num3), ptr)

	// 如果没有初始化，则变量默认为零值
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	const (
		helloLength = len(hello)
		helloSize   = unsafe.Sizeof(hello)
	)

	println(helloLength, helloSize)

	if hello == "hello" {
		iotaFunc()
	} else {
	}

	selectFunc()
}

func iotaFunc() {
	// iota 在 const关键字出现时将被重置为 0
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		z        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, z)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	// i=1：左移 0 位，不变仍为 1。
	// j=3：左移 1 位，变为二进制 110，即 6。
	// k=3：左移 2 位，变为二进制 1100，即 12。
	// l=3：左移 3 位，变为二进制 11000，即 24
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

func selectFunc() {
	ch := make(chan int, 1)

	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
}
