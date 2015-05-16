package main

import (
	"log"
	"runtime"

	"./mongodb"
	"./utils"
)

var stockcodeArray []string

func init() {

	runtime.GOMAXPROCS(5)

	stockcodeArray = parseFileToStringArray()
	//stockcodeArray = stockcodeArray[2020:len(stockcodeArray)]

}

func main() {

	_, settedDateArray := utils.ParseCmdFlagToDateFlag()

	for _, stockcode := range stockcodeArray {
		log.Println(stockcode)
		for _, date := range settedDateArray {
			mongodb.DeleteTranRecordWithDate(stockcode, date)
			log.Println("------ ", date)
		}
	}

}
