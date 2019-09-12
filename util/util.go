package util

import (
	"time"
	"strconv"
)

func ParseInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}

func DecodeTime(timeValue time.Time) string {
	return timeValue.Format("2006-01-02T15:04:05")
}
