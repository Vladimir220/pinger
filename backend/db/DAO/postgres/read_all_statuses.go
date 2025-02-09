package postgres

import (
	"fmt"
	"server/models"
)

// Запрашивает статусы отслеживаемых хостов
func (p DaoPostgres) ReadAllStatuses() (hosts []models.HostStatus, err error) {

	queryStr := `SELECT 
				ip_address,
				COALESCE(ping_time_ms, '0') AS ping_time_ms,
				COALESCE(last_success_date, '1970-01-01 00:00:00') AS last_success_date
				FROM host_statuses;`

	rows, err := p.db.Query(queryStr)
	if err != nil {
		err = fmt.Errorf("ошибка при чтении данных хостов: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		host := models.HostStatus{}
		err = rows.Scan(&host.Ip, &host.PingTimeMs, &host.LastSuccessDate)
		if err != nil {
			err = fmt.Errorf("ошибка чтения строк данных: %v", err)
			return
		}
		hosts = append(hosts, host)
	}

	return
}
