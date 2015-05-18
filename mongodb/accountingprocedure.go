package mongodb

import (
	//"log"
	"strconv"

	"../netservice"
	"../utils"
)

//func makeStockAccountingDoc(stockcode string) {
func makeStockAccountingDoc(stockcode string) StockAccounting {

	//netservice.MakeStockAccountingArray(stockcode)
	rankingInfos, accountingInfos := netservice.MakeStockAccountingArray(stockcode)
	var mSa StockAccounting
	mSa.Code = stockcode
	mSa.Date = utils.GetTodayWithIntFmt()
	mSa.IndustryClass = []string{rankingInfos[0], rankingInfos[3]}

	standardMarketVal, _ := strconv.Atoi(rankingInfos[1])
	standardIncomeVal, _ := strconv.Atoi(rankingInfos[2])
	IcbMarketVal, _ := strconv.Atoi(rankingInfos[4])
	IcbIncomeVal, _ := strconv.Atoi(rankingInfos[5])
	mSa.Ranking = []int{standardMarketVal, standardIncomeVal, IcbMarketVal, IcbIncomeVal}

	stopDate, _ := strconv.Atoi(accountingInfos[0])
	mSa.StopDate = stopDate
	target := getFloatValue(accountingInfos[1])
	lastyearCmpVal := getFloatValue(accountingInfos[2])
	lastPeriodCmpVal := getFloatValue(accountingInfos[3])
	mSa.EarningPerShare = []float32{target, lastyearCmpVal, lastPeriodCmpVal}
	target = getFloatValue(accountingInfos[4])
	lastyearCmpVal = getFloatValue(accountingInfos[5])
	lastPeriodCmpVal = getFloatValue(accountingInfos[6])
	mSa.Naps = []float32{target, lastyearCmpVal, lastPeriodCmpVal}
	target = getFloatValue(accountingInfos[7])
	lastyearCmpVal = getFloatValue(accountingInfos[8])
	lastPeriodCmpVal = getFloatValue(accountingInfos[9])
	mSa.SalesRevenue = []float32{target, lastyearCmpVal, lastPeriodCmpVal}
	target = getFloatValue(accountingInfos[10])
	lastyearCmpVal = getFloatValue(accountingInfos[11])
	lastPeriodCmpVal = getFloatValue(accountingInfos[12])
	mSa.Profit = []float32{target, lastyearCmpVal, lastPeriodCmpVal}
	target = getFloatValue(accountingInfos[13])
	lastyearCmpVal = getFloatValue(accountingInfos[14])
	lastPeriodCmpVal = getFloatValue(accountingInfos[15])
	mSa.GrossProfitSale = []float32{target, lastyearCmpVal, lastPeriodCmpVal}
	target = getFloatValue(accountingInfos[16])
	lastyearCmpVal = getFloatValue(accountingInfos[17])
	lastPeriodCmpVal = getFloatValue(accountingInfos[18])
	mSa.Others = []float32{target, lastyearCmpVal, lastPeriodCmpVal}

	//log.Println(mSa)
	return mSa

}

func getFloatValue(str string) float32 {
	val, _ := strconv.ParseFloat(str, 64)
	return float32(val)
}
