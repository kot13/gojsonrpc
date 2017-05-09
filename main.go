package main

import (
	"gojsonrpc/config"
	"gojsonrpc/logger"
	"gojsonrpc/handlers"
	"gojsonrpc/jsonrpc/v1"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"net"
)

func main() {
	conf := config.GetConfig()
	logFinalizer, err := logger.InitLogger(conf.Logger)

	if err != nil {
		log.Fatal(err)
	}
	defer logFinalizer()
	log.Info("start")

	http.HandleFunc("/health-check", handlers.HealthCheck)
	http.HandleFunc("/api/v1", v1.HandleRequest)

	listener, err := net.Listen("tcp4", conf.App.AppPort)
	if err != nil {
		log.Fatal("Listen: ", err)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
