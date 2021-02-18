package config

import "time"

func BookingStartTime() time.Time {
	t, _ := time.Parse(StartTimeFormat, configuration().GoLive.StartIsoDatetime)
	return t
}

func ServerAddr() string {
	return ":" + configuration().Server.Port
}

func BookingCode() string {
	return configuration().GoLive.BookingCode
}

func GoLiveTime() time.Time {
	start, _ := time.Parse(StartTimeFormat, configuration().GoLive.StartIsoDatetime)
	return start
}