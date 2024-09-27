package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/datetime"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {

	dt := datetime.New()
	// 初始化当前时间
	dt.Now()
	Loger.Printf("当前时间 : %s", dt.Result)
	// 下一年
	nextYear := dt.Swicth("year", 1)
	Loger.Printf("一年后时间 : %s", nextYear.Result)
	// 下一月
	nextMonth := nextYear.Swicth("month", 1)
	Loger.Printf("一月后时间 : %s", nextMonth.Result)
	// 前一天
	lastDay := nextMonth.Swicth("day", -1)
	Loger.Printf("前一天时间 : %s", lastDay.Result)
	// 后3小时
	nextHour := lastDay.Swicth("hour", 3)
	Loger.Printf("后3小时时间 : %s", nextHour.Result)
	// 后5分钟
	nextMinute := nextHour.Swicth("minute", 5)
	Loger.Printf("后5分钟时间 : %s", nextMinute.Result)
	// 后10秒
	nextSecond := nextMinute.Swicth("second", 10)
	Loger.Printf("后10秒时间 : %s", nextSecond.Result)

	// 日期时间转换为时间戳
	timeStamp := datetime.DateTime2TimeStamp("01/02/2006 15:04:05", "10/26/2025 12:00:00")
	Loger.Printf("时间戳 : %d", timeStamp)

	// 遍历出2年24个月
	// 初始化起始时间
	startTime := datetime.New()
	startTime.InitFromTimeStamp(datetime.DateTime2TimeStamp("01/02/2006 15:04:05", "01/01/2024 00:00:00"))
	// 循环24个月
	months := make([]*datetime.DateTime, 0)
	for i := 0; i < 24; i++ {
		months = append(months, startTime.Swicth("month", i))
	}
	for _, v := range months {
		Loger.Printf("%s-%s", v.YearString, v.MonthString)
	}

	// 遍历出2个月内的天
	days := make([]*datetime.DateTime, 0)
	for i := 0; i < 2; i++ {
		// 当前月
		currentMonth := startTime.Swicth("month", i)
		dayCount := currentMonth.CountDaysOfAMonth()
		for j := 0; j < dayCount; j++ {
			days = append(days, currentMonth.Swicth("day", j))
		}
	}
	for _, v := range days {
		Loger.Printf("%s-%s-%s", v.YearString, v.MonthString, v.DayString)
	}

	// 遍历12小时
	startTimeForHour := datetime.New()
	startTimeForHour.Now()
	fmt.Printf("startTimeForHour: %v\n", startTimeForHour)
	startTimeForHour.InitFromDatatime(
		startTimeForHour.YearString + "-" +
			startTimeForHour.MonthString + "-" +
			startTimeForHour.DayString + " 00:00:00")
	hours := make([]*datetime.DateTime, 0)
	for i := 0; i < 12; i++ {
		hour := startTimeForHour.Swicth("hour", i)
		hours = append(hours, hour)
	}
	for _, hour := range hours {
		fmt.Printf("hour: %v\n", hour.Result)
	}
}
