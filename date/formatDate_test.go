package date

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	now := time.Now()
	date := FormatDate(now, YYYY_MM_DD_HH_MM_SS_CN)
	fmt.Println("now -- ", now)
	fmt.Println("date -- ", date)
}
