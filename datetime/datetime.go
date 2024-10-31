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
	if dt.IsLeap() {
		daysOfMonth[1] = 29
	}
	dtNew := New()
	dtNew.InitFromTimeStamp(dt.TimeStamp)
	if tp == "year" {
		dtNew.Time = AddDate(dt.Time, val, 0)
	} else if tp == "month" {
		dtNew.Time = AddDate(dt.Time, 0, val)
	} else if tp == "day" {
		dtNew.Time = dt.Time.AddDate(0, 0, val)
	} else if tp == "hour" {
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Hour)
	} else if tp == "minute" {
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Minute)
	} else if tp == "second" {
		dtNew.Time = dt.Time.Add(time.Duration(val) * time.Second)
	}
	dtNew.initFromTimeStampBase()
	return dtNew
}

// 获取某月天数
func (st *DateTime) CountDaysOfAMonth() int {
	return time.Date(st.Year, st.Time.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// 闰年
func (dt *DateTime) IsLeap() bool {
	return dt.Year%4 == 0 && (dt.Year%100 != 0 || dt.Year%400 == 0)
}

var daysOfMonth = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func AddDate(t time.Time, years, months int) time.Time {
	month := t.Month()
	years, months = norm(years, months, 12)
	targetMonth := (int(month) + months) % 12
	targetYear := t.Year() + years + (int(month)+months)/12
	maxDayOfTargetMonth := daysOfMonth[targetMonth-1]
	if isLeap(targetYear) && targetMonth == 2 {
		maxDayOfTargetMonth += 1 // 闰年2月多一天
	}
	targetDay := t.Day()
	if targetDay > maxDayOfTargetMonth {
		targetDay = maxDayOfTargetMonth
	}
	return time.Date(targetYear, time.Month(targetMonth), targetDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
