package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mux "github.com/gorilla/mux"

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

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	})

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	go serv.ListenAndServe()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	timeout, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	if err := serv.Shutdown(timeout); err != nil {
		logrus.Fatal("Не удалось корректно завершить работу веб-сервера")
	}
}
