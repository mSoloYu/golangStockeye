package mongodb

import (
	"strconv"
	"time"

	"../utils"
)

func init() {
}

func FindStockinfoOpenSaleDate(stockcode string) (opensaleDate time.Time) {

	queryStat := makeQueryModelForOpenSaleDate(stockcode)

	connectToStockDbReadonly(stockdb)

	var queryStockinfo Stock
	connectToStockCollection().Find(queryStat).One(&queryStockinfo)

	strDate := strconv.Itoa(queryStockinfo.OpenSaleDate)
	return utils.ParseDateWithIntFmt(strDate)

}
