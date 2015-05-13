package utils

import (
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"log"
	"net/http"
	"strings"
)

const stockInfoSize = 78
const strSymbol = "&symbol="

const xlsUrlPrefix = "http://market.finance.sina.com.cn/downxls.php?date="
const stockInfoUrlPrefix = "http://stockdata.stock.hexun.com/gszl/s"

// MakeXlsUrl combine the date and the symbol which is stock code to an xls file download url.
func MakeXlsRecords(date string, symbol string) string {

	xlsUrl := xlsUrlPrefix + date + strSymbol + symbol

	res, err := http.Get(xlsUrl)
	checkError(err)

	rawRecords, err := ioutil.ReadAll(res.Body)
	checkError(err)

	return rawRecords

}

// MakeStockInfoArray
func MakeStockInfoArray(stockcode string) []string {

	stockInfoArray := make([]string, stockInfoSize)

	res, err := http.Get(stockInfoUrlPrefix + stockcode + ".shtml")
	checkError(err)

	doc, err := goquery.NewDocumentFromResponse(res)
	checkError(err)

	idx := 0
	doc.Find(".tab_xtable").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {
			val := td.Text()
			output, _ := iconv.ConvertString(val, "gbk", "utf-8")
			output = strings.TrimSpace(output)
			stockInfoArray[idx] = output
			idx++
		})

	})

	return stockInfoArray

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}
