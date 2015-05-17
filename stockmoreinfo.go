package main

import (
	"fmt"
	"strconv"
	"strings"

	"./mongodb"
)

var stockcodeArray []string

func init() {
	stockcodeArray = parseFileToStringArrayAll()
}

func main() {

	var cmd string

	fmt.Println("主要股东 ----> 1")
	fmt.Println("流通股东 ----> 2")
	fmt.Println("\n请选择要执行的功能，多选项时请使用一个逗号分隔：")
	fmt.Scanln(&cmd)

	indicatorArr := strings.Split(cmd, ",")

	for _, indicator := range indicatorArr {
		menu, _ := strconv.Atoi(indicator)
		switch {
		case menu == 1:
			fmt.Println("开始执行 ----> 主要股东")
			mongodb.StoreStockMainHolderModel(stockcodeArray)
		case menu == 2:
			fmt.Println("开始执行 ----> 流通股东")
			//stockcodeArray = stockcodeArray[2071:len(stockcodeArray)]
			mongodb.StoreStockPublicHolderModel(stockcodeArray)
		}
	}

}
