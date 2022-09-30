package main

import (
	"fmt"
	"math"
)

// 函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

// 闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/* 定义结构体 */
type Circle struct {
	radius float64
}

// Go 语言中同时有函数和方法
// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

	nextNumber := getSequence()
	fmt.Println(nextNumber()) // 1
	fmt.Println(nextNumber()) // 2
	fmt.Println(nextNumber()) // 3

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
