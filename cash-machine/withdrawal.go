package cash_machine

import (
	"fmt"
	"math"

	"code-katas/cash-machine/models"
)

func BreakIntoChange(amount float64) (string, error) {
	if amount < 0 {
		return "", fmt.Errorf("numbers less than zero are not allowed")
	}

	currGBP := InitializeGPB()

	change := models.ATMChange{}
	intValue := int(amount*math.Pow(10, currGBP.Decimals) + 0.01) // add 0.01 precision for float64 to int conversion

	for _, nc := range currGBP.Units {
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
