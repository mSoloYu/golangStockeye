package main

import (
	"log"
	"runtime"
	"sync"

	"./business"
	"./mongodb"
	"./netservice"
	"./utils"
)

//const stockcodeFilename = "stock_a.txt"
const stockcodeFilename = "stock_choose.txt"

var stockcodeArray []string

func init() {

	runtime.GOMAXPROCS(5)

	stockcodeArray = parseFileToStringArray(stockcodeFilename)
	//stockcodeArray = stockcodeArray[2020:len(stockcodeArray)]
}

func main() {

	var stockcodeWg sync.WaitGroup
	var stockcodeArrayWg sync.WaitGroup
	stocktran := new(mongodb.StockTran)

	counterChan := make(chan int)
	lastdayPriceChan := make(chan float32)
	dateChan := make(chan string)
	recordChan := make(chan string)

	unCompletedCounter := -1
	for idx, stockcode := range stockcodeArray {

		stockcodeArrayWg.Add(1)

		mongodb.ConnectToStockTranCollection(stockcode)

		// 获取首日上市日期
		// 获取交易日期数组
		dateArray := utils.GetDateArray(
			mongodb.FindStockinfoOpenSaleDate(stockcode))
		unCompletedCounter = len(dateArray)

		log.Printf("%4d - %s", idx, stockcode)
		// 获取每日交易数据
		go func() {
			for counter, date := range dateArray {
				counterChan <- counter
				dateChan <- date
				go func() {
					lastdayPriceChan <- netservice.MakeStockLastdayClosePrice(stockcode, date)
				}()
				go func() {
					recordChan <- netservice.MakeXlsRecords(stockcode, date)
				}()
			}
		}()

		go func() {
			defer stockcodeArrayWg.Done()

			var date string
			for unCompletedCounter != 0 {
				counter := <-counterChan
				stockcodeWg.Add(1)
				go func() {
					defer stockcodeWg.Done()

					date = <-dateChan
					lastdayPrice := <-lastdayPriceChan
					record := <-recordChan
					if len(record) != 0 {
						stocktran = business.ParseDaily(lastdayPrice, date, record)
					} else {
						stocktran = nil
					}
				}()
				stockcodeWg.Wait()
				if stocktran != nil {
					mongodb.StoreStockTranDailyModel(stocktran)
					log.Printf("     - %4d, %d\n", counter, date)
				} else {
					log.Printf("     - %4d, %d -- skip\n", counter, date)
				}
				unCompletedCounter--
			}
		}()

		stockcodeArrayWg.Wait()

	}

}
