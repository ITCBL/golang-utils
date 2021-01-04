package date

import (
	"strconv"
	"time"
)

// 获取每月的开始日期和结束日期
func GetMonthStartAndEnd(myYear string, myMonth string) map[string]string {
	// 数字月份必须前置补零
	if len(myMonth) == 1 {
		myMonth = "0" + myMonth
	}
	yInt, _ := strconv.Atoi(myYear)

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, myYear+"-"+myMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	t1 := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	t2 := time.Date(yInt, newMonth+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	result := map[string]string{"start": t1, "end": t2}
	return result
}
