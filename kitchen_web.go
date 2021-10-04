package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type KitchenWeb struct {
	kitchenServer http.Server
	orderHandler  OrderHandler
	kitchenClient http.Client
}

func (kweb *KitchenWeb) start() {
	kweb.kitchenServer.Addr = kitchenServerPort
	kweb.kitchenServer.Handler = &kweb.orderHandler

	fmt.Println(time.Now())
	fmt.Println("Kitchen is listening and serving on port" + kitchenServerPort)
	if err := kweb.kitchenServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (kweb *KitchenWeb) deliver(delivery *Delivery) {

	requestBody, marshallErr := json.Marshal(delivery)
	if marshallErr != nil {
		fmt.Println("Marshalling error:", marshallErr)
	}

	request, newRequestError := http.NewRequest(http.MethodPost, diningHallHost+diningHallPort+"/delivery", bytes.NewBuffer(requestBody))
	if newRequestError != nil {
		fmt.Println("Could not create new request. Error:", newRequestError)
	} else {
		_, doError := kweb.kitchenClient.Do(request)
		if doError != nil {
			fmt.Println("ERROR Sending request. ERR:",doError)
			return
		}
	}
}
