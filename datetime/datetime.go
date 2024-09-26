package datetime

import "time"

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
}

// 当前时间
func (dt *DateTime) Now() {
	dt.InitFromTimeStamp(time.Now().Unix())
}

// 切换年份
func (dt *DateTime) SwicthYear(year int) *DateTime {
	dtNew := New()
	dtNew.InitFromTimeStamp(dt.TimeStamp)
	dtNew.Time = dt.Time.AddDate(year, 0, 0)
	dtNew.initFromTimeStampBase()
	return dtNew
}

// 切换月份
func (dt *DateTime) SwicthMonth(month int) *DateTime {
	dtNew := New()
	dtNew.InitFromTimeStamp(dt.TimeStamp)
	dtNew.Time = dt.Time.AddDate(0, month, 0)
	dtNew.initFromTimeStampBase()
	return dtNew
}
