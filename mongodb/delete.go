package mongodb

import (
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func DeleteTranRecordWithDate(stockcode, trandate string) {

	connectToStockDbReadwrite(stocktrandb)
	mgoColl := connectToStockTranCollection(stockcode)

	date, _ := strconv.Atoi(strings.Replace(trandate, "-", "", -1))
	mgoColl.Remove(bson.M{"date": date})

}
