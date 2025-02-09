package handlers

import (
	"encoding/json"
	"net/http"
)

func (h Handlers) ReadAllStatuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Маршрут поддерживает только GET", http.StatusMethodNotAllowed)
		return
	}

	hosts, err := h.daoDB.ReadAllStatuses()
	if err != nil {
		h.dbErrorsLog.Println(err)
		http.Error(w, "БД не может обработать запрос", http.StatusInternalServerError)
		return
	}

	hostsJson, err := json.Marshal(hosts)
	if err != nil {
		h.dbErrorsLog.Println(err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(hostsJson)
}
