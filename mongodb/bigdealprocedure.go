package mongodb

import (
	//"log"

	"../netservice"
	"../utils"
)

//func makeStockBigDealDoc(stockcode string) {
func makeStockBigDealDoc(stockcode string) StockBigDeal {

	//netservice.MakeStockBigDealArray(stockcode)
	bigDealInfos := netservice.MakeStockBigDealArray(stockcode)

	var mBigDeal StockBigDeal
	mBigDeal.Code = stockcode
	mBigDeal.FetchDate = utils.GetTodayWithIntFmt()

	size := len(bigDealInfos)
	recSize := size / 6
	mBigDealArr := make([]BigDealStruct, recSize)

	idx, jdx := 0, 0
	for true {
		if idx >= recSize {
			break
		}

		mBigDealArr[idx].Date = utils.GetIntVal(bigDealInfos[jdx+0])

		buyPrice := utils.GetFloat32Val(bigDealInfos[jdx+1])
		buyVol := utils.GetFloat32Val(bigDealInfos[jdx+2])
		buyMoney := utils.GetFloat32Val(bigDealInfos[jdx+3])
		mBigDealArr[idx].Deal = []float32{buyPrice, buyVol, buyMoney}

		mBigDealArr[idx].Buyer = bigDealInfos[jdx+4]
		mBigDealArr[idx].Saller = bigDealInfos[jdx+5]

		jdx += 6
		idx++

	}

	mBigDeal.BigDeal = mBigDealArr
	//log.Println(mBigDeal)

	return mBigDeal

}
