package main

import (
	"fmt"
	"net/http"
	"server/db/DAO/postgres"
	"server/inits"
)

const Host = ":1234"

func main() {
	dbErrorsLog, errorsLog, netLogFile := inits.InitSystem()

	dao, err := postgres.CreateDaoPostgres()
	if err != nil {
		dbErrorsLog.Fatal(err)
	}
	defer dao.Close()

	router := inits.InitRouter(dao, errorsLog, dbErrorsLog, netLogFile)

	fmt.Printf("\nХост: %s/\n", Host)

	err = http.ListenAndServe(Host, router)
	if err != nil {
		errorsLog.Fatal(err)
	}
}
