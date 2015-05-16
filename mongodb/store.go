package mongodb

import (
	"log"

	"gopkg.in/mgo.v2"
)

var mgoColl *mgo.Collection

func StoreStockModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockCollection()

	for idx, stockcode := range stockcodeArr {
		stockDoc := makeStockModel(stockcode)

		log.Printf("%4d - %s : %s", idx, stockDoc.Code, stockDoc.FullName)
		coll.Insert(stockDoc)
	}

}

func StoreStockPublicHolderModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockPublicHolderCollection()

	for idx, stockcode := range stockcodeArr {
		publicHolderDoc := makeStockPublicHolderDoc(stockcode)

		log.Printf("----> %4d, %s", idx, stockcode)
		coll.Insert(publicHolderDoc)
	}

}

func StoreStockMainHolderModel(stockcodeArr []string) {

	//connectToStockDbReadwrite(stockdb)
	//coll := connectToStockMainHolderCollection()

	for _, stockcode := range stockcodeArr {
		stockcode = "601668"
		makeStockMainHolderDoc(stockcode)
		break
		//mainHolderDoc := makeStockMainHolderDoc(stockcode)

		//log.Printf("----> %4d, %s", idx, stockcode)
		//coll.Insert(mainHolderDoc)
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
