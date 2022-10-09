package main

import "fmt"

func main() {
	// 声明
	var a1 [5]float32

	// 初始化
	var a2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	a3 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// 不定长度
	a4 := [...]float32{1.2}

	//  将索引为 1 和 3 的元素初始化
	a5 := [5]float32{1: 2.0, 3: 7.0}

	// 二维数组
	a6 := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}

	fmt.Println(a1, a2, a3, a4, a5, a6)

	slice()
}

// 切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
func slice() {
	// make([]T, length, capacity)
	s1 := make([]int, 4)

	s2 := []int{1, 2, 3}

	// s2 的引用
	s3 := s2[:]

	// arr[startIndex:endIndex] ，将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片
	s4 := s2[1:]

	// len() 获取长度  cap() 获取最大容量
	fmt.Println(s1, s2, s3, s4, len(s1), cap(s1))

	// nil 空切片
	var s5 []int

	fmt.Println(s5, s5 == nil)

	// 追加元素
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println(s1, len(s1), cap(s1))

	s5 = append(s5, 1, 2)
	fmt.Println(s5, len(s5), cap(s5))

	// 创建两倍 s2 容量的新切片，并拷贝
	s6 := make([]int, len(s2), (cap(s2))*2)
	copy(s6, s2)
	fmt.Println(s6)
}
