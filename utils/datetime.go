package utils

import (
	"log"
	"strings"
	"sync"
	"time"
)

const intDateLayout = "20060102"
const hyphenDateLayout = "2006-01-02"

var dateBegin string = "2005-01-02"

func SetDateBegin(dateValue string) {
	dateBegin = dateValue
}

func GetDateArray(startDate time.Time) []string {

	var wg sync.WaitGroup

	dateArray := make([]string, 5000)
	ch := make(chan string, 1)

	go generateNextDateChannel(startDate, ch, wg)

	idx := 0
	for true {
		if val, ok := <-ch; ok {
			dateArray[idx] = val
			idx++
		} else {
			break
		}
	}

	wg.Wait()

	dateArray = dateArray[0:idx]

	return dateArray
}

func ParseDateWithIntFmt(intFmtDate string) time.Time {
	parseDate, _ := time.Parse(intDateLayout, intFmtDate)
	return parseDate
}

func generateNextDateChannel(startDate time.Time, ch chan string, wg sync.WaitGroup) {

	oneday, _ := time.ParseDuration("24h")
	negOneday, _ := time.ParseDuration("-24h")

	wg.Add(1)
	defer wg.Done()

	now := time.Now()
	if !(now.Hour() > 18 && now.Minute() > 15) {
		now = now.Add(negOneday)
	}

	nowStr := strings.Fields(now.Format("2006-01-02 15:04:05"))[0]
	log.Println("生成的日期截止于: ", nowStr)

	beginDate, _ := time.Parse(hyphenDateLayout, dateBegin)
	if startDate.After(beginDate) {
		beginDate = startDate
	}
	log.Println("初始日期：", beginDate)

	for !beginDate.After(now) {

		weekday := beginDate.Weekday()
		switch weekday {
		case time.Saturday:
			beginDate = beginDate.Add(oneday)
		case time.Sunday:
			beginDate = beginDate.Add(oneday)
		default:
			if !isHoliday(beginDate) {
				beginDateStr := strings.Fields(beginDate.Format("2006-01-02 15:04:05"))[0]
				ch <- strings.Split(beginDateStr, " ")[0]
			}
			beginDate = beginDate.Add(oneday)
		}

	}

	close(ch)

}

func isHoliday(checkDate time.Time) bool {

	_, month, day := checkDate.Date()

	isNewYearDay := month == time.January && (day == 1)

	isLaborDay := month == time.May && (day == 1 || day == 2 || day == 3)

	isNationalDay := month == time.October &&
		(day == 1 || day == 2 || day == 3 || day == 4 || day == 5 || day == 6 || day == 7)

	return isNewYearDay || isLaborDay || isNationalDay

}
