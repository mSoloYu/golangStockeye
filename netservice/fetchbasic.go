package netservice

import (
	"net/http"
	"strings"

	"../utils"
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

// MakeStockBasicArray fetch and parse the webpage for stock information like name, code, issue price, and pe, etc.
func MakeStockBasicArray(stockcode string) []string {

	stockBasicArray := make([]string, stockBasicSize)

	res, err := http.Get(stockInfoUrlPrefix + stockcode + ".shtml")
	utils.CheckError(err)

	doc, err := goquery.NewDocumentFromResponse(res)
	utils.CheckError(err)

	idx := 0
	doc.Find(".tab_xtable").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {
			val := td.Text()
			output, _ := iconv.ConvertString(val, "gbk", "utf-8")
			output = strings.TrimSpace(output)
			stockBasicArray[idx] = output
			idx++
		})

	})

	return stockBasicArray

}
