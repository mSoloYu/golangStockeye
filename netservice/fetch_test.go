package netservice

import (
	"log"
	"testing"
)

var testCasesMakeXlsUrl = []struct {
	inDate   string
	inSymbol string
	outUrl   string
}{
	{"2015-04-30", "sh601668", "http://market.finance.sina.com.cn/downxls.php?date=2015-04-30&symbol=sh601668"},
	{"2015-05-04", "sh601668", "http://market.finance.sina.com.cn/downxls.php?date=2015-04-30&symbol=sh601668"},
}

func TestMakeXlsRecords(t *testing.T) {
	for idx, testCase := range testCasesMakeXlsUrl {
		resultUrl := MakeXlsRecords(testCase.inDate, testCase.inSymbol)
		verify(t, idx, "MakeXlsUrl: ", testCase.inDate, testCase.inSymbol, resultUrl, testCase.outUrl)
	}
}

func TestMakeStockBasic(t *testing.T) {

	stockInfoArray := MakeStockBasicArray("601668")
	for _, stockInfoItem := range stockInfoArray {
		log.Println(stockInfoItem)
	}

}

func TestMakeStockCurrentQuote(t *testing.T) {

	quoteItems := MakeStockCurrentQuote("601668")
	log.Println(quoteItems)
	t.Fail()

}

func TestMakeStockLastdayClosePrice(t *testing.T) {
	log.Println("前一日报价：", MakeStockLastdayClosePrice("601668", "2015-05-06"))
	t.Fail()
}

func verify(t *testing.T, testnum int, testcase, inDate, inSymbol, output, expected string) {
	if expected != output {
		t.Errorf("\n%d. %s with input = %s, %s: output \n%s != \n%s",
			testnum, testcase, inDate, inSymbol, output, expected)
	}
}
