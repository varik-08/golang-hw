package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/varik-08/golang-hw/hw15_go_sql/config"
	"github.com/varik-08/golang-hw/hw15_go_sql/server/routes"
)

func main() {
	ctx := config.GetCTX()
	conf, _ := config.Init()

	config.InitDB(ctx, conf.DB)

	log.Println("Run server on", conf.HTTP.Host, ":", conf.HTTP.Port)

	mux := routes.SetupRoutes()

	server := &http.Server{
		Addr:         conf.HTTP.Host + ":" + strconv.Itoa(conf.HTTP.Port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
