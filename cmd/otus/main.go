package main

import (
	"fmt"
	"net/http"
	"os"

	logrus "github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("HelloWorld!!")
	// logrus.Warning("Warning")
	// logrus.Error("Error")
	// logrus.Fatal("Fatal")
	// logrus.Debug("debug")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Не указан порт для работы веб-приложения")
	}
	logrus.Info(fmt.Sprintf("Приложение будет работать на %s порту", port))

	http.ListenAndServe(":"+port, nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	})
}
