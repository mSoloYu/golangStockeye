package mongodb

import (
	"testing"
)

func TestConnectToStockDb(t *testing.T) {
	connectToStockDb()
	t.Fail()
}

func TestConnectToStockCollection(t *testing.T) {
	connectToStockCollection()
	t.Fail()
}
