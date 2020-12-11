/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/9
   Description :
-------------------------------------------------
*/

package zutils

import (
	"math"
	"strings"
)

var Geo = newGeo()

type geoUtil struct {
	municipalities map[string]struct{}
}

func newGeo() *geoUtil {
	return &geoUtil{
		municipalities: map[string]struct{}{
			"北京": {}, "上海": {}, "天津": {}, "重庆": {},
			"北京市": {}, "上海市": {}, "天津市": {}, "重庆市": {},
			"北京城区": {}, "上海城区": {}, "天津城区": {}, "重庆城区": {},
			"11": {}, "31": {}, "12": {}, "50": {}, // 代号
			"911": {}, "931": {}, "912": {}, "955": {}, // 代号
			"BJ": {}, "SH": {}, "TJ": {}, "CQ": {}, "bj": {}, "sh": {}, "tj": {}, "cq": {}, // 简写
		},
	}
}

// 计算两个坐标的距离, 输出单位:米
func (u *geoUtil) Distance(lon1, lat1, lon2, lat2 float64) float64 {
	if lon1 == 0 && lat1 == 0 {
		return 0
	}
	if lon2 == 0 && lat2 == 0 {
		return 0
	}

	radians := func(d float64) float64 {
		r := d * math.Pi / 180.0
		if d < 0 {
			r = -math.Abs(r)
		}
		return r
	}
	lon1, lat1, lon2, lat2 = radians(lon1), radians(lat1), radians(lon2), radians(lat2)
	dLon, dLat := lon2-lon1, lat2-lat1

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	return 2 * math.Asin(math.Sqrt(a)) * 6370996.81
}

// 中心点经纬度, 将每个经纬度转化成x,y,z的坐标值。然后根据根据x,y,z的值，寻找3D坐标系中的中心点
// GeoMidPoint({lon,lat},{lon,lat}) lon,lat
func (*geoUtil) MidPoint(points ...[]float64) (float64, float64) {
	if len(points) == 0 {
		return 0, 0
	}
	if len(points) == 1 {
		return points[0][0], points[0][1]
	}

	var x, y, z float64
	for _, point := range points {
		lon := point[0] * math.Pi / 180
		lat := point[1] * math.Pi / 180

		a := math.Cos(lat) * math.Cos(lon)
		b := math.Cos(lat) * math.Sin(lon)
		c := math.Sin(lat)

		x += a
		y += b
		z += c
	}

	pointsNum := float64(len(points))
	x /= pointsNum
	y /= pointsNum
	z /= pointsNum

	lon := math.Atan2(y, x)
	hyp := math.Sqrt(x*x + y*y)
	lat := math.Atan2(z, hyp)

	lon *= 180 / math.Pi
	lat *= 180 / math.Pi

	return lon, lat
}

// 中心点经纬度, 将经纬度坐标看成是平面坐标，直接计算经度和纬度的平均值
// GeoAveragePoint({lon,lat},{lon,lat}) lon,lat
//
// 该方法只是大致的估算方法，仅适合距离在400KM以内的点, 且极点附近不精确, 但是速度比MidPoint快
func (*geoUtil) MidPointPlane(points ...[]float64) (float64, float64) {
	if len(points) == 0 {
		return 0, 0
	}
	if len(points) == 1 {
		return points[0][0], points[0][1]
	}

	var x, y float64
	for _, point := range points {
		lon := point[0]
		lat := point[1]
		x += lon
		y += lat
	}

	pointsNum := float64(len(points))
	x /= pointsNum
	y /= pointsNum
	return x, y
}

// 是否为直辖市
func (u *geoUtil) IsMunicipalities(name string) bool {
	_, ok := u.municipalities[strings.ToUpper(name)]
	return ok
}

// 判断是否在中国, 只能简单判断, 不是很精确
func (u *geoUtil) GeoIsInsideChina(lon, lat float64) bool {
	InSideRectangle := [][]float64{
		// 左上，右下
		{79.446200, 49.220400, 96.330000, 42.889900},
		{109.687200, 54.141500, 135.000200, 39.374200},
		{73.124600, 42.889900, 124.143255, 29.529700},
		{82.968400, 29.529700, 97.035200, 26.718600},
		{97.025300, 29.529700, 124.367395, 20.414096},
		{107.975793, 20.414096, 111.744104, 17.871542},
	}
	OutSideRectangle := [][]float64{
		{119.921265, 25.398623, 122.497559, 21.785006},
		{101.865200, 22.284000, 106.665000, 20.098800},
		{106.452500, 21.542200, 108.051000, 20.487800},
		{109.032300, 55.817500, 119.127000, 50.325700},
		{127.456800, 55.817500, 137.022700, 49.557400},
		{131.266200, 44.892200, 137.022700, 42.569200},
	}
	for _, inRect := range InSideRectangle {
		if !u.InGeoRectangle(lon, lat, inRect) {
			continue
		}
		for _, outRect := range OutSideRectangle {
			if u.InGeoRectangle(lon, lat, outRect) {
				return false
			}
		}
		return true
	}
	return false
}

// 判断点是否在一个矩形内, rectangle{左上经度, 左上纬度, 右下经度, 右下纬度}
func (*geoUtil) InGeoRectangle(lon, lat float64, rectangle []float64) bool {
	if len(rectangle) != 4 {
		return false
	}
	return math.Min(rectangle[1], rectangle[3]) <= lat &&
		lat <= math.Max(rectangle[1], rectangle[3]) &&
		math.Min(rectangle[0], rectangle[2]) <= lon &&
		lon <= math.Max(rectangle[0], rectangle[2])
}
