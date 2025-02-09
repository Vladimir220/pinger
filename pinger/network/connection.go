package network

import "log"

type connection struct {
	url          string
	errorsLog    *log.Logger
	netErrorsLog *log.Logger
}

func CreateConnection(url string, errorsLog, netErrorsLog *log.Logger) connection {
	return connection{url: url, errorsLog: errorsLog, netErrorsLog: netErrorsLog}
}
