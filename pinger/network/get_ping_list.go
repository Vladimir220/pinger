package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pinger/models"
)

func (c connection) GetPingList() (ips []string, success bool) {
	resp, err := http.Get(fmt.Sprintf("%s/read_ip_list", c.url))
	if err != nil {
		c.netErrorsLog.Printf("не удалось получить доступ к ресурсу: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.netErrorsLog.Printf("информация не найдена: %v\n", err)
		return
	}

	var hosts []models.HostIp
	err = json.NewDecoder(resp.Body).Decode(&hosts)
	if err != nil {
		c.netErrorsLog.Printf("не удалось распоковать ответ ресурса обогащения: %v\n", err)
		return
	}

	for _, v := range hosts {
		ips = append(ips, v.Ip)
	}

	success = true
	return
}
