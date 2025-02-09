package DAO

import (
	"server/models"
	"time"
)

type Dao interface {
	// Создаёт новое отслеживание состояния хоста
	CreateTracking(ip string) (err error)

	// Запрашивает статусы отслеживаемых хостов
	ReadAllStatuses() (hosts []models.HostStatus, err error)

	// Запрашивает список отслеживаемых хостов
	ReadIpList() (ips []models.HostIp, err error)

	// Обновляет информацию о состоянии отслеживаемого хоста
	UpdateStatus(ip string, pingTimeMs int, lastSuccessDate time.Time) (err error)

	// Удаляет отслеживание состояние хоста
	DeleteTracking(ip string) (err error)
}
