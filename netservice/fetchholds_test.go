package netservice

import (
	"log"
	"testing"
)

func TestMakeStockMoreInfoArray(t *testing.T) {

	log.Println("testing starting ...")
	MakeStockMoreInfoArray("601668")
	t.Fail()

}
