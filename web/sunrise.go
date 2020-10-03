package web

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// move this into another package
// middle of the day should be 0
const Julian2000 = 2451545

const testLat = -74.02

var tt = time.Date(2000, 1, 1, 12, 0, 0, 0, time.Local)

type Sun struct {
	Sunrise   string
	SolarNoon string
	Sunset    string
}

func Calculate(long, lat float64, year, month, day int) map[string]interface{} {
	jd := float64(calcJD(year, month, day))

	geomMeanLong := CalcGeomMeanLong(jd)
	meanAnom := CalcMeanAnom(jd)
	eqOfCenter := CalcEqOfCenter(meanAnom, jd)
	trueLong := CalcSunTrueLong(geomMeanLong, eqOfCenter)
	appLong := CalcSunAppLong(jd, trueLong)
	obliqCorr := CalcObliqCorr(jd)
	sunDeclin := CalcSunDeclin(appLong, obliqCorr)
	eccent := CalcEarthEccent(jd)
	eqOfTime := CalcEqOfTime(geomMeanLong, meanAnom, eccent, obliqCorr)
	haSunrise := CalcHASunrise(lat, sunDeclin)
	// hard coded
	offset := -4.0
	solarNoon := CalcSolarNoon(eqOfTime, long, offset)

	sunrise := CalcSunriseTime(solarNoon, haSunrise)
	sunset := CalcSunsetTime(solarNoon, haSunrise)

	return map[string]interface{}{
		"sunrise":   DecimalToTime(sunrise),
		"solarnoon": DecimalToTime(solarNoon),
		"sunset":    DecimalToTime(sunset),
	}
}

func calcJulianDate() float64 {
	since := time.Since(tt)
	n := since.Hours() / 24
	return Julian2000 + n
	//fmt.Println(since)
	//fmt.Println(since.Hours()
	//fmt.Println(since.Minutes())
}

func daysSince() int {
	since := time.Since(tt)
	n := int(math.Floor(since.Hours() / 24))
	return n
}

func degToRadians(n float64) float64 {
	return (math.Pi / 180) * n
}

func radToDegrees(n float64) float64 {
	return n * (180 / math.Pi)

}

func getJD(year, month, day float64) float64 {
	if month <= 2 {
		year -= 1
		month += 12
	}
	A := math.Floor(year / 100)
	B := 2 - A + math.Floor(A/4)
	return math.Floor(365.25*(year+4176)) + math.Floor(30.6001*(month+1)) + day + B - 1524.5
}

func calcJD(year, month, day int) int {
	return (1461*(year+4800+(month-14)/12))/4 + (367*(month-2-12*((month-14)/12)))/12 - (3*((year+4900+(month-14)/12)/100))/4 + day - 32075

}

func CalcMeanAnom(jd float64) float64 {
	// slightly rounded compared to the excels
	return 0.98560027751677*jd - 2415885.90323
}

func CalcGeomMeanLong(jd float64) float64 {
	return math.Mod((2.27273472596*math.Pow(10, -13))*math.Pow(jd, 2)+0.985646245822*jd-2416077.02518, 360)
}

func CalcEarthEccent(jd float64) float64 {
	return -9.4972127236*math.Pow(10, -17)*math.Pow(jd, 2) - 6.85253448057*math.Pow(10, -10)*jd + 0.018959353071
}

func CalcEarthEccent1(jd float64) float64 {
	return -9.4972127236 * math.Pow(10, -17) * math.Pow(jd, 2)
}

// This should be reworked because it is slightly off
func CalcEqOfCenter(meanAnom, jd float64) float64 {
	return 0.000289*math.Sin(3*degToRadians(meanAnom)) + math.Sin(2*degToRadians(meanAnom))*(0.026772084052-2.765229295*math.Pow(10, -9)) + math.Sin(degToRadians(meanAnom))*(-1.04941577056*math.Pow(10, -14)*jd*jd-8.04284727112*math.Pow(10, -8)*jd+2.17484667283)
}

func CalcSunTrueLong(geomMeanLong, sunEqOfCenter float64) float64 {
	return geomMeanLong + sunEqOfCenter
}

func CalcSunTrueAnom(geomMeanAnom, sunEqOfCenter float64) float64 {
	return geomMeanAnom + sunEqOfCenter
}

func CalcSunAppLong(jd, sunTrueLong float64) float64 {
	return sunTrueLong - 0.005569 - 0.00478*math.Sin(degToRadians((129943.559921-0.0529537577*jd)))
}

func CalcMeanObliqEcliptic(jd float64) float64 {
	return (23 + (26+(21.448-((jd-2451545)/36525)*(46.815+((jd-2451545)/36525)*(0.00059-((jd-2451545)/36525)*0.001813)))/60)/60)
}

func CalcObliqCorr(jd float64) float64 {
	return CalcMeanObliqEcliptic(jd) + 0.00256*math.Cos(degToRadians(125.04-1934.136*((jd-2451545)/36525)))
}

// saving this for last
func CalcSunRtAscent() float64 {
	// return radToDegrees(math.
	return 0
}

func CalcSunDeclin(sunAppLong, obliqCorr float64) float64 {
	return radToDegrees(math.Asin(math.Sin(degToRadians(obliqCorr)) * math.Sin(degToRadians(sunAppLong))))
}

// looks slighly off
func CalcEqOfTime(geomMeanLong, geomMeanAnom, eccentEarth, obliqCorr float64) float64 {
	y := CalcVarY(obliqCorr)
	return 4 * radToDegrees(y*math.Sin(2*degToRadians(geomMeanLong))-2*eccentEarth*math.Sin(degToRadians(geomMeanAnom))+4*eccentEarth*y*math.Sin(degToRadians(geomMeanAnom))*math.Cos(2*degToRadians(geomMeanLong))-0.5*y*y*math.Sin(4*degToRadians(geomMeanLong))-1.25*eccentEarth*eccentEarth*math.Sin(2*degToRadians(geomMeanAnom)))
}

func CalcVarY(obliqCorr float64) float64 {
	return math.Tan(degToRadians(obliqCorr/2)) * math.Tan(degToRadians(obliqCorr/2))
}

func CalcHASunrise(lat, sunDeclin float64) float64 {
	return radToDegrees(math.Acos(math.Cos(degToRadians(90.833))/(math.Cos(degToRadians(lat))*math.Cos(degToRadians(sunDeclin))) - math.Tan(degToRadians(lat))*math.Tan(degToRadians(sunDeclin))))
}

func CalcSolarNoon(eqOfTime, long, offset float64) float64 {
	return (720 - 4*long - eqOfTime + offset*60) / 1440
}

func CalcSunriseTime(solarNoon, haSunrise float64) float64 {
	return solarNoon - haSunrise*4/1440
}

func CalcSunsetTime(solarNoon, haSunrise float64) float64 {
	return solarNoon + haSunrise*4/1440
}

func DecimalToTime(t float64) string {
	b := strings.Builder{}

	hours := fmt.Sprint(math.Floor(t * 24))
	t = t*24 - math.Floor(t*24)

	minutes := fmt.Sprint(math.Floor(t * 60))
	t = t*60 - math.Floor(t*60)

	seconds := fmt.Sprint(math.Floor(t * 60))

	b.WriteString(hours)
	b.WriteString(":")
	b.WriteString(minutes)
	b.WriteString(":")
	b.WriteString(seconds)

	return b.String()

}

//func Calculate()
