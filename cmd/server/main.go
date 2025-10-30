package main

import (
	"log"
	"net/http"
	"status-dashboard/internal/db"
	
)

func main() {
	dsn := "file:status.db?_fk=1&_busy_timeout=500"

	conn, err := db.InitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.RunMigrations(conn); err != nil {
		log.Fatal(err)
	}
	log.Println("Migraciones completadas")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`{"status":"ok"}`))
})

log.Println("🚀 Servidor escuchando en :8080")
http.ListenAndServe(":8080", nil)

}
