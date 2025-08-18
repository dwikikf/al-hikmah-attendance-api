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

	_ = db

}
