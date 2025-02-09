package models

import "time"

type HostStatus struct {
	Ip              string    `json:"ip"`
	PingTimeMs      int       `json:"ping_time_ms"`
	LastSuccessDate time.Time `json:"last_success_date"`
}

type HostIp struct {
	Ip string `json:"ip"`
}

type PageRange struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
