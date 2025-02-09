package inits

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitSystem() (dbErrorsLog, errorsLog, netLog *log.Logger) {

	err := godotenv.Load("./env/.env")
	if err != nil {
		panic(err.Error())
	}

	dbErrorsLogFile, err := os.OpenFile(os.Getenv("DB_ERRORS_LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	errorsLogFile, err := os.OpenFile(os.Getenv("ERRORS_LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	netLogFile, err := os.OpenFile(os.Getenv("NET_LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	dbErrorsLog = log.New(dbErrorsLogFile, "DB_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	errorsLog = log.New(errorsLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	netLog = log.New(netLogFile, "NET_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	return
}
