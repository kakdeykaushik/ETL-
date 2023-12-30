package main

import (
	"log"
	"math"
	"time"
)

func fatalErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s - %s", msg, err)
	}
}

func panicErr(err error, msg string) {
	if err != nil {
		log.Panicf("%s - %s", msg, err)
	}
}

func tsToEpoch(dt string) int64 {
	layout := "2006-01-02T15:04:05" // prof will shout
	parsedTime, err := time.Parse(layout, dt)
	panicErr(err, "Error parsing time:")
	return parsedTime.UnixMilli()
}

func round(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value*multiplier) / multiplier
}
