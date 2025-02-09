package inits

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitSystem() (netErrorsLog *log.Logger, errorsLog *log.Logger) {

	err := godotenv.Load("./env/.env")
	if err != nil {
		panic(err.Error())
	}

	netErrorsLogFile, err := os.OpenFile(os.Getenv("NET_ERRORS_LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	errorsLogFile, err := os.OpenFile(os.Getenv("ERRORS_LOG_PATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}

	netErrorsLog = log.New(netErrorsLogFile, "NET: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	errorsLog = log.New(errorsLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	return
}
