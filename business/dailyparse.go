package business

import (
	//"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"../mongodb"
)

var sepTimeStamp = []string{"14:45:00", "14:30:00", "14:15:00", "14:00:00",
	"13:45:00", "13:30:00", "13:15:00", "13:00:00",
	"11:15:00", "11:00:00", "10:45:00", "10:30:00",
	"10:15:00", "10:00:00", "09:45:00", "09:30:00"}

func initMemory(tranSize int) (
	times []string, prices []float32, vols, moneys []int64,
	timePriceMap map[string]string, pricePriceMap map[string]float32,
	volPriceMap, moneyPriceMap map[string]int64,
	timeTimeMap map[string]int, priceTimeMap map[string]string, priceCountTimeMap map[string]int,
	volTimeMap, moneyTimeMap map[string]int64) {

	times = make([]string, tranSize)
	prices = make([]float32, tranSize)
	vols = make([]int64, tranSize)
	moneys = make([]int64, tranSize)

	timePriceMap = make(map[string]string)
	pricePriceMap = make(map[string]float32)
	volPriceMap = make(map[string]int64)
	moneyPriceMap = make(map[string]int64)

	timeTimeMap = make(map[string]int)
	priceTimeMap = make(map[string]string)
	priceCountTimeMap = make(map[string]int)
	volTimeMap = make(map[string]int64)
	moneyTimeMap = make(map[string]int64)

	return

}

// ParseDaily complete the statistical transaction information with below task:
// 1. basic : open, close, highest, lowest, vol, money, priceCount, tranCount
// 2. each 15 minute : tranCount, priceChangeCount, vol, money
// 3. each price : firstTime, price, vol, money
func ParseDaily(lastdayPrice float32, date, recordCollection string) *mongodb.StockTran {

	stockTran := new(mongodb.StockTran)

	records := strings.Split(recordCollection, "\n")
	sort.Strings(records)
	records = records[1 : len(records)-1]

	tranSize := len(records)

	tranTimes, tranPrices, tranVols, tranMoneys,
		tranTimePriceMap, tranPricePriceMap, tranVolPriceMap, tranMoneyPriceMap,
		tranTimeTimeMap, tranPriceTimeMap, tranPriceCountTimeMap,
		tranVolTimeMap, tranMoneyTimeMap := initMemory(tranSize)

	openPrice, _ := strconv.ParseFloat(regexp.MustCompile("\\s").Split(records[0], -1)[1], 32)
	closePrice, _ := strconv.ParseFloat(regexp.MustCompile("\\s").Split(records[tranSize-1], -1)[1], 32)

	highestPrice, lowestPrice := 0, 100000
	var totalVol, totalMoney int64
	for idx, rec := range records {

		fields := regexp.MustCompile("\\s").Split(rec, -1)

		trantime, price, vol, money :=
			setUpRecordFields(idx, fields, tranTimes, tranPrices, tranVols, tranMoneys)
		totalVol += vol
		totalMoney += money

		intPrice, _ := strconv.Atoi(strings.Replace(price, ".", "", -1))
		if highestPrice < intPrice {
			highestPrice = intPrice
		}
		if lowestPrice > intPrice {
			lowestPrice = intPrice
		}

		if _, ok := tranPricePriceMap[price]; ok {
			doClassifyByPirce(vol, money, price, tranVolPriceMap, tranMoneyPriceMap)
		} else {
			setUpMaps(vol, money, tranPrices[idx], price, trantime,
				tranTimePriceMap, tranPricePriceMap, tranVolPriceMap, tranMoneyPriceMap)
		}

		setUpFifteenMinRecordFields(vol, money, price, trantime,
			tranTimeTimeMap, tranPriceTimeMap, tranPriceCountTimeMap, tranVolTimeMap, tranMoneyTimeMap)

	}

	priceSize := len(tranPricePriceMap)

	intFmtDate, _ := strconv.Atoi(strings.Replace(date, "-", "", -1))
	stockTran.Date = intFmtDate
	stockTran.KeyPrices = []float32{lastdayPrice,
		float32(openPrice), float32(closePrice), float32(highestPrice) / 100, float32(lowestPrice) / 100}
	stockTran.Count = []int{priceSize, tranSize}
	stockTran.VolMoney = []int64{totalVol, totalMoney}
	stockTran.Fifteen = getFifteen(tranTimeTimeMap, tranPriceTimeMap,
		tranPriceCountTimeMap, tranVolTimeMap, tranMoneyTimeMap)
	stockTran.EachPrices = getEachPrices(priceSize, tranTimePriceMap,
		tranPricePriceMap, tranVolPriceMap, tranMoneyPriceMap)

	//test output
	//log.Println(stockTran)

	return stockTran

}

func getFifteen(tranTimeTimeMap map[string]int, tranPriceTimeMap map[string]string,
	tranPriceCountTimeMap map[string]int,
	tranVolTimeMap, tranMoneyTimeMap map[string]int64) []mongodb.FifteenMinuteTran {

	size := len(sepTimeStamp) - 1
	fifMinTran := make([]mongodb.FifteenMinuteTran, size+1)
	for i, timeStamp := range sepTimeStamp {
		idx := size - i
		fifMinTran[idx].Time = timeStamp
		tranCount, _ := tranTimeTimeMap[timeStamp]
		priceChangeCount, _ := tranPriceCountTimeMap[timeStamp]
		vol, _ := tranVolTimeMap[timeStamp]
		money, _ := tranMoneyTimeMap[timeStamp]
		fifMinTran[idx].Count = []int{priceChangeCount, tranCount}
		fifMinTran[idx].VolMoney = []int64{vol, money}
	}

	return fifMinTran

}

func getEachPrices(size int, tranTimePriceMap map[string]string, tranPricePriceMap map[string]float32,
	tranVolPriceMap, tranMoneyPriceMap map[string]int64) []mongodb.EachPriceTran {

	eachPrices := make([]mongodb.EachPriceTran, size)
	idx := 0
	for k, v := range tranPricePriceMap {
		tranTime, _ := tranTimePriceMap[k]
		vol, _ := tranVolPriceMap[k]
		money, _ := tranMoneyPriceMap[k]
		eachPrices[idx].Time = tranTime
		eachPrices[idx].Price = v
		eachPrices[idx].VolMoney = []int64{vol, money}
		idx++
	}

	return eachPrices

}

func doClassifyByPirce(vol, money int64, price string,
	tranVolPriceMap, tranMoneyPriceMap map[string]int64) {

	curPriceVol, _ := tranVolPriceMap[price]
	tranVolPriceMap[price] = curPriceVol + vol
	curPriceMoney, _ := tranMoneyPriceMap[price]
	tranMoneyPriceMap[price] = curPriceMoney + money

}

func getFakeCompareTime(checkTime int) (fakeTime string) {

	switch {
	case checkTime > 144500:
		fakeTime = sepTimeStamp[0]
	case checkTime > 143000:
		fakeTime = sepTimeStamp[1]
	case checkTime > 141500:
		fakeTime = sepTimeStamp[2]
	case checkTime > 140000:
		fakeTime = sepTimeStamp[3]
	case checkTime > 134500:
		fakeTime = sepTimeStamp[4]
	case checkTime > 133000:
		fakeTime = sepTimeStamp[5]
	case checkTime > 131500:
		fakeTime = sepTimeStamp[6]
	case checkTime > 120000:
		fakeTime = sepTimeStamp[7]
	case checkTime > 111500:
		fakeTime = sepTimeStamp[8]
	case checkTime > 110000:
		fakeTime = sepTimeStamp[9]
	case checkTime > 104500:
		fakeTime = sepTimeStamp[10]
	case checkTime > 103000:
		fakeTime = sepTimeStamp[11]
	case checkTime > 101500:
		fakeTime = sepTimeStamp[12]
	case checkTime > 100000:
		fakeTime = sepTimeStamp[13]
	case checkTime > 94500:
		fakeTime = sepTimeStamp[14]
	default:
		fakeTime = sepTimeStamp[15]
	}

	return

}

func setUpFifteenMinRecordFields(vol, money int64, price, trantime string,
	tranTimeTimeMap map[string]int, tranPriceTimeMap map[string]string,
	tranPriceCountTimeMap map[string]int, tranVolTimeMap, tranMoneyTimeMap map[string]int64) {

	intFmtTime, _ := strconv.Atoi(strings.Replace(trantime, ":", "", -1))
	strTime := getFakeCompareTime(intFmtTime)
	if _, ok := tranTimeTimeMap[strTime]; ok {
		comparePrice, _ := tranPriceTimeMap[strTime]
		if !strings.Contains(price, comparePrice) {
			tranPriceTimeMap[strTime] = price
			counter, _ := tranPriceCountTimeMap[strTime]
			tranPriceCountTimeMap[strTime] = counter + 1
		}
		tranCount, _ := tranTimeTimeMap[strTime]
		tranTimeTimeMap[strTime] = tranCount + 1
		curTimeVol, _ := tranVolTimeMap[strTime]
		tranVolTimeMap[strTime] = curTimeVol + vol
		curTimeMoney, _ := tranMoneyTimeMap[strTime]
		tranMoneyTimeMap[strTime] = curTimeMoney + money

		return
	}

	tranTimeTimeMap[strTime] = 1
	tranPriceTimeMap[strTime] = price
	tranPriceCountTimeMap[strTime] = 1
	tranVolTimeMap[strTime] = vol
	tranMoneyTimeMap[strTime] = money

}

func setUpMaps(vol, money int64, priceFloat float32, price, time string,
	tranTimePriceMap map[string]string, tranPricePriceMap map[string]float32,
	tranVolPriceMap, tranMoneyPriceMap map[string]int64) {

	tranTimePriceMap[price] = time
	tranPricePriceMap[price] = priceFloat
	tranVolPriceMap[price] = vol
	tranMoneyPriceMap[price] = money

}

func setUpRecordFields(idx int, fields, tranTimes []string, tranPrices []float32, tranVols,
	tranMoneys []int64) (trantime, price string, vol, money int64) {

	trantime = strings.TrimSpace(fields[0])
	tranTimes[idx] = trantime

	volParsed, _ := strconv.ParseInt(strings.TrimSpace(fields[3]), 10, 64)
	vol = volParsed
	tranVols[idx] = vol
	moneyParsed, _ := strconv.ParseInt(strings.TrimSpace(fields[4]), 10, 64)
	money = moneyParsed
	tranMoneys[idx] = money

	price = strings.TrimSpace(fields[1])
	if (len(price) - strings.Index(price, ".")) == 2 {
		price = price + "0"
	}
	priceFloat, _ := strconv.ParseFloat(price, 32)
	tranPrices[idx] = float32(priceFloat)

	return

}
