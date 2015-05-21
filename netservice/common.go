package netservice

import (
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

type holderFunc func(a string) ([]string, bool)
type fundtranFunc func(a string) ([]string, []string, bool)

const stockBasicSize = 78
const strSymbol = "&symbol="

const quoteUrlPrefix = "http://hq.sinajs.cn/list="
const xlsUrlPrefix = "http://market.finance.sina.com.cn/downxls.php?date="
const stockInfoUrlPrefix = "http://stockdata.stock.hexun.com/gszl/s"
const historyQuoteUrlModel = "http://vip.stock.finance.sina.com.cn/quotes_service" +
	"/view/vMS_tradehistory.php?symbol=--------&date=**********"
const mainHoldersUrlModel = "http://vip.stock.finance.sina.com.cn/corp/go.php/vCI_StockHolder/stockid/------.phtml"
const publicHoldersUrlModel = "http://vip.stock.finance.sina.com.cn/corp/go.php/vCI_CirculateStockHolder/stockid/------.phtml"

const rankingUrlModel = "http://stockdata.stock.hexun.com/------.shtml"
const accountingUrlModel = "http://q.stock.sohu.com/cn/------/index.shtml"
const stockFundingUrlModel = "http://q.stock.sohu.com/cn/------/jjcc.shtml"

const bigDealUrlModel = "http://q.stock.sohu.com/cn/------/dzjy.shtml"
const marginTradingUrlModel = "http://q.stock.sohu.com/app2/mpssTrade.up?code=------&sd=&ed=&p=*"

//const internalTradingUrlModel = "http://q.stock.sohu.com/app2/rpsholder.up?code=------&sd=&ed=&type=date&dir=1"
const moneyFlowUrlModel = "http://app.finance.ifeng.com/hq/trade/stock_zijin.php?code=--------&begin_day=**********&end_day=**********"

// 十大重仓
const fundUrlModelBySohu = "http://q.fund.sohu.com/------/index.shtml?code=------"

// 投资行业分布明细
const fundUrlModelBySina = "http://finance.sina.com.cn/fund/quotes/------/bc.shtml"

const fundbasicUrlModel = "http://q.fund.sohu.com/q/gk.php?code=------"

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

func getUsedStockcode(stockcode string) (usedStockcode string) {

	if strings.HasPrefix(stockcode, "00") {
		usedStockcode = "sz" + stockcode
	} else {
		usedStockcode = "sh" + stockcode
	}

	return

}

func getUtfTextFromGoQuerySelection(selection *goquery.Selection) (text string) {

	val := selection.Text()
	output, _ := iconv.ConvertString(val, "gbk", "utf-8")
	text = strings.TrimSpace(output)

	return

}

func getGoQueryDocument(urlModel, stockcode, strPanic string) *goquery.Document {

	urlStr := strings.Replace(urlModel, "------", stockcode, -1)
	res, err := getTimeoutHttpClient().Get(urlStr)
	if err != nil {
		panic(strPanic)
	}

	doc, _ := goquery.NewDocumentFromResponse(res)

	return doc

}

func printLog(e interface{}) {
	log.Println("出错：", e)
}
