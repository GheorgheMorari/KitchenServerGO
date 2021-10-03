package main

import "sync"

type Apparatus struct {
	busy int32
	meal *Meal
	mx sync.Mutex
}
