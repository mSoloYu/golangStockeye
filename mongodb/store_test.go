package mongodb

import (
	"testing"
)

func TestStoreStockModel(t *testing.T) {

	StoreStockModel([]string{"601668"})
	t.Fail()

}
