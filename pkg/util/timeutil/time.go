package timeutil

import (
	"errors"
	"strconv"
	"time"
)

const DATE = `2006-01-02`
const TIME = `15:04:05`
const DATE_TIME = `2006-01-02 15:04:05`
const DATE_TIME_START_SECOND = `2006-01-02 15:04:00`
const DATE_TIME_ZERO_HOUR = `2006-01-02 00:00:00`

func ParseInLocal(layout, value string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layout, value, loc)
}

// 获取月份开始、结束
func GetMonthStartEnd(dd time.Time) (start time.Time, end time.Time) {
	year, month, _ := dd.Date()
	loc := dd.Location()

	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	return startOfMonth, endOfMonth
}

// CalculateTimeDelay delay = two - one 单位是s
func CalculateTimeDelay(one string, two string) (int64, error) {

	loc, ok := time.LoadLocation("Local")
	if ok != nil {
		return 0, ok
	}

	timeLayout := "2006-01-02 15:04:05"

	oneTime, ok1 := time.ParseInLocation(timeLayout, one, loc)
	if ok1 != nil {
		return 0, ok1
	}
	one1 := oneTime.Unix() //转化为时间戳 类型是int64

	twoTime, ok2 := time.ParseInLocation(timeLayout, two, loc)
	if ok2 != nil {
		return 0, ok2
	}
	two1 := twoTime.Unix()

	return two1 - one1, nil
}

// CalculateTimeDelayDay delay = two - one 单位是day
func CalculateTimeDelayDay(one1 string, two1 string) (int64, error) {

	loc, ok := time.LoadLocation("Local")
	if ok != nil {
		return 0, ok
	}

	if len(one1) != 19 || len(two1) != 19 {
		return -1, errors.New("inputs error")
	}

	one := one1[:11] + "00:00:00"
	two := two1[:11] + "00:00:00"

	timeLayout := "2006-01-02 15:04:05"

	oneTime, ok1 := time.ParseInLocation(timeLayout, one, loc)
	if ok1 != nil {
		return 0, ok1
	}
	twoTime, ok2 := time.ParseInLocation(timeLayout, two, loc)
	if ok2 != nil {
		return 0, ok2
	}
	delays := int64(timeSubDays(twoTime, oneTime))

	if delays == -1 {
		return -1, errors.New("sub days error")
	}
	return delays, nil
}

func timeSubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}

	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}

	if hours < 24 {
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		}
		return 1
	}

	if (hours/24)-float64(int(hours/24)) == 0 {
		return int(hours / 24)
	}

	return int(hours/24) + 1

}

// GetMonthStartAndEnd
func GetMonthStartAndEnd(myYear string, myMonth string) map[string]string {

	if len(myMonth) == 1 {
		myMonth = "0" + myMonth
	}
	yInt, _ := strconv.Atoi(myYear)

	loc, _ := time.LoadLocation("Local")

	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", myYear+"-"+myMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	t1 := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	t2 := time.Date(yInt, newMonth+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	result := map[string]string{"start": t1, "end": t2}
	return result
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}
