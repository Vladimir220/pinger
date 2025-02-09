package postgres

import (
	"errors"
	"fmt"
	"server/db/DAO"
	"time"
)

// Обновляет информацию о состоянии отслеживаемого хоста
func (p DaoPostgres) UpdateStatus(ip string, pingTimeMs int, lastSuccessDate time.Time) (err error) {
	if !DAO.IsValidIP(ip) {
		return errors.New("в UpdateStatus передан некорректный IP")
	}

	queryStr := "UPDATE host_statuses SET ping_time_ms = $1, last_success_date = $2 WHERE ip_address = $3;"

	_, err = p.db.Exec(queryStr, pingTimeMs, lastSuccessDate.Format("2006-01-02 15:04:05"), ip)
	if err != nil {
		err = fmt.Errorf("ошибка при обновлении статуса: %v", err)
		return
	}

	return
}
