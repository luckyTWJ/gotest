package main

import "fmt"

func main() {
	//声明map key string; value int
	var m map[string]int
	if m == nil {
		fmt.Println("m 是一个空map")
	}
	//第一种 初始化分配空间
	m = make(map[string]int, 10)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)
	//第二种
	map2 := make(map[string]string)
	map2["one"] = "java"
	map2["two"] = "c++"
	map2["three"] = "go"

	fmt.Println(map2)

	//	第三种
	map3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "go",
	}
	fmt.Println(map3)

	fmt.Println("-------------------")
	map_test()
}

func map_test() {
	cityMap := make(map[string]string)
	cityMap["china"] = "北京"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "new york"
	//	遍历
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
	// 删除
	delete(cityMap, "usa")
	fmt.Println(cityMap)

}
