package goutil

import (
	"fmt"
	"time"
)

// 时间戳转为日期
func TimeToDate(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02")
}

// 时间戳转为datetime格式
func TimeToDateTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 日期转为时间戳
func DateTimeToTime(date string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation("2006-01-02 15:04:05", date, LOC)
	return tim.Unix()
}

func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取一天开始的时间戳
func GetTodayStartTime() int64 {
	now := time.Now().Unix()
	return now - (now+28800)%86400
}

// 获取一周开始的时间戳
func GetWeekStartTime() int64 {
	now := time.Now()
	week := int(now.Weekday())
	if week == 0 {
		week = 7
	}
	t := now.AddDate(0, 0, 1-week).Unix()
	return t - (t+28800)%86400
}

// 获取一月开始的时间戳
func GetMonthStartTime() int64 {
	now := time.Now().AddDate(0, 0, 1-time.Now().Day()).Unix()
	return now - (now+28800)%86400
}

func TimeToWeek(t int64) string {
	week := time.Unix(t, 0)
	switch week.Weekday() {
	case time.Monday:
		return "星期一"
	case time.Tuesday:
		return "星期二"
	case time.Wednesday:
		return "星期三"
	case time.Thursday:
		return "星期四"
	case time.Friday:
		return "星期五"
	case time.Saturday:
		return "星期六"
	case time.Sunday:
		return "星期天"
	}
	return ""
}

// 时间戳转时间提示
func TimeToTips(t int64) string {
	now := time.Now().Unix()
	span := now - t

	result := ""
	if span > 0 {
		if span < 60 {
			result = "刚刚"
		} else if span <= 1800 {
			result = fmt.Sprintf("%d分钟前", span/60)
		} else if span <= 3600 {
			result = "半小时前"
		} else if span <= 86400 {
			result = fmt.Sprintf("%d小时前", span/3600)
		} else if span <= 86400*30 {
			result = fmt.Sprintf("%d天前", span/(86400))
		} else if span <= 86400*30*12 {
			result = fmt.Sprintf("%d月前", span/(86400*30))
		} else {
			result = fmt.Sprintf("%d年前", span/(86400*30*12))
		}
	} else {
		span = 0 - span
		if span < 60 {
			result = fmt.Sprintf("%d秒后", span)
		} else if span <= 1800 {
			result = fmt.Sprintf("%d分钟后", span/60)
		} else if span <= 3600 {
			result = "半小时后"
		} else if span <= 86400 {
			result = fmt.Sprintf("%d小时后", span/3600)
		} else if span <= 86400*30 {
			result = fmt.Sprintf("%d天后", span/(86400))
		} else if span <= 86400*30*12 {
			result = fmt.Sprintf("%d月后", span/(86400*30))
		} else {
			result = fmt.Sprintf("%d年后", span/(86400*30*12))
		}
	}

	return result
}
