package mongodb

import (
	//"log"
	"strconv"

	"../netservice"
	"../utils"
)

//func makeStockFundingDoc(stockcode string) {
func makeStockFundingDoc(stockcode string) StockFunding {

	fundingInfos := netservice.MakeStockFundingArray(stockcode)

	var mSf StockFunding
	if len(fundingInfos) == 1 {
		mSf.Date = -1
		return mSf
		//return
	}

	mSf.Code = stockcode
	mSf.Date = utils.GetTodayWithIntFmt()
	fetchDate, _ := strconv.Atoi(fundingInfos[0])
	mSf.FetchDate = fetchDate

	totalFundings, _ := strconv.Atoi(fundingInfos[1])
	newFundings, _ := strconv.Atoi(fundingInfos[2])
	addFundings, _ := strconv.Atoi(fundingInfos[3])
	subFundings, _ := strconv.Atoi(fundingInfos[4])
	quitFundings, _ := strconv.Atoi(fundingInfos[5])
	mSf.FundingCounting = []int{totalFundings, newFundings, addFundings, subFundings, quitFundings}

	totalVolOfFundings, _ := strconv.Atoi(fundingInfos[6])
	totalVolChangeOfFundings, _ := strconv.Atoi(fundingInfos[7])
	totalVolPercentageOfFundings, _ := strconv.Atoi(fundingInfos[8])
	mSf.VolCounting = []int{totalVolOfFundings, totalVolChangeOfFundings, totalVolPercentageOfFundings}

	mTenFundingArr := make([]FundingStruct, 10)
	mTenFundingArr[0].Name = fundingInfos[10]
	mTenFundingArr[1].Name = fundingInfos[11]
	mTenFundingArr[2].Name = fundingInfos[12]
	mTenFundingArr[3].Name = fundingInfos[13]
	mTenFundingArr[4].Name = fundingInfos[14]
	mTenFundingArr[5].Name = fundingInfos[15]
	mTenFundingArr[6].Name = fundingInfos[16]
	mTenFundingArr[7].Name = fundingInfos[17]
	mTenFundingArr[8].Name = fundingInfos[18]
	mTenFundingArr[9].Name = fundingInfos[19]

	mTenFundingArr[0].Type = fundingInfos[20]
	vol := utils.GetIntVal(fundingInfos[21])
	mTenFundingArr[0].Vol = vol
	mTenFundingArr[1].Type = fundingInfos[22]
	vol = utils.GetIntVal(fundingInfos[23])
	mTenFundingArr[1].Vol = vol
	mTenFundingArr[2].Type = fundingInfos[24]
	vol = utils.GetIntVal(fundingInfos[25])
	mTenFundingArr[2].Vol = vol
	mTenFundingArr[3].Type = fundingInfos[26]
	vol = utils.GetIntVal(fundingInfos[27])
	mTenFundingArr[3].Vol = vol
	mTenFundingArr[4].Type = fundingInfos[28]
	vol = utils.GetIntVal(fundingInfos[29])
	mTenFundingArr[4].Vol = vol
	mTenFundingArr[5].Type = fundingInfos[30]
	vol = utils.GetIntVal(fundingInfos[31])
	mTenFundingArr[5].Vol = vol
	mTenFundingArr[6].Type = fundingInfos[32]
	vol = utils.GetIntVal(fundingInfos[33])
	mTenFundingArr[6].Vol = vol
	mTenFundingArr[7].Type = fundingInfos[34]
	vol = utils.GetIntVal(fundingInfos[35])
	mTenFundingArr[7].Vol = vol
	mTenFundingArr[8].Type = fundingInfos[36]
	vol = utils.GetIntVal(fundingInfos[37])
	mTenFundingArr[8].Vol = vol
	mTenFundingArr[9].Type = fundingInfos[38]
	vol = utils.GetIntVal(fundingInfos[39])
	mTenFundingArr[9].Vol = vol
	mSf.TenFundingChange = mTenFundingArr

	size := len(fundingInfos)
	detailLen := (size - 40) / 6
	mFundDetailArr := make([]FundingDetailStruct, detailLen)
	idx, counter := 0, 40
	totalShares, changeShares, percentageOfFunding := 0, 0, 0
	for true {
		if idx == detailLen {
			break
		}
		mFundDetailArr[idx].Name = fundingInfos[counter]
		mFundDetailArr[idx].Code = fundingInfos[counter+1]
		totalShares = utils.GetIntVal(fundingInfos[counter+2])
		changeShares = utils.GetIntVal(fundingInfos[counter+3])
		percentageOfFunding = utils.GetIntVal(fundingInfos[counter+5])
		mFundDetailArr[idx].VolDetail = []int{totalShares, changeShares, percentageOfFunding}
		counter += 6
		idx++
	}
	mSf.FundingDetail = mFundDetailArr

	//log.Println("--->", mSf)
	return mSf

}
