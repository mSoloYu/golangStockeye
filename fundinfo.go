package main

import (
	"fmt"
	"strconv"
	"strings"

	"./mongodb"
)

var fundcodeArray []string

func init() {
	fundcodeArray = parseFileToStringArrayAll(typefund)
}

func main() {

	var cmd string

	fmt.Println("基金资料 ----> 1")
	fmt.Println("基金持仓 ----> 2")
	fmt.Println("")
	fmt.Println("退出程序 ----> 回车")
	fmt.Println("\n请选择要执行的功能，多选项时请使用一个逗号分隔：")
	fmt.Scanln(&cmd)

	indicatorArr := strings.Split(cmd, ",")

	for _, indicator := range indicatorArr {
		menu, _ := strconv.Atoi(indicator)
		switch {
		case menu == 1:
			fmt.Println("开始执行 ----> 基金资料")
			fundcodeArray = fundcodeArray[470:len(fundcodeArray)]
			mongodb.StoreFundBasicModel(fundcodeArray)
		case menu == 2:
			fmt.Println("开始执行 ----> 基金持仓")
			mongodb.StoreFundHoldingStockModel(fundcodeArray)
		}
	}

}
