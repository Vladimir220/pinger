package postgres

import (
	"fmt"
	"server/models"
)

// Запрашивает список отслеживаемых хостов
func (p DaoPostgres) ReadIpList() (ips []models.HostIp, err error) {

	queryStr := `SELECT ip_address FROM host_statuses;`

	rows, err := p.db.Query(queryStr)
	if err != nil {
		err = fmt.Errorf("ошибка при чтении списка отслеживаемых хостов: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var host models.HostIp
		err = rows.Scan(&host.Ip)
		if err != nil {
			err = fmt.Errorf("ошибка чтения строк данных: %v", err)
			return
		}
		ips = append(ips, host)
	}

	return
}
