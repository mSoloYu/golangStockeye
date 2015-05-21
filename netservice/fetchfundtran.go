package netservice

import (
	//"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
func MakeFundTranArray(fundcode string) (tenFundHoldingStocks, fiveBondInfos, tradeDistributions, perfromanceInfos []string) {

	tenFundHoldingStocks, fiveBondInfos = doMakeFundTranInfoArray(fundcode, doMakeTenFundHoldingStocksArray)
	tradeDistributions, perfromanceInfos = doMakeFundTranInfoArray(fundcode, doMakeTradeDistributionsArray)

	return

}

func doMakeFundTranInfoArray(fundcode string, opFunc fundtranFunc) (infoArr1, infoArr2 []string) {

	var isPanic bool
	infoArr1, infoArr2, isPanic = opFunc(fundcode)
	for isPanic {
		infoArr1, infoArr2, isPanic = opFunc(fundcode)
	}

	return

}

//
func doMakeTenFundHoldingStocksArray(fundcode string) (infoArray, bondinfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	infoArray = make([]string, 51)
	bondinfoArray = make([]string, 50)

	doc := getGoQueryDocument(fundUrlModelBySohu, fundcode, "十大重仓股, 五大债券持仓")

	idx, jdx := 0, 0

	doc.Find(".date_more").Each(func(i int, div *goquery.Selection) {
		if i == 0 {
			div.Find("span").Each(func(j int, span *goquery.Selection) {
				str := getUtfTextFromGoQuerySelection(span)
				infoArray[idx] = strings.Replace(strings.Split(str, "：")[1], "-", "", -1)
				idx++
			})
		}
	})

	doc.Find(".rtable02").Each(func(i int, table *goquery.Selection) {

		if i == 0 {
			table.Find("tbody").Each(func(j int, tbody *goquery.Selection) {

				tbody.Find("td").Each(func(k int, td *goquery.Selection) {

					str := getUtfTextFromGoQuerySelection(td)
					if (k < 5 && k != 0 && k != 5) || (k > 5 && k%6 != 0 && (k+1)%6 != 0) {
						infoArray[idx] = strings.Replace(strings.Replace(str, "--", "0.00", -1), "%", "", -1)
						idx++
					}

				})

			})
		}

		if i == 1 {
			table.Find("tbody").Each(func(j int, tbody *goquery.Selection) {

				tbody.Find("td").Each(func(k int, td *goquery.Selection) {

					str := getUtfTextFromGoQuerySelection(td)
					if (k < 4 && k != 0) || (k > 4 && k%5 != 0) {
						bondinfoArray[jdx] = strings.Replace(strings.Replace(str, "--", "0.00", -1), "%", "", -1)
						jdx++
					}

				})

			})
		}

	})

	infoArray = infoArray[0:idx]
	bondinfoArray = bondinfoArray[0:jdx]
	//log.Println("--->", infoArray)
	//log.Println("--->", bondinfoArray)

	return

}

func doMakeTradeDistributionsArray(fundcode string) (infoArray, performanceinfoArr []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	infoArray = make([]string, 128)
	performanceinfoArr = make([]string, 128)

	doc := getGoQueryDocument(fundUrlModelBySina, fundcode, "投资行业分布")

	idx, jdx := 0, 0
	doc.Find("#table-distribution-list").Each(func(i int, table *goquery.Selection) {

		table.Find("span").Each(func(j int, span *goquery.Selection) {

			str := getUtfTextFromGoQuerySelection(span)
			if j == 0 || (j > 1 && j%2 == 0) {
				infoArray[idx] = str
			} else {
				infoArray[idx] = strings.Replace(str, "%", "", -1)
			}
			idx++

		})

	})

	doc.Find("#table-fund-performance-rank").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody").Each(func(j int, tbody *goquery.Selection) {
			table.Find("td").Each(func(k int, td *goquery.Selection) {

				if k < 28 {
					str := getUtfTextFromGoQuerySelection(td)
					str = strings.Replace(strings.Replace(str, "-", "", -1), "%", "", -1)
					performanceinfoArr[jdx] = str
					jdx++
				}

			})
		})

	})

	infoArray = infoArray[0:idx]
	performanceinfoArr = performanceinfoArr[0:jdx]
	//log.Println("--->", infoArray)
	//log.Println("--->", performanceinfoArr)

	return

}
