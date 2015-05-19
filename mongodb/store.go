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
		//stockcode = "601188"
		//makeStockPublicHolderDoc(stockcode)
		//break
		publicHolderDoc := makeStockPublicHolderDoc(stockcode)

		log.Printf("----> %4d, %s", idx, stockcode)
		coll.Insert(publicHolderDoc)
	}

}

func StoreStockMainHolderModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockMainHolderCollection()

	for idx, stockcode := range stockcodeArr {
		//stockcode = "000018"
		//makeStockMainHolderDoc(stockcode)
		//break
		mainHolderDoc := makeStockMainHolderDoc(stockcode)

		log.Printf("----> %4d, %s", idx, stockcode)
		coll.Insert(mainHolderDoc)
	}

}

func StoreStockAccountingModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockAccountingCollection()

	for idx, stockcode := range stockcodeArr {
		//stockcode = "600023"
		//makeStockAccountingDoc(stockcode)
		//break
		accountingDoc := makeStockAccountingDoc(stockcode)

		log.Printf("----> %4d, %s", idx, stockcode)
		coll.Insert(accountingDoc)
	}

}

func StoreStockFundingModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockFundingCollection()

	for idx, stockcode := range stockcodeArr {
		//stockcode = "600023"
		//makeStockFundingDoc(stockcode)
		//break
		fundingDoc := makeStockFundingDoc(stockcode)
		if fundingDoc.Date != -1 {
			coll.Insert(fundingDoc)
			log.Printf("----> %4d, %s", idx, stockcode)
		} else {
			log.Printf("----> %4d, %s -- skip", idx, stockcode)
		}
	}

}

func StoreStockMarginTradingModel(stockcodeArr []string) {

	connectToStockDbReadwrite(stockdb)
	coll := connectToStockMarginTradingCollection()

	for idx, stockcode := range stockcodeArr {
		//stockcode = "000001"
		//makeStockMarginTradingDoc(stockcode)
		//break
		marginTradingDoc := makeStockMarginTradingDoc(stockcode)

		log.Printf("----> %4d, %s", idx, stockcode)
		coll.Insert(marginTradingDoc)
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
