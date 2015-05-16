package netservice

import (
	"net"
	"net/http"
	"strings"
	"time"
)

const stockBasicSize = 78
const strSymbol = "&symbol="

const quoteUrlPrefix = "http://hq.sinajs.cn/list="
const xlsUrlPrefix = "http://market.finance.sina.com.cn/downxls.php?date="
const stockInfoUrlPrefix = "http://stockdata.stock.hexun.com/gszl/s"
const historyQuoteUrlModel = "http://vip.stock.finance.sina.com.cn/quotes_service" +
	"/view/vMS_tradehistory.php?symbol=--------&date=**********"
const rankingUrlModel = "http://stockdata.stock.hexun.com/------.shtml"
const mainHoldersUrlModel = "http://vip.stock.finance.sina.com.cn/corp/go.php/vCI_StockHolder/stockid/------.phtml"
const publicHoldersUrlModel = "http://vip.stock.finance.sina.com.cn/corp/go.php/vCI_CirculateStockHolder/stockid/------.phtml"
const bigDealUrlModel = "http://q.stock.sohu.com/cn/------/dzjy.shtml"
const marginTradingUrlModel = "http://q.stock.sohu.com/app2/mpssTrade.up?code=------&sd=&ed="

//const internalTradingUrlModel = "http://q.stock.sohu.com/app2/rpsholder.up?code=------&sd=&ed=&type=date&dir=1"

const accountingUrlModel = "http://q.stock.sohu.com/cn/------/index.shtml"
const fundPositionsUrlModel = "http://q.stock.sohu.com/cn/------/jjcc.shtml"

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
