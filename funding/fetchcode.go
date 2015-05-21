package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

func main() {

	res, _ := http.Get("http://fund.jrj.com.cn/funddata/family1.shtml#1")

	doc, _ := goquery.NewDocumentFromResponse(res)

	doc.Find(".tab3").Each(func(i int, table *goquery.Selection) {

		table.Find("td").Each(func(j int, td *goquery.Selection) {

			str := td.Text()
			if len(str) != 0 {
				text, _ := iconv.ConvertString(str, "gbk", "utf-8")
				fmt.Println(text)
			}

		})

	})

}
