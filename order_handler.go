package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

type OrderHandler struct {
	packetsReceived int32
	postReceived      int32
	latestOrder       PostOrder
	latestOrderString string
}


func (oh *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Received something")
	atomic.StoreInt32(&oh.packetsReceived, oh.packetsReceived+1)
	if r.Method == http.MethodPost {
		atomic.StoreInt32(&oh.postReceived, oh.postReceived+1)
		//TODO make buffer static
		var buffer = make([]byte, r.ContentLength)
		r.Body.Read(buffer)
		var currentOrder PostOrder
		json.Unmarshal(buffer, &currentOrder)
		if oh.latestOrder.Id < currentOrder.Id {
			oh.latestOrder = currentOrder
			oh.latestOrderString = string(buffer)
		}
		//parseDiningHallRequest(buffer)
		fmt.Fprintln(w, "Kitchen http request post method detected.")
		fmt.Fprintln(w, "Kitchen request detected.\nPost Method Body:\n"+string(buffer))
	} else {
		fmt.Fprintln(w, "Kitchen server is UP on port "+kitchenServerPort)
		fmt.Fprintf(w, "Recieved %d requests. Post requests: %d\n", oh.packetsReceived, oh.postReceived)
		fmt.Fprintln(w, "Latest order:"+oh.latestOrderString)
	}
}
