package zutils

import (
	"time"
)

var T = &timeUtil{
	SecondStamp:   1e3,
	MinuteStamp:   60e3,
	HourStamp:     3600e3,
	DayStamp:      86400e3,
	Year2100Stamp: 4102416000000,

	Nanosecond:  time.Nanosecond,
	Microsecond: time.Microsecond,
	Millisecond: time.Millisecond,
	Second:      time.Second,
	Minute:      time.Minute,
	Hour:        time.Hour,
	Day:         time.Hour * 24,
	Week:        time.Hour * 24 * 7,
	Year365:     time.Hour * 24 * 365,
	Year366:     time.Hour * 24 * 366,

	Layout:           "2006-01-02 15:04:05",
	LayoutDate:       "2006-01-02",
	LayoutTime:       "15:04:05",
	LayoutTimeMinute: "15:04",
	LayoutDateMinute: "2006-01-02 15:04",
}

// 时间工具
//
// 这个工具内的所有时间戳均以毫秒为单位
type timeUtil struct {
	SecondStamp   int64 // 一秒的毫秒数
	MinuteStamp   int64 // 一分的毫秒数
	HourStamp     int64 // 一小时的毫秒数
	DayStamp      int64 // 一天的毫秒数
	Year2100Stamp int64 // 2100年的时间戳

	Nanosecond  time.Duration // 纳秒
	Microsecond time.Duration // 微秒
	Millisecond time.Duration // 毫秒
	Second      time.Duration // 秒
	Minute      time.Duration // 分
	Hour        time.Duration // 时
	Day         time.Duration // 天
	Week        time.Duration // 周
	Year365     time.Duration // 年
	Year366     time.Duration // 年

	Layout           string // 默认时间样式 YYYY-MM-DD hh:mm:ss
	LayoutDate       string // 日期样式 YYYY-MM-DD
	LayoutTime       string // 时间样式 hh:mm:ss
	LayoutTimeMinute string // 分钟样式 hh:mm
	LayoutDateMinute string // 带日期的分钟样式 YYYY-MM-DD hh:mm
}

func Time(loc *time.Location) *TimeUtil {
	return &TimeUtil{
		loc: loc,
	}
}

type TimeUtil struct {
	loc *time.Location
}

// region 基础
// 获取当前毫秒级时间戳
func (z *TimeUtil) GetStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前时间
func (z *TimeUtil) GetTime() time.Time {
	return time.Now().In(z.loc)
}

// 获取当前时间默认样式, YYYY-MM-DD hh:mm:ss
func (z *TimeUtil) GetText() string {
	return time.Now().In(z.loc).Format(T.Layout)
}

// 获取当前时间日期样式, YYYY-MM-DD
func (z *TimeUtil) GetDateText() string {
	return time.Now().In(z.loc).Format(T.LayoutDate)
}

// 获取当前时间时间样式, hh:mm:ss
func (z *TimeUtil) GetTimeText() string {
	return time.Now().In(z.loc).Format(T.LayoutTime)
}

// 获取当前时间分钟样式, hh:mm
func (z *TimeUtil) GetTimeMinuteText() string {
	return time.Now().In(z.loc).Format(T.LayoutTimeMinute)
}

// 获取当前时间带日期的分钟样式, YYYY-MM-DD hh:mm
func (z *TimeUtil) GetDateMinuteTextHour() string {
	return time.Now().In(z.loc).Format(T.LayoutDateMinute)
}

// endregion

// region 转换

// 将时间转为毫秒级时间戳
func (z *TimeUtil) TimeToStamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 将时间转为默认样式的字符串
func (z *TimeUtil) TimeToText(t time.Time) string {
	return t.In(z.loc).Format(T.Layout)
}

// 将时间转为日期样式的字符串
func (z *TimeUtil) TimeToDateText(t time.Time) string {
	return t.In(z.loc).Format(T.LayoutDate)
}

// 将时间转为指定样式的字符串
func (z *TimeUtil) TimeToTextOfLayout(t time.Time, layout string) string {
	return t.In(z.loc).Format(layout)
}

// 毫秒级时间戳转时间
func (z *TimeUtil) StampToTime(stamp int64) time.Time {
	return time.Unix(0, stamp*1e6).In(z.loc)
}

// 将毫秒级时间戳转为默认样式的字符串
func (z *TimeUtil) StampToText(stamp int64) string {
	return time.Unix(0, stamp*1e6).In(z.loc).Format(T.Layout)
}

// 将毫秒级时间戳转为日期样式的字符串
func (z *TimeUtil) StampToDateText(stamp int64) string {
	return time.Unix(0, stamp*1e6).In(z.loc).Format(T.LayoutDate)
}

// 将毫秒级时间戳转为指定样式的字符串
func (z *TimeUtil) StampToTextOfLayout(stamp int64, layout string) string {
	return time.Unix(0, stamp*1e6).In(z.loc).Format(layout)
}

// 将标准样式时间文本转为时间
func (z *TimeUtil) TextToTime(text string) (time.Time, error) {
	return time.ParseInLocation(T.Layout, text, z.loc)
}

// 将日期样式时间文本转为时间
func (z *TimeUtil) DateTextToTime(text string) (time.Time, error) {
	return time.ParseInLocation(T.LayoutDate, text, z.loc)
}

// 将指定样式时间文本转为时间
func (z *TimeUtil) TextToTimeOfLayout(text, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, text, z.loc)
}

// 将标准样式时间文本转为毫秒级时间戳
func (z *TimeUtil) TextToStamp(text string) (int64, error) {
	return z.TextToStampOfLayout(text, T.Layout)
}

// 将日期样式时间文本转为毫秒级时间戳
func (z *TimeUtil) DateTextToStamp(text string) (int64, error) {
	return z.TextToStampOfLayout(text, T.LayoutDate)
}

// 将时间文本转为毫秒级时间戳
func (z *TimeUtil) TextToStampOfLayout(text, layout string) (int64, error) {
	t, e := time.ParseInLocation(layout, text, z.loc)
	if e != nil {
		return 0, e
	}
	return t.UnixNano() / 1e6, nil
}

// endregion

// 获取当天开始时的毫秒级时间戳(0时0分0秒)
func (z *TimeUtil) GetDayStartTime() time.Time {
	t := time.Now().In(z.loc)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
}

// 获取当天开始时的毫秒级时间戳(0时0分0秒)
func (z *TimeUtil) GetDayStartStamp() int64 {
	t := time.Now().In(z.loc)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
	return t.Unix() * 1e3
}

// 获取传入时间戳当天的开始时间戳(0时0分0秒)
func (z *TimeUtil) GetDayStartTimeOfTime(t time.Time) time.Time {
	t = t.In(z.loc)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
}

// 获取传入时间戳当天的开始时间戳(0时0分0秒)
func (z *TimeUtil) GetDayStartStampOfStamp(stamp int64) int64 {
	t := time.Unix(0, stamp*1e6).In(z.loc)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
	return t.Unix() * 1e3
}

// GetWeekStartTimeOfWeek0 获取指定时间当周的起始时间(周日凌晨0点)
func (z *TimeUtil) GetWeekStartTimeOfWeek0(t time.Time) time.Time {
	t = t.In(z.loc)
	// 计算当前周的起始天（周日）
	weekday := int(t.Weekday())
	t = t.AddDate(0, 0, -weekday)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
	return t
}

// GetWeekStartTimeOfWeek1 获取指定时间当周的起始时间(周一凌晨0点)
func (z *TimeUtil) GetWeekStartTimeOfWeek1(t time.Time) time.Time {
	t = t.In(z.loc)
	// 计算当前周的起始天（周一）
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	offset := weekday - 1
	t = t.AddDate(0, 0, -offset)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, z.loc)
	return t
}
