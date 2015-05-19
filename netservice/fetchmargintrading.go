package netservice

import (
	//"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
func MakeStockMarginTradingArray(stockcode string) (marginTradingInfos []string) {

	marginTradingInfos = doMakeStockMarginTradingInfosArray(stockcode, doMakeStockMarginTradingArray)

	return

}

func doMakeStockMarginTradingInfosArray(stockcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(stockcode)
	for isPanic {
		recArr, isPanic = opFunc(stockcode)
	}

	return recArr

}

//
func doMakeStockMarginTradingArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 10240)

	pageCount, arrIdx := makeStockMarginTradingInfoPerPage(0, 1, stockcode, stockInfoArray)

	jdx := 2
	for true {
		if jdx > pageCount {
			break
		}

		_, idx := makeStockMarginTradingInfoPerPage(arrIdx, jdx, stockcode, stockInfoArray)
		arrIdx = idx

		jdx++
	}

	//log.Println("--->", arrIdx)
	stockInfoArray = stockInfoArray[0:arrIdx]
	//log.Println("--->", stockInfoArray)

	return

}

func makeStockMarginTradingInfoPerPage(arrIdx, pageIdx int, stockcode string,
	stockInfoArray []string) (pageCount, idx int) {

	urlModel := strings.Replace(marginTradingUrlModel, "*", strconv.Itoa(pageIdx), -1)
	doc := getGoQueryDocument(urlModel, stockcode, "搜狐-融资融券")

	idx = arrIdx
	doc.Find("#BIZ_MS_JDPH").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {

			str := getUtfTextFromGoQuerySelection(td)
			if j == 0 || (j > 10 && j%11 == 0) {
				stockInfoArray[idx] = strings.Replace(str, "-", "", -1)
				idx++
			} else if (j < 12 && (j != 1 && j != 2 && j != 9 && j != 10)) ||
				(j > 11 && (j-1)%11 != 0 && (j-2)%11 != 0 && (j-9)%11 != 0 && (j-10)%11 != 0) {
				//log.Println(str)
				stockInfoArray[idx] = strings.Replace(str, ",", "", -1)
				idx++
			}

		})

	})

	if arrIdx == 0 {
		doc.Find("#pageDiv").Each(func(i int, div *goquery.Selection) {

			str := getUtfTextFromGoQuerySelection(div)
			pageAfterText := strings.Split(str, "/")[1]
			pageText := strings.Split(pageAfterText, "\n")[0]
			//log.Println("--->pages : ", pageText)
			count, _ := strconv.Atoi(pageText)
			pageCount = count

		})
	}

	return

}
