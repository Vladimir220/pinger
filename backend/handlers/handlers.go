package handlers

import (
	"log"
	"server/db/DAO"
)

// Содержит обработчики для маршрутов.
type Handlers struct {
	daoDB       DAO.Dao
	errorsLog   *log.Logger
	dbErrorsLog *log.Logger
	netLogFile  *log.Logger
}

// Создаёт и инициализирует обработчики для маршрутов.
func CreateHandlers(daoDB DAO.Dao, errorsLog, dbErrorsLog, netLogFile *log.Logger) (h Handlers) {
	h = Handlers{daoDB: daoDB, errorsLog: errorsLog, dbErrorsLog: dbErrorsLog}
	return
}
