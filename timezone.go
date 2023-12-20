/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/11/16
   Description :
-------------------------------------------------
*/

package zutils

import (
	"time"

	"gopkg.in/ugjka/go-tz.v2/tz"
)

var TZ = &tzUtil{
	CSTTimeZone: time.FixedZone("CST", 8*3600),
	UTCTimeZone: time.UTC,
}

type tzUtil struct {
	CSTTimeZone *time.Location // cst时区 / 中国时区
	UTCTimeZone *time.Location // utc时区
}

// 根据经纬度获取时区(经度, 纬度, 失败时默认时区=UTC)
func (*tzUtil) GetTimezoneOfGeo(lon, lat float64, def ...*time.Location) *time.Location {
	zone, _ := tz.GetZone(tz.Point{Lon: lon, Lat: lat})
	if len(zone) > 0 {
		loc, _ := time.LoadLocation(zone[0])
		if loc != nil {
			return loc
		}
	}

	if len(def) > 0 {
		return def[0]
	}
	return time.UTC
}

// 获取时区x相对于UTC的时差(时区, 当前时间), 单位小时
func (*tzUtil) GetTimezoneDiff(loc *time.Location, t ...time.Time) float32 {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
		now.In(time.UTC) // 不能写在一行, 否则可能导致入参的数据被改变
	} else {
		now = time.Now().In(time.UTC)
	}
	tLoc := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, loc)
	return float32(now.Unix()-tLoc.Unix()) / 3600 // 时间戳和时间文本是相反的
}

// 获取时区a相对于时区b的时差, 以时区b为准(时区a, 时区b, 当前时间), 单位小时
func (*tzUtil) GetTimezoneDiffOfZone(a, b *time.Location, t ...time.Time) float32 {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
		now.In(time.UTC) // 不能写在一行, 否则可能导致入参的数据被改变
	} else {
		now = time.Now().In(time.UTC)
	}
	ta := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, a)
	tb := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, b)
	return float32(tb.Unix()-ta.Unix()) / 3600 // 时间戳和时间文本是相反的
}
