package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"time"
)

func parseDateTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func (h Handlers) readBodyWithIp(r *http.Request) (hostIp models.HostIp, err error) {
	contentType := r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		err = json.NewDecoder(r.Body).Decode(&hostIp)
		if err != nil {
			h.netLogFile.Printf("Ошибка кодирования в JSON: %v\n", err)
			return
		}
	case "application/x-www-form-urlencoded":
		err = r.ParseForm()
		if err != nil {
			h.netLogFile.Printf("Ошибка парсинга формы: %v\n", err)
			return
		}
		hostIp.Ip = r.FormValue("ip")
	}
	return
}
