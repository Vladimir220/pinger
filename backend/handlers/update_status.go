package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"strconv"
)

func (h Handlers) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Маршрут поддерживает только POST", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	var err error
	hostStatus := models.HostStatus{}

	switch contentType {
	case "application/json":
		err = json.NewDecoder(r.Body).Decode(&hostStatus)
		if err != nil {
			h.netLogFile.Printf("Ошибка кодирования в JSON: %v\n", err)
			break
		}
	case "application/x-www-form-urlencoded":
		err = r.ParseForm()
		if err != nil {
			h.netLogFile.Printf("Ошибка парсинга формы: %v\n", err)
			break
		}
		hostStatus.Ip = r.FormValue("ip")

		pingTime, err := strconv.Atoi(r.FormValue("ping_time_ms"))
		if err != nil {
			h.netLogFile.Printf("Ошибка конвертации: %v\n", err)
			break
		}

		hostStatus.PingTimeMs = pingTime

		lastSuccessDate, err := parseDateTime(r.FormValue("last_success_date"))
		if err != nil {
			h.netLogFile.Printf("Недопустимый формат даты: %v\n", err)
			break
		}

		hostStatus.LastSuccessDate = lastSuccessDate
	}

	if err != nil {
		http.Error(w, "Ожидаются поля ip, ping_time_ms, last_success_date", http.StatusBadRequest)
		return
	}

	err = h.daoDB.UpdateStatus(hostStatus.Ip, hostStatus.PingTimeMs, hostStatus.LastSuccessDate)
	if err != nil {
		h.dbErrorsLog.Println(err)
		http.Error(w, "БД не может обработать запрос", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Для %s обновлен статус: %d, %v\n", hostStatus.Ip, hostStatus.PingTimeMs, hostStatus.LastSuccessDate)
}
