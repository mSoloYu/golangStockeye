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
	fmt.Println("简要财务 ----> 3")
	fmt.Println("基金持仓 ----> 4")
	fmt.Println("融资融券 ----> 5")
	fmt.Println("大宗交易 ----> 6")
	fmt.Println("")
	fmt.Println("退出程序 ----> 回车")
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
		case menu == 3:
			fmt.Println("开始执行 ----> 简要财务")
			//stockcodeArray = stockcodeArray[460:len(stockcodeArray)]
			mongodb.StoreStockAccountingModel(stockcodeArray)
		case menu == 4:
			fmt.Println("开始执行 ----> 基金持仓")
			//stockcodeArray = stockcodeArray[2053:len(stockcodeArray)]
			mongodb.StoreStockFundingModel(stockcodeArray)
		case menu == 5:
			fmt.Println("开始执行 ----> 融资融券")
			//stockcodeArray = stockcodeArray[537:len(stockcodeArray)]
			mongodb.StoreStockMarginTradingModel(stockcodeArray)
		case menu == 6:
			fmt.Println("开始执行 ----> 大宗交易")
			//stockcodeArray = stockcodeArray[537:len(stockcodeArray)]
			mongodb.StoreStockBigDealModel(stockcodeArray)
		case menu == 0:
			fallthrough
		default:
			break
		}
	}

}
