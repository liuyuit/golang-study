package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "Japan")
	// 修改
	cityMap["USA"] = "HS"
	fmt.Println("========")
	// 遍历
	printMap(cityMap)
	fmt.Println("========")
	printMap(cityMap)
	fmt.Println("========")
	changeMap(cityMap)
	printMap(cityMap)
}

// map 是引用传值
func printMap(cityMap map[string]string) {
	// 遍历
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
}

func changeMap(cityMap map[string]string) {
	cityMap["English"] = "England"
}
