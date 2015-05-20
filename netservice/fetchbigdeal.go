package netservice

import (
	//"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
func MakeStockBigDealArray(stockcode string) (bigDealInfos []string) {

	bigDealInfos = doMakeStockBigDealInfosArray(stockcode, doMakeStockBigDealArray)

	return

}

func doMakeStockBigDealInfosArray(stockcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(stockcode)
	for isPanic {
		recArr, isPanic = opFunc(stockcode)
	}

	return recArr

}

//
func doMakeStockBigDealArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 1024)

	doc := getGoQueryDocument(bigDealUrlModel, stockcode, "搜狐-大宗交易")

	idx := 0
	doc.Find("#BIZ_hq_historySearch").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {

			str := getUtfTextFromGoQuerySelection(td)
			if j == 0 || (j > 5 && j%6 == 0) {
				stockInfoArray[idx] = strings.Replace(str, "-", "", -1)
			} else {
				stockInfoArray[idx] = str
			}
			//log.Println(stockInfoArray[idx])
			idx++

		})

	})

	stockInfoArray = stockInfoArray[0:idx]
	//log.Println("--->", stockInfoArray)

	return

}
