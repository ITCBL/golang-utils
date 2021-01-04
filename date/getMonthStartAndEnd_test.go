package date

import (
	"fmt"
	"testing"
)

func TestGetMonthStartAndEnd(t *testing.T) {
	monthStartAndEnd := GetMonthStartAndEnd("2020", "02")
	fmt.Println("monthStartAndEnd -- ", monthStartAndEnd)
}
