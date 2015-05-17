package mongodb

import (
	//"log"
	"strconv"
	"strings"

	"../netservice"
	"../utils"
)

//func makeStockPublicHolderDoc(stockcode string) {
func makeStockPublicHolderDoc(stockcode string) StockPublicHolder {

	publicHolderInfos := netservice.MakeStockHoldersArray(stockcode, netservice.HolderTypePublic)
	size := len(publicHolderInfos)
	fetchDate := utils.GetTodayWithIntFmt()
	mPublicArr := make([]PublicStruct, size)
	mPublicHolder := StockPublicHolder{stockcode, fetchDate, mPublicArr}

	for idx, info := range publicHolderInfos {
		recs := strings.SplitAfterN(info, "|", 2)
		intStopDate, _ := strconv.Atoi(strings.TrimRight(recs[0], "|"))
		mPublicArr[idx].StopDate = intStopDate
		holders := strings.Split(recs[1], "|")
		size = len(holders)
		mHolderArr := make([]HolderStruct, size/3)
		mPublicArr[idx].PublicHolder = mHolderArr

		jdx, kdx := 0, 0
		for true {
			vol, _ := strconv.ParseInt(holders[kdx+1], 10, 64)
			percentage, _ := strconv.ParseFloat(holders[kdx+2], 32)

			mHolderArr[jdx].Name = holders[kdx]
			mHolderArr[jdx].Vol = vol
			mHolderArr[jdx].Percentage = float32(percentage)

			jdx++
			kdx += 3
			if kdx >= size {
				break
			}
		}

	}

	//log.Println(mPublicHolder)
	return mPublicHolder

}

//func makeStockMainHolderDoc(stockcode string) {
func makeStockMainHolderDoc(stockcode string) StockMainHolder {

	mainHolderInfos := netservice.MakeStockHoldersArray(stockcode, netservice.HolderTypeMain)
	size := len(mainHolderInfos)
	fetchDate := utils.GetTodayWithIntFmt()
	mMainArr := make([]MainStruct, size)
	mMainHolder := StockMainHolder{stockcode, fetchDate, mMainArr}

	for idx, info := range mainHolderInfos {
		recs := strings.SplitAfterN(info, "|", 4)
		intStopDate, _ := strconv.Atoi(strings.TrimRight(recs[0], "|"))
		numHolder, _ := strconv.Atoi(strings.TrimRight(recs[1], "|"))
		evenVol, _ := strconv.Atoi(strings.TrimRight(recs[2], "|"))

		mMainArr[idx].StopDate = intStopDate
		mMainArr[idx].NumHolder = numHolder
		mMainArr[idx].EvenVol = evenVol

		holders := strings.Split(recs[3], "|")
		size = len(holders)
		mHolderArr := make([]HolderStruct, size/3)
		mMainArr[idx].MainHolder = mHolderArr

		jdx, kdx := 0, 0
		for true {
			vol, _ := strconv.ParseInt(holders[kdx+1], 10, 64)
			percentage, _ := strconv.ParseFloat(holders[kdx+2], 32)

			mHolderArr[jdx].Name = holders[kdx]
			mHolderArr[jdx].Vol = vol
			mHolderArr[jdx].Percentage = float32(percentage)

			jdx++
			kdx += 3
			if kdx >= size {
				break
			}
		}

	}

	//log.Println(mMainHolder)
	return mMainHolder

}
