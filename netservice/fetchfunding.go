package netservice

import (
	//"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

//
func MakeStockFundingArray(stockcode string) (fundingInfos []string) {

	fundingInfos = doMakeStockFundingInfosArray(stockcode, doMakeStockFundingArray)

	//log.Println("--->", fundingInfos)

	return

}

func doMakeStockFundingInfosArray(stockcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(stockcode)
	for isPanic {
		recArr, isPanic = opFunc(stockcode)
	}

	return recArr

}

//
func doMakeStockFundingArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			if isPanic == false {
				stockInfoArray = []string{"0"}
				return
			}
			return
		}
	}()

	isPanic = true
	stockInfoArray = make([]string, 4096)

	doc := getGoQueryDocument(fundingsUrlModel, stockcode, "搜狐-基金持仓")

	idx, jdx := 0, 0
	doc.Find(".innerStr2ColumnBL").Each(func(i int, div *goquery.Selection) {

		div.Find("h4").Each(func(h int, header *goquery.Selection) {
			val := header.Text()
			output, _ := iconv.ConvertString(val, "gbk", "utf-8")
			str := strings.Replace(strings.TrimSpace(output), "-", "", -1)
			stockInfoArray[idx] = strings.Split(str, "：")[1]
			idx++
		})

		div.Find("table").Each(func(j int, table *goquery.Selection) {
			table.Find("td").Each(func(k int, td *goquery.Selection) {

				str := getUtfTextFromGoQuerySelection(td)
				if len(str) == 0 {
					str = "0"
				}
				if k == 7 {
					str = strings.Replace(strings.Replace(str, ".", "", -1), "%", "", -1)
				}
				stockInfoArray[idx] = str
				idx++

			})
		})
	})

	doc.Find(".innerStr2ColumnBR").Each(func(i int, div *goquery.Selection) {
		div.Find("table").Each(func(j int, table *goquery.Selection) {
			table.Find("th").Each(func(k int, th *goquery.Selection) {
				str := getUtfTextFromGoQuerySelection(th)
				stockInfoArray[idx] = str
				//log.Println("--->", str, "--->", jdx)
				idx++
				jdx++
			})
		})
		if jdx == 0 {
			isPanic = false
			panic("zero")
		}
		for jdx <= 9 {
			//log.Println("--->", stockInfoArray[idx-1], "--->", jdx)
			stockInfoArray[idx] = stockInfoArray[idx-1]
			idx++
			jdx++
		}
		jdx = 0
		div.Find("table").Each(func(j int, table *goquery.Selection) {
			table.Find("td").Each(func(k int, td *goquery.Selection) {
				str := getUtfTextFromGoQuerySelection(td)
				stockInfoArray[idx] = str
				//log.Println("--->", str, "--->", jdx)
				idx++
				jdx++
			})
		})
		for jdx <= 18 {
			stockInfoArray[idx] = stockInfoArray[idx-2]
			//log.Println("--->", stockInfoArray[idx], "--->")
			idx++
			stockInfoArray[idx] = stockInfoArray[idx-2]
			//log.Println("--->", stockInfoArray[idx], "--->")
			idx++
			jdx += 2
		}
	})

	doc.Find(".tableJ").Each(func(i int, table *goquery.Selection) {
		table.Find("td").Each(func(j int, td *goquery.Selection) {

			str := getUtfTextFromGoQuerySelection(td)
			if (j-5)%6 == 0 {
				str = strings.Replace(strings.Replace(str, ".", "", -1), "%", "", -1)
			}
			stockInfoArray[idx] = str
			idx++

		})
	})

	stockInfoArray = stockInfoArray[0:idx]
	isPanic = false

	return

}
