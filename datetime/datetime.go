package datetime

import (
	"strconv"
	"time"
)

// Time stamp to date time
func TimeStampToDatatime(timeStamp int64) string {
	tm := time.Unix(timeStamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// Time stamp to slice [] string
func TimeStampToDatatimeSlice(timeStamp int64) []string {
	tm := time.Unix(timeStamp, 0)
	tm.Format("2006-01-02 15:04:05")
	return []string{
		tm.Format("2006"),
		tm.Format("01"),
		tm.Format("02"),
		tm.Format("15"),
		tm.Format("04"),
		tm.Format("05"),
	}
}

// Date time to time stamp
func DateTimeToTimeStamp(datetime string) int64 {
	res, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	if err == nil {
		return res.Unix()
	} else {
		return 0
	}
}

// Calculate past time
func FormatPastTime(pastTimeStamp int64) (int, string, string) {
	currentTime := time.Now().Unix()
	timeDifference := int(currentTime - pastTimeStamp)
	if timeDifference < 180 {
		return timeDifference, "just", "刚刚"
	} else if timeDifference >= 180 && timeDifference < 3600 {
		resultInt := int(timeDifference / 60)
		return timeDifference, strconv.Itoa(resultInt) + " minutes ago", strconv.Itoa(resultInt) + "分钟前"
	} else if timeDifference >= 3600 && timeDifference < 86400 {
		resultInt := int(timeDifference / 3600)
		return timeDifference, strconv.Itoa(resultInt) + " hours ago", strconv.Itoa(resultInt) + "小时前"
	} else if timeDifference >= 86400 && timeDifference < 2592000 {
		resultInt := int(timeDifference / 86400)
		return timeDifference, strconv.Itoa(resultInt) + " days ago", strconv.Itoa(resultInt) + "天前"
	} else {
		return timeDifference, TimeStampToDatatime(pastTimeStamp), TimeStampToDatatime(pastTimeStamp)
	}
}

// Count Days Of A Month
func CountDaysOfAMonth(yearAndMonth string) (int, error) {
	if len([]rune(yearAndMonth)) < 6 {
		return 0, nil
	}
	year, erry := strconv.Atoi(yearAndMonth[0:4])
	if erry != nil {
		return 0, erry
	}
	monthStr := yearAndMonth[4:6]
	if monthStr[0:1] == "0" {
		monthStr = monthStr[1:2]
	}
	month, _ := strconv.Atoi(monthStr)
	var days int = 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return days, nil
}

// Is Leap Year
func IsLeapYear(year int) bool {
	if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
		return true
	} else {
		return false
	}
}
