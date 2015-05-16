package mongodb

import (
	"time"

	"../utils"

	"gopkg.in/mgo.v2"
)

const stockdb = "stock"
const stocktrandb = "stocktran"

const stockCollection = "stock_info"
const stockPublicHolderCollection = "stock_public_holder"
const stockMainHolderCollection = "stock_main_holder"
const stockBigDealCollection = "stock_big_deal"

var stockUser *mgo.DialInfo = new(mgo.DialInfo)
var mgoSession *mgo.Session

func init() {

	stockUser.Addrs = []string{"192.168.1.105:9988"}
	stockUser.Direct = false
	stockUser.Timeout = time.Minute
	stockUser.FailFast = true
	//stockUser.Source
	//stockUser.Service
	//stockUser.Mechanism
	//DialServer func(addr *ServerAddr) (net.Conn, error)
}

func connectToStockDbReadwrite(db string) {

	if mgoSession != nil {
		mgoSession.Close()
	}

	stockUser.Username = "stockrw"
	stockUser.Password = "STOCK@#!34406eyeZz"
	stockUser.Database = db

	session, err := mgo.DialWithInfo(stockUser)
	utils.CheckError(err)
	mgoSession = session

}

func connectToStockDbReadonly(db string) {

	stockUser.Username = "stockro"
	stockUser.Password = "123456"
	stockUser.Database = db

	session, err := mgo.DialWithInfo(stockUser)
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

func connectToStockTranCollection(coll string) *mgo.Collection {

	mgoCollection := mgoSession.DB("").C(coll)
	return mgoCollection

}
