package handlers

import (
	"fmt"
	"net/http"
)

func (h Handlers) CreateTracking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Маршрут поддерживает только POST", http.StatusMethodNotAllowed)
		return
	}

	hostIp, err := h.readBodyWithIp(r)

	if err != nil {
		http.Error(w, "Ожидается поле ip", http.StatusBadRequest)
		return
	}

	err = h.daoDB.CreateTracking(hostIp.Ip)
	if err != nil {
		h.dbErrorsLog.Println(err)
		http.Error(w, "БД не может обработать запрос", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s теперь отслеживается\n", hostIp.Ip)
}
