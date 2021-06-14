package main

import (
	"fmt"
	router2 "gin-api/router"
	"gin-api/utils"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"time"
)

func main() {
	initialize()
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 30 * time.Second
	router := router2.Routers()
	server := endless.NewServer(":8080", router)
	server.BeforeBegin = func(add string) {
		log.Printf("pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func initialize() {
	utils.Config.LoadConfig()
	utils.Gorm.InitGorm()
}
