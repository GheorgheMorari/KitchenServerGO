package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type KitchenWeb struct {
	kitchenServer http.Server
	orderHandler OrderHandler
}

func (kweb *KitchenWeb) start(){
	kweb.kitchenServer.Addr = kitchenServerPort
	kweb.kitchenServer.Handler = &kweb.orderHandler

	fmt.Println(time.Now())
	fmt.Println("Kitchen is listening and serving on port" + kitchenServerPort)
	if err := kweb.kitchenServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}