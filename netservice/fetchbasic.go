package netservice

import (
	"net/http"

	"../utils"
	"github.com/PuerkitoBio/goquery"
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
			output := getUtfTextFromGoQuerySelection(td)
			stockBasicArray[idx] = output
			idx++
		})

	})

	return stockBasicArray

}
