package main

import (
	"pinger/inits"
)

func main() {
	netErrorsLog, errorsLog := inits.InitSystem()

	server := CreateServer(errorsLog, netErrorsLog)
	server.StartPingServer()
}
