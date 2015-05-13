package utils

import (
	"log"
	"testing"
)

var testCasesMakeXlsUrl = []struct {
	inDate string
	inSymbol string
	outUrl string
}{
	{"2015-04-30", "sh601668", "http://market.finance.sina.com.cn/downxls.php?date=2015-04-30&symbol=sh601668"},
	{"2015-05-04", "sh601668", "http://market.finance.sina.com.cn/downxls.php?date=2015-04-30&symbol=sh601668"},
}

func TestMakeXlsUrl(t *testing.T) {
	for idx, testCase := range testCasesMakeXlsUrl {
		resultUrl := MakeXlsUrl(testCase.inDate, testCase.inSymbol)
		verify(t, idx, "MakeXlsUrl: ", testCase.inDate, testCase.inSymbol, resultUrl, testCase.outUrl)
	}
}

func TestMakeStockInfo(t * testing.T) {

	stockInfoArray := MakeStockInfoArray("601668")
	for _, stockInfoItem := range stockInfoArray {
		log.Println(stockInfoItem)
	}

}

func verify(t *testing.T, testnum int, testcase, inDate, inSymbol, output, expected string)  {
	if expected != output {
        t.Errorf("\n%d. %s with input = %s, %s: output \n%s != \n%s", 
              testnum, testcase, inDate, inSymbol, output, expected)
    }
}
