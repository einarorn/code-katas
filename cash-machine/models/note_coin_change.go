package models

import (
	"fmt"
	"strings"
)

type NoteCoinChange struct {
	Amount string
	Count  int
}

type ATMChange []NoteCoinChange

func (a ATMChange) NotesAndCoinsToString() string {
	n := len(a) - 1
	b := strings.Builder{}
	b.WriteString("{")

	for i, change := range a {
		fmt.Fprintf(&b, "'%s': %d", change.Amount, change.Count)
		if i < n {
			fmt.Fprintf(&b, ", ")
		}
	}

	b.WriteString("}")

	return b.String()
}
