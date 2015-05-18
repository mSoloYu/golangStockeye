package netservice

import (
	//"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

//
func MakeStockAccountingArray(stockcode string) (rankingInfos, accountingInfos []string) {

	rankingInfos = doMakeStockInfosArray(stockcode, doMakeStockRankingArray)
	accountingInfos = doMakeStockInfosArray(stockcode, doMakeStockAccountingArray)

	return

}

func doMakeStockInfosArray(stockcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(stockcode)
	for isPanic {
		recArr, isPanic = opFunc(stockcode)
	}

	return recArr

}

//
func doMakeStockRankingArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 6)

	doc := getGoQueryDocument(rankingUrlModel, stockcode, "和讯-排名")

	idx := 0
	doc.Find(".box6").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {

			val := td.Text()
			output, _ := iconv.ConvertString(val, "gbk", "utf-8")
			str := strings.TrimSpace(output)
			switch {
			case j == 24:
				fallthrough
			case j == 27:
				stockInfoArray[idx] = str
				idx++
			case j == 25:
				fallthrough
			case j == 28:
				vals := strings.Split(str, "名")
				stockInfoArray[idx] = strings.Split(vals[1], "位")[0]
				idx++
				stockInfoArray[idx] = strings.Split(vals[2], "位")[0]
				idx++
			}
		})

	})

	//log.Println("--->", stockInfoArray)

	return

}

func doMakeStockAccountingArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 128)

	doc := getGoQueryDocument(accountingUrlModel, stockcode, "搜狐-简要财报")

	idx, jdx := 0, 0
	doc.Find(".tablebox_l").Each(func(i int, table *goquery.Selection) {

		if i < 1 {
			table.Find("td").Each(func(j int, td *goquery.Selection) {

				val := td.Text()
				output, _ := iconv.ConvertString(val, "gbk", "utf-8")
				str := strings.TrimSpace(output)
				if j == 0 {
					stockInfoArray[idx] = strings.Replace(str, "-", "", -1)
					idx++
				}
				if j > 3 && j < 24 {
					if jdx > 0 {
						if j == 13 || j == 17 {
							str = strings.Replace(str, "亿", "", -1)
						} else if strings.EqualFold(str, "-") {
							str = "0.00"
						}
						stockInfoArray[idx] = str
						idx++
					}
					jdx++
					if jdx > 3 {
						jdx = 0
					}
				} else if j == 29 || j == 31 || j == 33 {
					stockInfoArray[idx] = str
					idx++
				}
			})
		}

	})

	stockInfoArray = stockInfoArray[0:idx]
	//log.Println("--->", stockInfoArray)

	return

}
