package main

import (
	"log"

	"github.com/dwikikf/al-hikmah-attendance-api/internal/app"
)

func main() {
	config, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi : %v", err)
	}

	db, err := app.InitDB(config.Database)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database :%v", err)
	}

	router := app.NewRouter(db)

	log.Println("Server starting on port :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}

}
