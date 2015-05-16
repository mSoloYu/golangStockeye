package mongodb

import (
	"fmt"
	"strconv"
	"strings"

	"../netservice"
	"../utils"
)

func makeStockPublicHolderDoc(stockcode string) StockPublicHolder {

	publicHolderInfos := netservice.MakeStockHoldersArray(stockcode, netservice.HolderTypePublic)
	size := len(publicHolderInfos)
	fetchDate := utils.GetTodayWithIntFmt()
	mPublicArr := make([]PublicStruct, size)
	mPublicHolder := StockPublicHolder{stockcode, fetchDate, mPublicArr}

	for idx, info := range publicHolderInfos {
		recs := strings.SplitAfterN(info, "_", 2)
		intStopDate, _ := strconv.Atoi(strings.TrimRight(recs[0], "_"))
		mPublicArr[idx].StopDate = intStopDate
		holders := strings.Split(recs[1], "_")
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

	return mPublicHolder

}

//func MakeStockMainHolderDoc(stockcode string) mongodb.StockPublicHolder {
func makeStockMainHolderDoc(stockcode string) {

	mainHolderInfos := netservice.MakeStockHoldersArray(stockcode, netservice.HolderTypeMain)
	for _, info := range mainHolderInfos {
		fmt.Println(info)
	}

	/*
		size := len(mainHolderInfos)
		fetchDate := utils.GetTodayWithIntFmt()
		mPublicArr := make([]mongodb.PublicStruct, size)
		mPublicHolder := mongodb.StockPublicHolder{stockcode, fetchDate, mPublicArr}

		for idx, info := range publicHolderInfos {
			recs := strings.SplitAfterN(info, "_", 2)
			intStopDate, _ := strconv.Atoi(strings.TrimRight(recs[0], "_"))
			mPublicArr[idx].StopDate = intStopDate
			holders := strings.Split(recs[1], "_")
			size = len(holders)
			mHolderArr := make([]mongodb.HolderStruct, size/3)
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
	*/

	//return mPublicHolder

}
