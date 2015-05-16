package main

import (
	"fmt"

	"./mongodb"
)

var stockcodeArray []string

func init() {
	stockcodeArray = parseFileToStringArrayAll()
}

func main() {

	var menu int

	fmt.Println("主要股东 ----> 1")
	fmt.Println("流通股东 ----> 2")
	fmt.Println("\n请选择要执行的功能：")
	fmt.Scanln(&menu)

	switch {
	case menu == 1:
		mongodb.StoreStockMainHolderModel(stockcodeArray)
	case menu == 2:
		mongodb.StoreStockPublicHolderModel(stockcodeArray)
	}

}
