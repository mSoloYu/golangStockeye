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

var stockcodeArray []string

func init() {

	runtime.GOMAXPROCS(5)

	stockcodeArray = parseFileToStringArray()
	//stockcodeArray = stockcodeArray[2020:len(stockcodeArray)]
}

func main() {

	hasCmdFlag, settedDateArray := utils.ParseCmdFlag()

	var stockcodeWg sync.WaitGroup
	var stockcodeArrayWg sync.WaitGroup
	stocktran := new(mongodb.StockTran)

	counterChan := make(chan int, 1)
	dateChan := make(chan string)
	lastdayPriceChan := make(chan float32, 1)
	recordChan := make(chan string, 1)

	unCompletedCounter := -1
	for idx, stockcode := range stockcodeArray {

		stockcodeArrayWg.Add(1)

		mongodb.ConnectToStockTranCollection(stockcode)

		var dateArray []string
		if hasCmdFlag {
			dateArray = settedDateArray
		} else {
			// 获取首日上市日期
			// 获取交易日期数组
			dateArray = utils.GetDateArray(
				mongodb.FindStockinfoOpenSaleDate(stockcode))
		}
		unCompletedCounter = len(dateArray)

		log.Printf("%4d - %s", idx, stockcode)
		// 获取每日交易数据
		go func() {
			for counter, date := range dateArray {
				go func() {
					lastdayPriceChan <- netservice.MakeStockLastdayClosePrice(stockcode, date)
				}()
				go func() {
					recordChan <- netservice.MakeXlsRecords(stockcode, date)
				}()
				counterChan <- counter
				dateChan <- date
			}
		}()

		go func() {
			defer stockcodeArrayWg.Done()

			var date string
			for unCompletedCounter != 0 {
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

				counter := <-counterChan
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
