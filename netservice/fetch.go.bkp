package netservice

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../utils"
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

const stockInfoSize = 78
const strSymbol = "&symbol="

const quoteUrlPrefix = "http://hq.sinajs.cn/list="
const xlsUrlPrefix = "http://market.finance.sina.com.cn/downxls.php?date="
const stockInfoUrlPrefix = "http://stockdata.stock.hexun.com/gszl/s"
const historyQuoteUrlModel = "http://vip.stock.finance.sina.com.cn/quotes_service" +
	"/view/vMS_tradehistory.php?symbol=--------&date=**********"

func getTimeoutHttpClient() *http.Client {

	c := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	return &c

}

// MakeXlsUrl combine the date and the symbol which is stock code to an xls file download url.
func MakeXlsRecords(symbol, date string) string {

	rec, isPanic := doMakeXlsRecords(symbol, date)
	for isPanic {
		rec, isPanic = doMakeXlsRecords(symbol, date)
	}
	return rec

}

func doMakeXlsRecords(symbol, date string) (record string, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			log.Println("交易情况网络出错：", e)
			record = ""
			isPanic = true
			return
		}
	}()

	isPanic = false

	xlsUrl := xlsUrlPrefix + date + strSymbol + getUsedStockcode(symbol)

	res, err := getTimeoutHttpClient().Get(xlsUrl)
	if err != nil {
		panic("xls")
	}

	rawRecords, _ := ioutil.ReadAll(res.Body)

	strRecords := string(rawRecords)

	if strings.Contains(strRecords, "javascript") &&
		strings.Contains(strRecords, "window.close") {
		record = ""
		return
	}

	if !isStockTranRecordsComplete(strRecords) {
		record = ""
		return
	}

	record = string(rawRecords)
	return

}

// MakeStockInfoArray fetch and parse the webpage for stock information like name, code, issue price, and pe, etc.
func MakeStockInfoArray(stockcode string) []string {

	stockInfoArray := make([]string, stockInfoSize)

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
			stockInfoArray[idx] = output
			idx++
		})

	})

	return stockInfoArray

}

//
//
// var hq_str_sz002024="苏宁云商,11.26,11.22,11.33,11.45,11.17,11.33,11.34,131241390,1484755809.02,
// 153559,11.33,358350,11.32,396400,11.31,654870,11.30,151700,11.29,420400,11.34,315415,11.35,170927,11.36,202600,11.37,111001,11.38,
// 2015-03-13,15:05:53,00";
// name,open,lastClose,current,highest,lowest,close,**,vol(手),money(元),
// 0,   1,   2,        3,      4,      5,     6,    7, 8,      9,
func MakeStockCurrentQuote(stockcode string) []string {

	quoteUrl := quoteUrlPrefix + getUsedStockcode(stockcode)

	res, err := http.Get(quoteUrl)
	utils.CheckError(err)

	rawRecords, err := ioutil.ReadAll(res.Body)
	utils.CheckError(err)

	quoteItemArray := strings.Split(string(rawRecords), ",")

	return quoteItemArray[1:10]

}

//
func MakeStockLastdayClosePrice(stockcode, currentday string) float32 {

	if utils.IsToday(currentday) {
		return getTodayLastdayPrice(stockcode)
	} else {
		return getHistoryLastdayPrice(stockcode, currentday)
	}

}

func getTodayLastdayPrice(stockcode string) float32 {

	infos := MakeStockCurrentQuote(stockcode)
	thePrice, _ := strconv.ParseFloat(infos[1], 32)
	return float32(thePrice)

}

func getHistoryLastdayPrice(stockcode, currentday string) float32 {

	price, isPanic := doMakeStockLastdayClosePrice(stockcode, currentday)
	for isPanic {
		price, isPanic = doMakeStockLastdayClosePrice(stockcode, currentday)
	}
	return price

}

func doMakeStockLastdayClosePrice(stockcode, currentday string) (lastdayPrice float32, isPanic bool) {

	defer func() {
		if e := recover(); e != nil {
			log.Println("昨日报价网络出错：")
			lastdayPrice = 0.00
			isPanic = true
			return
		}
	}()

	isPanic = false

	lastdayClosePriceUrl := strings.Replace(
		strings.Replace(historyQuoteUrlModel, "--------", getUsedStockcode(stockcode), -1),
		"**********", currentday, -1)

	res, err := getTimeoutHttpClient().Get(lastdayClosePriceUrl)
	if err != nil {
		panic("lastday price")
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	utils.CheckError(err)

	mainTag := doc.Find(".main").First()
	if mainTag == nil || strings.Contains(mainTag.Text(), "输入的日期为非交易日期") {
		lastdayPrice = 0.00
		return
	}

	quoteareaTag := doc.Find("#quote_area").First()
	if quoteareaTag == nil {
		lastdayPrice = 0.00
		return
	}

	quoteareaTag.Find("td").Each(func(i int, tag *goquery.Selection) {
		if i == 5 {
			val, _ := strconv.ParseFloat(tag.Text(), 32)
			lastdayPrice = float32(val)
		}
	})

	return

}

//
func getUsedStockcode(stockcode string) (usedStockcode string) {

	if strings.HasPrefix(stockcode, "00") {
		usedStockcode = "sz" + stockcode
	} else {
		usedStockcode = "sh" + stockcode
	}

	return

}

//
func isStockTranRecordsComplete(strRecords string) bool {

	isComplete := (strings.Contains(strRecords, "9:30") ||
		strings.Contains(strRecords, "9:37") || strings.Contains(strRecords, "9:41")) &&
		(strings.Contains(strRecords, "9:48") ||
			strings.Contains(strRecords, "9:53") || strings.Contains(strRecords, "9:59")) &&
		(strings.Contains(strRecords, "10:04") ||
			strings.Contains(strRecords, "10:07") || strings.Contains(strRecords, "10:12")) &&
		(strings.Contains(strRecords, "10:17") ||
			strings.Contains(strRecords, "10:22") || strings.Contains(strRecords, "10:29")) &&
		(strings.Contains(strRecords, "10:34") ||
			strings.Contains(strRecords, "10:40") || strings.Contains(strRecords, "10:43")) &&
		(strings.Contains(strRecords, "10:46") ||
			strings.Contains(strRecords, "10:52") || strings.Contains(strRecords, "10:58")) &&
		(strings.Contains(strRecords, "11:03") ||
			strings.Contains(strRecords, "11:06") || strings.Contains(strRecords, "11:13")) &&
		(strings.Contains(strRecords, "11:18") ||
			strings.Contains(strRecords, "11:23") || strings.Contains(strRecords, "11:28")) &&
		(strings.Contains(strRecords, "13:03") ||
			strings.Contains(strRecords, "13:07") || strings.Contains(strRecords, "13:13")) &&
		(strings.Contains(strRecords, "13:18") ||
			strings.Contains(strRecords, "13:23") || strings.Contains(strRecords, "13:29")) &&
		(strings.Contains(strRecords, "13:34") ||
			strings.Contains(strRecords, "13:37") || strings.Contains(strRecords, "13:42")) &&
		(strings.Contains(strRecords, "13:47") ||
			strings.Contains(strRecords, "13:52") || strings.Contains(strRecords, "13:59")) &&
		(strings.Contains(strRecords, "14:04") ||
			strings.Contains(strRecords, "14:10") || strings.Contains(strRecords, "14:13")) &&
		(strings.Contains(strRecords, "14:16") ||
			strings.Contains(strRecords, "14:22") || strings.Contains(strRecords, "14:28")) &&
		(strings.Contains(strRecords, "14:33") ||
			strings.Contains(strRecords, "14:36") || strings.Contains(strRecords, "14:43")) &&
		(strings.Contains(strRecords, "14:48") ||
			strings.Contains(strRecords, "14:53") || strings.Contains(strRecords, "14:48"))

	return isComplete

}
