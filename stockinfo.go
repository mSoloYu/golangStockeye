package main

import (
	"./mongodb"
)

const stockcodeFilename = "stock_a.txt"

var stockcodeArray []string

func init() {
	stockcodeArray = parseFileToStringArray(stockcodeFilename)
	//stockcodeArray = stockcodeArray[2020:len(stockcodeArray)]
}

func main() {

	mongodb.StoreStockModel(stockcodeArray)

}
