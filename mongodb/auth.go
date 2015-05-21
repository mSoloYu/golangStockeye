package mongodb

import (
	"time"

	"../utils"

	"gopkg.in/mgo.v2"
)

const stockdb = "stock"
const stocktrandb = "stocktran"
const funddb = "fund"

const stockCollection = "stock_info"
const stockPublicHolderCollection = "stock_public_holder"
const stockMainHolderCollection = "stock_main_holder"
const stockAccountingCollection = "stock_accounting"
const stockFundingCollection = "stock_funding"
const stockBigDealCollection = "stock_big_deal"
const stockMarginTradingCollection = "stock_margin_trading"
const fundCollection = "fund_info"
const fundtranCollection = "fund_tran"

var authUser *mgo.DialInfo = new(mgo.DialInfo)
var mgoSession *mgo.Session
var mgoColl *mgo.Collection

func init() {

	authUser.Addrs = []string{"192.168.1.105:9988"}
	authUser.Direct = false
	authUser.Timeout = time.Minute
	authUser.FailFast = true
	//authUser.Source
	//authUser.Service
	//authUser.Mechanism
	//DialServer func(addr *ServerAddr) (net.Conn, error)
}

func connectToStockDbReadwrite(db string) {

	if mgoSession != nil {
		mgoSession.Close()
	}

	authUser.Username = "stockrw"
	authUser.Password = "STOCK@#!34406eyeZz"
	authUser.Database = db

	session, err := mgo.DialWithInfo(authUser)
	utils.CheckError(err)
	mgoSession = session

}

func connectToStockDbReadonly(db string) {

	authUser.Username = "stockro"
	authUser.Password = "123456"
	authUser.Database = db

	session, err := mgo.DialWithInfo(authUser)
	utils.CheckError(err)
	mgoSession = session

}

func connectToFundDbReadwrite(db string) {

	if mgoSession != nil {
		mgoSession.Close()
	}

	authUser.Username = "fundrw"
	authUser.Password = "STOCK@#!34406eyeZz"
	authUser.Database = db

	session, err := mgo.DialWithInfo(authUser)
	utils.CheckError(err)
	mgoSession = session

}

func connectToFundDbReadonly(db string) {

	authUser.Username = "fundro"
	authUser.Password = "123456"
	authUser.Database = db

	session, err := mgo.DialWithInfo(authUser)
	utils.CheckError(err)
	mgoSession = session

}

func connectToStockCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockCollection)
	return mgoCollection

}

func connectToStockPublicHolderCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockPublicHolderCollection)
	return mgoCollection

}

func connectToStockMainHolderCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockMainHolderCollection)
	return mgoCollection

}

func connectToStockAccountingCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockAccountingCollection)
	return mgoCollection

}

func connectToStockFundingCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockFundingCollection)
	return mgoCollection

}

func connectToStockMarginTradingCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockMarginTradingCollection)
	return mgoCollection

}

func connectToStockBigDealCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(stockBigDealCollection)
	return mgoCollection

}

func connectToStockTranCollection(coll string) *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(coll)
	return mgoCollection

}

func connectToFundCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(fundCollection)
	return mgoCollection

}

func connectToFundTranCollection() *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(fundtranCollection)
	return mgoCollection

}
