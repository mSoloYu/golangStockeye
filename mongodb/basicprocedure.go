package mongodb

import (
	"strings"

	"../netservice"
	"../utils"
)

func makeStockModel(stockcode string) *Stock {

	stockArray := netservice.MakeStockBasicArray(stockcode)

	stock := new(Stock)

	stock.SimpleName = strings.TrimSpace(stockArray[1])
	stock.Code = strings.TrimSpace(stockArray[3])
	stock.FullName = strings.TrimSpace(stockArray[5])
	stock.EnglishName = strings.TrimSpace(stockArray[7])
	stock.OldName = strings.TrimSpace(stockArray[9])
	stock.PublishDate = utils.ParseIntByReplace(stockArray[11], "-")
	stock.ForIndustry = strings.TrimSpace(stockArray[13])
	stock.ForStockConcept = strings.Split(strings.TrimSpace(stockArray[15]), "、")
	stock.Area = strings.TrimSpace(stockArray[17])
	stock.LegalOwner = strings.TrimSpace(stockArray[19])
	stock.IndependentDirector = strings.Split(strings.TrimSpace(stockArray[21]), "、")
	stock.AdvisoryOrga = strings.TrimSpace(stockArray[23])
	stock.AccountingOrga = strings.TrimSpace(stockArray[25])
	stock.SecuritiesRepresentative = strings.TrimSpace(stockArray[27])

	stock.Capital = utils.ParseFloatByReplace(strings.Replace(stockArray[29], ",", "", -1), "万元")

	stock.RegisterAddr = strings.TrimSpace(stockArray[31])
	stock.RateOfTax = utils.ParseFloatByReplace(stockArray[33], "%")

	stock.BusinessAddr = strings.TrimSpace(stockArray[35])
	stock.MainProduct = strings.TrimSpace(stockArray[37])
	stock.IssueDate = utils.ParseIntByReplace(stockArray[39], "-")
	stock.OpenSaleDate = utils.ParseIntByReplace(stockArray[41], "-")
	stock.WhichMarket = strings.TrimSpace(stockArray[43])
	stock.SecurityType = strings.TrimSpace(stockArray[45])

	stock.OutstandingCapitalStock = utils.ParseFloatByReplace(strings.Replace(stockArray[47], ",", "", -1), "万股")
	stock.TotalCapitalStock = utils.ParseFloatByReplace(strings.Replace(stockArray[49], ",", "", -1), "万股")

	stock.SallerAgent = strings.TrimSpace(stockArray[51])
	stock.IssuePrice = utils.ParseFloatByReplace(stockArray[53], "元")
	stock.OpenSaleOpenPrice = utils.ParseFloatByReplace(stockArray[55], "元")
	stock.OpenSalePriceRate = utils.ParseFloatByReplace(stockArray[57], "%")
	stock.OpenSaleExchangeRage = utils.ParseFloatByReplace(stockArray[59], "%")
	stock.SpecialIssue = strings.TrimSpace(stockArray[61])
	stock.IssuePE = utils.ParseFloatByReplace(stockArray[63], "倍")
	stock.CurrentPE = utils.ParseFloatByReplace(stockArray[65], "倍")
	stock.Tel = strings.TrimSpace(stockArray[67])
	stock.Email = strings.TrimSpace(stockArray[71])
	stock.WebSite = strings.TrimSpace(stockArray[73])
	stock.ContactPerson = strings.TrimSpace(stockArray[75])

	return stock

}
