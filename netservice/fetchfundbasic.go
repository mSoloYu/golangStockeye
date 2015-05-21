package netservice

import (
	//"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
func MakeFundBasicArray(fundcode string) (marginTradingInfos []string) {

	marginTradingInfos = doMakeFundBasicInfoArray(fundcode, doMakeFundBasicArray)

	return

}

func doMakeFundBasicInfoArray(fundcode string, opFunc holderFunc) []string {

	recArr, isPanic := opFunc(fundcode)
	for isPanic {
		recArr, isPanic = opFunc(fundcode)
	}

	return recArr

}

//
func doMakeFundBasicArray(fundcode string) (infoArray []string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			printLog(e)
			isPanic = true
			return
		}
	}()

	isPanic = false
	infoArray = make([]string, 64)

	doc := getGoQueryDocument(fundbasicUrlModel, fundcode, "搜狐-基金概况")

	idx := 0
	re := regexp.MustCompile("\\s+")
	doc.Find("table").Each(func(i int, table *goquery.Selection) {

		if i == 1 {
			table.Find("td").Each(func(j int, td *goquery.Selection) {

				str := getUtfTextFromGoQuerySelection(td)
				text := re.ReplaceAllString(str, "")
				//log.Println(text)
				if !strings.EqualFold(text, "其中") {
					infoArray[idx] = strings.Replace(strings.Replace(str, "-", "", -1), ".", "", -1)
					idx++
				}

			})
		}

	})

	infoArray = infoArray[0:idx]
	//log.Println("--->", infoArray)

	return

}
