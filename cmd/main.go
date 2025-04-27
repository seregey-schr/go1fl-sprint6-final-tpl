package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	srv := server.NewServer(logger)

	if err := srv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
