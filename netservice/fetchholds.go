package netservice

import (
	//"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

const HolderTypeMain = 1
const HolderTypePublic = 2

//
func MakeStockHoldersArray(stockcode string, holderType int) []string {

	if holderType == 1 {
		return doMakeStockHoldersArray(stockcode, doMakeStockMainHoldersArray)
	}

	return doMakeStockHoldersArray(stockcode, doMakeStockPublicHoldersArray)

}

func doMakeStockHoldersArray(stockcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(stockcode)
	for isPanic {
		recArr, isPanic = opFunc(stockcode)
	}

	return recArr

}

//
func doMakeStockMainHoldersArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 128)

	doc := getGoQueryDocument(mainHoldersUrlModel, stockcode, "主要股东")

	idx, jdx := 0, 0
	re := regexp.MustCompile("聽")
	doc.Find("#Table1").Each(func(i int, table *goquery.Selection) {

		table.Find("tbody").Each(func(j int, tbody *goquery.Selection) {

			tbody.Find("td").Each(func(k int, td *goquery.Selection) {
				val := td.Text()
				output, _ := iconv.ConvertString(val, "gbk", "utf-8")
				str := strings.TrimSpace(output)
				strItem := re.ReplaceAllString(str, "")
				if strings.HasPrefix(strItem, "截至日期") {
					jdx = 0
					idx++
					//log.Println(stockInfoArray[idx-1])
				} else if jdx == 1 || jdx == 7 || jdx == 9 ||
					(jdx > 14 && ((jdx-11)%5 == 0 || (jdx-12)%5 == 0 || (jdx-13)%5 == 0)) {
					if jdx == 7 {
						strItem = strings.Replace(strItem, "查看变化趋势", "", -1)
						if len(strItem) == 0 {
							strItem = "0"
						}
					}
					if jdx == 9 {
						if strings.Contains(strItem, "股") {
							strItem = strings.Split(strItem, "股")[0]
						} else {
							strItem = "0"
						}
					}
					if (jdx-12)%5 == 0 {
						strItem = re.ReplaceAllString(strItem, "")
					}
					if jdx == 1 {
						strItem = strings.Replace(strItem, "-", "", -1)
						stockInfoArray[idx] = strItem
					} else {
						strVal := stockInfoArray[idx]
						stockInfoArray[idx] = strVal + "|" + strItem
					}
				}
				jdx++
			})

		})

	})

	stockInfoArray = stockInfoArray[1:(idx + 1)]

	return

}

func doMakeStockPublicHoldersArray(stockcode string) (stockInfoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	stockInfoArray = make([]string, 128)

	doc := getGoQueryDocument(publicHoldersUrlModel, stockcode, "流通股股东")

	idx, jdx := 0, 0
	re := regexp.MustCompile("\\s+")
	doc.Find("#CirculateShareholderTable").Each(func(i int, table *goquery.Selection) {

		table.Find("tr").Each(func(j int, tr *goquery.Selection) {

			val := tr.Text()
			str := strings.TrimSpace(val)
			output := re.ReplaceAllString(str, "|")
			rec, _ := iconv.ConvertString(output, "gbk", "utf-8")

			if j > 0 {

				if strings.Contains(rec, "截止日期") {
					jdx = 0
					idx++
					stockInfoArray[idx] = strings.Replace(strings.Split(rec, "|")[1], "-", "", -1)
					//log.Println(stockInfoArray[idx-1])
				} else {
					if len(rec) > 0 &&
						!strings.HasPrefix(rec, "公告日期") &&
						!strings.HasPrefix(rec, "编号") {
						holdersArr := strings.Split(output, "|")
						item := stockInfoArray[idx]
						holderName, _ := iconv.ConvertString(holdersArr[1], "gbk", "utf-8")
						stockInfoArray[idx] = item + "|" +
							holderName + "|" + holdersArr[2] + "|" + holdersArr[3]
					}
				}
				jdx++

			}

		})

	})

	stockInfoArray = stockInfoArray[1:(idx + 1)]

	return

}
