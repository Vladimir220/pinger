package main

import (
	"log"
	"os"
	"pinger/models"
	"pinger/network"
	"pinger/ping"
	"strconv"
	"time"
)

const defaultHost = "http://:1234"

type Server struct {
	errorsLog    *log.Logger
	netErrorsLog *log.Logger
}

func (s Server) StartPingServer() {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = defaultHost
	}

	pingFreqStr := os.Getenv("PING_FREQUENCY_SEC")
	pingFreq, err := strconv.Atoi(pingFreqStr)
	if err != nil {
		s.errorsLog.Printf("ошибка преобразования строки в число: %v\n", err)
		return
	}

	conn := network.CreateConnection(host, s.errorsLog, s.netErrorsLog)

	//запускаем бесконечную проверку
	for {
		ips, success := conn.GetPingList()
		if !success {
			time.Sleep(10 * time.Second)
			continue
		}

		for _, ip := range ips {
			pingTime, err := ping.Ping(ip)
			if err != nil {
				s.netErrorsLog.Printf("Пинг на '%s' не прошел. Ошибка: %v", ip, err)
				time.Sleep(10 * time.Second)
				continue
			}

			hostStatus := models.HostStatus{
				Ip:              ip,
				PingTimeMs:      int(pingTime.Milliseconds()),
				LastSuccessDate: time.Now(),
			}

			_ = conn.UpdatePingStatus(hostStatus)
		}
		time.Sleep(time.Duration(pingFreq) * time.Second)

	}
}

func CreateServer(errorsLog, netErrorsLog *log.Logger) Server {
	return Server{errorsLog: errorsLog, netErrorsLog: netErrorsLog}
}
