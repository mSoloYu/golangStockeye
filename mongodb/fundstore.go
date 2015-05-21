package mongodb

import (
	"log"
)

func StoreFundBasicModel(fundcodeArr []string) {

	connectToFundDbReadwrite(funddb)
	coll := connectToFundCollection()

	for idx, fundcode := range fundcodeArr {
		fundDoc := makeFundBasicDoc(fundcode)
		if fundDoc.FetchDate == -1 {
			log.Printf("%4d - %s : %s", idx, fundDoc.Code, " --- skip")
			continue
		}

		log.Printf("%4d - %s : %s", idx, fundDoc.Code, fundDoc.Name)
		coll.Insert(fundDoc)

		//makeFundBasicDoc(fundcode)
		//break
	}

}

func StoreFundHoldingStockModel(fundcodeArr []string) {

	//connectToFundDbReadwrite(funddb)
	//coll := connectToFundTranCollection()

	for _, fundcode := range fundcodeArr {
		makeFundTranDoc(fundcode)
		break
		//fundtranDoc := makeFundTranDoc(fundcode)

		//log.Printf("%4d - %s : %s", idx, fundtranDoc.Code)
		//coll.Insert(fundtranDoc)
	}

}
