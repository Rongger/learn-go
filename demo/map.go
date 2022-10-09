package main

import "fmt"

func main() {
	var map1 map[string]string = make(map[string]string)

	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India "] = "新德里"

	fmt.Println(map1)

	// 遍历
	for country := range map1 {
		fmt.Println(country, "首都是", map1[country])
	}

	// 查询
	capital, ok := map1["American"]
	fmt.Println(capital, ok)

	// 删除
	delete(map1, "Japan")
}
