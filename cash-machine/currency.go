package cash_machine

type (
	Currency struct {
		Code     string
		Decimals float64
		Units    []NoteCoin
	}

	NoteCoin struct {
		Name  string
		Value int
	}
)

func InitializeGPB() Currency {
	nc := []NoteCoin{
		{Name: "£50", Value: 5000},
		{Name: "£20", Value: 2000},
		{Name: "£10", Value: 1000},
		{Name: "£5", Value: 500},
		{Name: "£2", Value: 200},
		{Name: "£1", Value: 100},
		{Name: "50p", Value: 50},
		{Name: "20p", Value: 20},
		{Name: "10p", Value: 10},
		{Name: "5p", Value: 5},
		{Name: "2p", Value: 2},
		{Name: "1p", Value: 1},
	}

	return Currency{
		Code:     "GBP",
		Decimals: 2,
		Units:    nc,
	}
}
