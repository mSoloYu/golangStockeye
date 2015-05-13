package mongodb

import (
	"strconv"
	"strings"
)

type Stock struct {
	SimpleName               string
	Code                     string
	FullName                 string
	EnglishName              string
	OldName                  string
	PublishDate              int
	ForIndustry              string
	ForStockConcept          []string
	Area                     string
	LegalOwner               string
	IndependentDirector      []string
	AdvisoryOrga             string
	AccountingOrga           string
	SecuritiesRepresentative string
	Capital                  float32
	RegisterAddr             string
	RateOfTax                float32
	BusinessAddr             string
	MainProduct              string
	IssueDate                int
	OpenSaleDate             int
	WhichMarket              string
	SecurityType             string
	OutstandingCapitalStock  float32
	TotalCapitalStock        float32
	SallerAgent              string
	IssuePrice               float32
	OpenSaleOpenPrice        float32
	OpenSalePriceRate        float32
	OpenSaleExchangeRage     float32
	SpecialIssue             string
	IssuePE                  float32
	CurrentPE                float32
	Tel                      string
	Email                    string
	WebSite                  string
	ContactPerson            string
}

type FifteenMinuteTran struct {
	Time     string
	Count    []int // tranCount, priceChangeCount
	VolMoney []int64
}

type EachPriceTran struct {
	Time     string
	Price    float32
	VolMoney []int64
}

type StockTran struct {
	Date       int
	KeyPrices  []float32 // last, open, close, highest, lowest
	Count      []int     // priceCount, tranCount
	VolMoney   []int64
	Fifteen    []FifteenMinuteTran
	EachPrices []EachPriceTran
}

func makeStockModel(stockArray []string) *Stock {

	stock := new(Stock)

	stock.SimpleName = strings.TrimSpace(stockArray[1])
	stock.Code = strings.TrimSpace(stockArray[3])
	stock.FullName = strings.TrimSpace(stockArray[5])
	stock.EnglishName = strings.TrimSpace(stockArray[7])
	stock.OldName = strings.TrimSpace(stockArray[9])
	stock.PublishDate = parseInt(stockArray[11], "-")
	stock.ForIndustry = strings.TrimSpace(stockArray[13])
	stock.ForStockConcept = strings.Split(strings.TrimSpace(stockArray[15]), "、")
	stock.Area = strings.TrimSpace(stockArray[17])
	stock.LegalOwner = strings.TrimSpace(stockArray[19])
	stock.IndependentDirector = strings.Split(strings.TrimSpace(stockArray[21]), "、")
	stock.AdvisoryOrga = strings.TrimSpace(stockArray[23])
	stock.AccountingOrga = strings.TrimSpace(stockArray[25])
	stock.SecuritiesRepresentative = strings.TrimSpace(stockArray[27])

	stock.Capital = parseFloat(strings.Replace(stockArray[29], ",", "", -1), "万元")

	stock.RegisterAddr = strings.TrimSpace(stockArray[31])
	stock.RateOfTax = parseFloat(stockArray[33], "%")

	stock.BusinessAddr = strings.TrimSpace(stockArray[35])
	stock.MainProduct = strings.TrimSpace(stockArray[37])
	stock.IssueDate = parseInt(stockArray[39], "-")
	stock.OpenSaleDate = parseInt(stockArray[41], "-")
	stock.WhichMarket = strings.TrimSpace(stockArray[43])
	stock.SecurityType = strings.TrimSpace(stockArray[45])

	stock.OutstandingCapitalStock = parseFloat(strings.Replace(stockArray[47], ",", "", -1), "万股")
	stock.TotalCapitalStock = parseFloat(strings.Replace(stockArray[49], ",", "", -1), "万股")

	stock.SallerAgent = strings.TrimSpace(stockArray[51])
	stock.IssuePrice = parseFloat(stockArray[53], "元")
	stock.OpenSaleOpenPrice = parseFloat(stockArray[55], "元")
	stock.OpenSalePriceRate = parseFloat(stockArray[57], "%")
	stock.OpenSaleExchangeRage = parseFloat(stockArray[59], "%")
	stock.SpecialIssue = strings.TrimSpace(stockArray[61])
	stock.IssuePE = parseFloat(stockArray[63], "倍")
	stock.CurrentPE = parseFloat(stockArray[65], "倍")
	stock.Tel = strings.TrimSpace(stockArray[67])
	stock.Email = strings.TrimSpace(stockArray[71])
	stock.WebSite = strings.TrimSpace(stockArray[73])
	stock.ContactPerson = strings.TrimSpace(stockArray[75])

	return stock

}

func parseInt(targetStr, replaceStr string) int {
	val, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(targetStr), replaceStr, "", -1))
	return val
}

func parseFloat(targetStr, replaceStr string) float32 {

	val, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(targetStr), replaceStr, "", -1), 32)
	return float32(val)

}
