package mongodb

import (
	//"log"
	"strings"

	"../netservice"
	"../utils"
)

//func makeFundBasicDoc(fundcode string) {
func makeFundBasicDoc(fundcode string) FundBasic {

	//netservice.MakeFundBasicArray(fundcode)
	infoArray := netservice.MakeFundBasicArray(fundcode)

	var mFundBasic FundBasic
	mFundBasic.Code = fundcode

	if !strings.Contains(infoArray[11], ",") {
		mFundBasic.FetchDate = -1
		return mFundBasic
		//return
	}
	mFundBasic.FetchDate = utils.GetTodayWithIntFmt()

	mFundBasic.LegalName = infoArray[1]
	mFundBasic.Name = infoArray[3]
	mFundBasic.InvestType = infoArray[5]

	publishDate := utils.GetIntVal(infoArray[7])
	RaiseStartDate := utils.GetIntVal(infoArray[9])
	RaiseStopDate := utils.GetIntVal(infoArray[13])
	DailyBuyStartDate := utils.GetIntVal(infoArray[17])
	DailyRedeemStartDate := utils.GetIntVal(infoArray[21])
	mFundBasic.Date = []int{publishDate, RaiseStartDate, RaiseStopDate,
		DailyBuyStartDate, DailyRedeemStartDate}

	totalVol := utils.GetInt64Val(strings.Replace(infoArray[11], ",", "", -1)) * 100
	realVol := utils.GetInt64Val(strings.Replace(infoArray[15], ",", "", -1)) * 100
	managerVol := utils.GetInt64Val(strings.Replace(infoArray[19], ",", "", -1)) * 100
	interestMoney := utils.GetInt64Val(infoArray[23]) * 100
	newestTotalVol := utils.GetInt64Val(strings.Replace(infoArray[27], ",", "", -1)) * 100
	mFundBasic.SharesVol = []int64{totalVol, realVol, managerVol, interestMoney, newestTotalVol}

	mFundBasic.Manager = infoArray[25]
	mFundBasic.FundManager = infoArray[29]
	mFundBasic.FundHolder = infoArray[31]
	mFundBasic.InvestStyle = infoArray[33]
	mFundBasic.RateStructure = infoArray[35]
	mFundBasic.ResultCmp = infoArray[37]
	mFundBasic.InvestTarget = infoArray[39]
	mFundBasic.InvestPhilosophy = infoArray[41]
	mFundBasic.InvestRange = infoArray[43]
	mFundBasic.InvestStrategy = infoArray[45]
	mFundBasic.InvestStandard = infoArray[47]
	mFundBasic.RiskFeature = infoArray[49]
	mFundBasic.RiskManageTool = infoArray[51]

	//log.Println("--->", mFundBasic)
	return mFundBasic

}
