package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type KitchenHandler struct{}

var packetsReceived int32= 0
var postReceived int32 = 0
var latestOrder Order
var latestOrderString string ="No order"

func (KitchenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Received something")
	atomic.StoreInt32(&packetsReceived, packetsReceived+1)
	if r.Method == http.MethodPost {
		atomic.StoreInt32(&postReceived, postReceived+1)
		//TODO make buffer static
		var buffer = make([]byte, r.ContentLength)
		r.Body.Read(buffer)
		var currentOrder Order
		json.Unmarshal(buffer,&currentOrder)
		if latestOrder.Id < currentOrder.Id {
			latestOrder = currentOrder
			latestOrderString = string(buffer)
		}
		//parseDiningHallRequest(buffer)
		fmt.Fprintln(w, "Kitchen http request post method detected.")
		fmt.Fprintln(w, "Kitchen request detected.\nPost Method Body:\n"+string(buffer))
	} else {
		fmt.Fprintln(w, "Kitchen server is UP on port "+kitchenServerPort)
		fmt.Fprintf(w, "Recieved %d requests. Post requests: %d\n", packetsReceived, postReceived)
		fmt.Fprintln(w,"Latest order:"+latestOrderString)
	}
}

//func parseDiningHallRequest(buffer []byte) map[string]string {
//	decoder := json.Decoder{}
//	result := decoder.Decode(string(buffer))
//
//	//TODO return a order type
//	return result
//}

const diningHallPort = ":7500"
const kitchenServerPort = ":8000"

func main() {
	var kitchenServer http.Server
	kitchenServer.Addr = kitchenServerPort
	kitchenServer.Handler = KitchenHandler{}

	fmt.Println(time.Now())
	fmt.Println("Kitchen is listening and serving on port:" + kitchenServerPort)
	if err := kitchenServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
