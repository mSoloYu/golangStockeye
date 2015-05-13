package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

func makeQueryModelForOpenSaleDate(stockcode string) bson.M {

	queryStat := bson.M{"code": stockcode}
	return queryStat

}
