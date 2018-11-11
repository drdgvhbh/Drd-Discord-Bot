package stock

import "fmt"

type StockDoesNotExist struct {
	name string
}

func CreateStockDoesNotExistError(
	name string,
) *StockDoesNotExist {
	return &StockDoesNotExist{
		name,
	}
}

func (stockDoesNotExist StockDoesNotExist) Error() string {
	return fmt.Sprintf(
		"Cannot find stock matching name %s",
		stockDoesNotExist.name)
}
