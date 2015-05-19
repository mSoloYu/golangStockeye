package mongodb

import (
	//"log"

	"../netservice"
	"../utils"
)

//func makeStockMarginTradingDoc(stockcode string) {
func makeStockMarginTradingDoc(stockcode string) StockMarginTrading {

	marginTradingInfos := netservice.MakeStockMarginTradingArray(stockcode)

	var mMarginT StockMarginTrading
	mMarginT.Code = stockcode
	mMarginT.FetchDate = utils.GetTodayWithIntFmt()

	size := len(marginTradingInfos)
	recSize := size / 7
	mMarginTradingArr := make([][7]int64, recSize)
	idx, jdx := 0, 0
	for true {
		if idx >= recSize {
			break
		}

		dateIntFmt := utils.GetInt64Val(marginTradingInfos[jdx+0])
		buyMoney := utils.GetInt64Val(marginTradingInfos[jdx+1])
		payMoney := utils.GetInt64Val(marginTradingInfos[jdx+2])
		balanceMoney := utils.GetInt64Val(marginTradingInfos[jdx+3])
		saleStockVol := utils.GetInt64Val(marginTradingInfos[jdx+4])
		payStockVol := utils.GetInt64Val(marginTradingInfos[jdx+5])
		balanceStockVol := utils.GetInt64Val(marginTradingInfos[jdx+6])

		mMarginTradingArr[idx] = [7]int64{
			dateIntFmt, buyMoney, payMoney, balanceMoney, saleStockVol, payStockVol, balanceStockVol,
		}

		jdx += 7
		idx++

	}

	mMarginT.MarginTrading = mMarginTradingArr
	//log.Println(mMarginT.MarginTrading)

	return mMarginT

}
