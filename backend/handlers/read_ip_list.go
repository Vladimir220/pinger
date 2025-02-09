package handlers

import (
	"encoding/json"
	"net/http"
)

func (h Handlers) ReadIpList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Маршрут поддерживает только GET", http.StatusMethodNotAllowed)
		return
	}

	ips, err := h.daoDB.ReadIpList()
	if err != nil {
		h.dbErrorsLog.Println(err)
		http.Error(w, "БД не может обработать запрос", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	ipsJSON, err := json.Marshal(ips)
	if err != nil {
		h.errorsLog.Println(err)
		http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
		return
	}

	w.Write(ipsJSON)
}
