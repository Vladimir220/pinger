package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pinger/models"
)

func (c connection) UpdatePingStatus(hostStatus models.HostStatus) (success bool) {
	data, err := json.Marshal(hostStatus)
	if err != nil {
		c.errorsLog.Printf("ошибка преобразования в Json: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/update_status", c.url), bytes.NewBuffer(data))
	if err != nil {
		c.errorsLog.Printf("ошибка создания запроса: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.netErrorsLog.Printf("ошибка при отправке запроса: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.netErrorsLog.Println("запрос не выполнен")
		return
	}

	success = true
	return
}
