package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DaoPostgres struct {
	db *sql.DB
}

func (dp *DaoPostgres) init() (err error) {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
		host     = os.Getenv("DB_HOST")
	)
	if user == "" || password == "" || dbName == "" || host == "" {
		err = errors.New("в env не указана одна из следующих переменных: DB_USER, DB_PASSWORD, DB_NAME, DB_HOST")
		return
	}

	loginInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, "postgres")

	db, err := sql.Open("postgres", loginInfo)
	if err != nil {
		err = fmt.Errorf("ошибка подключения к БД: %v", err)
		return
	}

	query := "SELECT EXISTS (SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)"

	var isDbExist bool
	err = db.QueryRow(query, dbName).Scan(&isDbExist)
	if err != nil {
		err = fmt.Errorf("ошибка при проверке существования базы данных: %v", err)
		return
	}

	if !isDbExist {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
		if err != nil {
			err = fmt.Errorf("ошибка при создании базы данных: %v", err)
			return
		}
	}

	loginInfo = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)

	dp.db, err = sql.Open("postgres", loginInfo)
	if err != nil {
		err = fmt.Errorf("ошибка подключения к БД: %v", err)
		return
	}

	return
}

func (dao DaoPostgres) checkMigrations() (err error) {
	driver, err := postgres.WithInstance(dao.db, &postgres.Config{})
	if err != nil {
		err = fmt.Errorf("ошибка создания драйвера: %v", err)
		return
	}

	tmod := os.Getenv("TEST_MOD")
	dbName := os.Getenv("DB_NAME")
	var path string
	if tmod != "1" {
		path = ""
	} else {
		currentDir, _ := os.Getwd()
		path = filepath.ToSlash(filepath.Dir(currentDir))
		if path != "" {
			path = path + "/"
		}
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+path+"db/migrations", dbName, driver)
	if err != nil {
		err = fmt.Errorf("ошибка создания мигратора: %v", err)
		return
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		err = fmt.Errorf("ошибка применения миграций: %v", err)
		return
	} else {
		err = nil
	}

	return
}

func (dao *DaoPostgres) Close() (err error) {
	err = dao.db.Close()
	if err != nil {
		err = fmt.Errorf("ошибка закрытия БД: %v", err)
	}
	return
}

func CreateDaoPostgres() (dao *DaoPostgres, err error) {
	psql := &DaoPostgres{}
	err = psql.init()
	if err != nil {
		return
	}

	err = psql.checkMigrations()
	if err != nil {
		return
	}

	return psql, err
}
