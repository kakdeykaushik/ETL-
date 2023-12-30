package utils

import (
	"log"
	"math"
	"time"
)

func FatalErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s - %s", msg, err)
	}
}

func PanicErr(err error, msg string) {
	if err != nil {
		log.Panicf("%s - %s", msg, err)
	}
}

func TsToEpoch(dt string) int64 {
	layout := "2006-01-02T15:04:05" // prof will shout
	parsedTime, err := time.Parse(layout, dt)
	PanicErr(err, "Error parsing time:")
	return parsedTime.UnixMilli()
}

func Round(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value*multiplier) / multiplier
}
