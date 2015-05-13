package mongodb

import (
	"log"

	"../netservice"
	"gopkg.in/mgo.v2"
)

var mgoColl *mgo.Collection

func StoreStockModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)

	for idx, stockcode := range stockcodeArr {
		stockArray := netservice.MakeStockInfoArray(stockcode)
		stock := makeStockModel(stockArray)

		log.Printf("%4d - %s : %s", idx, stock.Code, stock.FullName)
		connectToStockCollection().Insert(stock)
	}

}

func ConnectToStockTranCollection(stockcode string) {
	connectToStockDbReadwrite(stocktrandb)
	mgoColl = connectToStockTranCollection(stockcode)
	index := mgo.Index{
		Key:        []string{"date"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	mgoColl.EnsureIndex(index)
}

func StoreStockTranDailyModel(stocktran *StockTran) {
	mgoColl.Insert(stocktran)
}
