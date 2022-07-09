package cash_machine

import (
	"fmt"

	"code-katas/cash-machine/models"
)

type noteCoinConstant struct {
	Name  string
	Value int
}

var noteCoinConstants = [...]noteCoinConstant{
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

func BreakIntoChange(amount float64) (string, error) {
	if amount < 0 {
		return "", fmt.Errorf("numbers less than zero are not allowed")
	}

	change := models.ATMChange{}
	intValue := int(amount * 100)

	for _, nc := range noteCoinConstants {
		count := intValue / nc.Value
		if count > 0 {
			change = append(change, models.NoteCoinChange{
				Amount: nc.Name,
				Count:  count,
			})
			intValue -= count * nc.Value
		}
	}

	return change.NotesAndCoinsToString(), nil
}
