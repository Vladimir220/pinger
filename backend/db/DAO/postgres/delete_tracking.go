package postgres

import (
	"errors"
	"fmt"
	"server/db/DAO"
)

// Удаляет отслеживание состояние хоста
func (p DaoPostgres) DeleteTracking(ip string) (err error) {
	if !DAO.IsValidIP(ip) {
		return errors.New("в DeleteTracking передан некорректный IP")
	}

	queryStr := `DELETE FROM host_statuses WHERE ip_address = $1;`

	_, err = p.db.Exec(queryStr, ip)
	if err != nil {
		err = fmt.Errorf("ошибка при удалении отслеживания: %v", err)
		return
	}

	return
}
