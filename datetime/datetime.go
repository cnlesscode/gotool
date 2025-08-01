package datetime

import (
	"time"
)

type DateTime struct {
	Year         int
	YearString   string
	Month        int
	MonthString  string
	Day          int
	DayString    string
	Hour         int
	HourString   string
	Minute       int
	MinuteString string
	Second       int
	SecondString string
	TimeStamp    int64
	Format       string
	Result       string
	Time         time.Time
}

func New() *DateTime {
	return &DateTime{
		Format: "2006-01-02 15:04:05",
	}
}

func (dt *DateTime) InitFromTimeStamp(timeStamp int64) {
	dt.TimeStamp = timeStamp
	dt.Time = time.Unix(timeStamp, 0)
	dt.initFromTimeStampBase()
}

func (dt *DateTime) initFromTimeStampBase() {
	dt.Year = dt.Time.Year()
	dt.YearString = dt.Time.Format("2006")
	dt.Month = int(dt.Time.Month())
	dt.MonthString = dt.Time.Format("01")
	dt.Day = dt.Time.Day()
	dt.DayString = dt.Time.Format("02")
	dt.Hour = dt.Time.Hour()
	dt.HourString = dt.Time.Format("15")
	dt.Minute = dt.Time.Minute()
	dt.MinuteString = dt.Time.Format("04")
	dt.Second = dt.Time.Second()
	dt.SecondString = dt.Time.Format("05")
	dt.Result = dt.Time.Format(dt.Format)
	dt.TimeStamp = dt.Time.Unix()
}

func (dt *DateTime) InitFromDatatime(datetime string) {
	dt.Time, _ = time.ParseInLocation(dt.Format, datetime, time.Local)
	dt.initFromTimeStampBase()
}

// 当前时间
func (dt *DateTime) Now() {
	dt.InitFromTimeStamp(time.Now().Unix())
}

// 切换时间
func (dt *DateTime) Swicth(tp string, val int) *DateTime {
	dtNew := New()
	dtNew.InitFromTimeStamp(dt.TimeStamp)
	switch tp {
	case "year":
		dtNew.Time = dt.Time.AddDate(val, 0, 0)
	case "month":
		month := val + dt.Month
		dayMax := dt.CountDaysOfAMonth(month + 1)
		useDay := dt.Day
		if dayMax < useDay {
			useDay = dayMax
		}
		dtNew.Time = time.Date(dt.Year, time.Month(month), useDay, dt.Hour, dt.Minute, dt.Second, 0, time.Local)
	case "day":
		dtNew.Time = dt.Time.AddDate(0, 0, val)
	case "hour":
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Hour)
	case "minute":
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Minute)
	case "second":
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Second)
	}
	dtNew.initFromTimeStampBase()
	return dtNew
}

// 获取某月天数
func (st *DateTime) CountDaysOfAMonth(month int) int {
	return time.Date(st.Year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Day()
}

// 闰年
func (dt *DateTime) IsLeap() bool {
	return dt.Year%4 == 0 && (dt.Year%100 != 0 || dt.Year%400 == 0)
}

func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
