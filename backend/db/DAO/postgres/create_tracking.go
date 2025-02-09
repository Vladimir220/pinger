package postgres

import (
	"errors"
	"fmt"
	"server/db/DAO"
)

// Создаёт новое отслеживание состояния хоста
func (p DaoPostgres) CreateTracking(ip string) (err error) {
	if !DAO.IsValidIP(ip) {
		return errors.New("в CreateTracking передан некорректный IP")
	}

	queryStr := `INSERT INTO host_statuses (ip_address) VALUES ($1);`

	_, err = p.db.Exec(queryStr, ip)
	if err != nil {
		err = fmt.Errorf("ошибка при создании нового отслеживания: %v", err)
		return
	}

	return
}
