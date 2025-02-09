package inits

import (
	"fmt"
	"log"
	"server/db/DAO"
	"server/handlers"

	"github.com/go-chi/chi/v5"
)

// Здесь происходит вся настройка роутера.
func InitRouter(dao DAO.Dao, errorsLog, dbErrorsLog, netLogFile *log.Logger) (router *chi.Mux) {
	router = chi.NewRouter()

	h := handlers.CreateHandlers(dao, errorsLog, dbErrorsLog, netLogFile)

	router.HandleFunc("/create_tracking", h.CreateTracking)
	router.HandleFunc("/delete_tracking", h.DeleteTracking)
	router.HandleFunc("/read_all_statuses", h.ReadAllStatuses)
	router.HandleFunc("/update_status", h.UpdateStatus)
	router.HandleFunc("/read_ip_list", h.ReadIpList)

	fmt.Printf("%-25s %-25s\n", "Маршрут", "Описание")
	fmt.Println("--------------------------------------")
	fmt.Printf("%-25s %-25s\n", "/create_tracking", "Создаёт новое отслеживание состояния хоста")
	fmt.Printf("%-25s %-25s\n", "/delete_tracking", "Удаляет отслеживание состояние хоста")
	fmt.Printf("%-25s %-25s\n", "/read_all_statuses", "Запрашивает статусы отслеживаемых хостов")
	fmt.Printf("%-25s %-25s\n", "/update_status", "Обновляет информацию о состоянии отслеживаемого хоста")
	fmt.Printf("%-25s %-25s\n", "/read_ip_list", "Запрашивает список отслеживаемых хостов")

	return
}
