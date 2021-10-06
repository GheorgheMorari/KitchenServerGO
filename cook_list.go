package main

import "math/rand"

type CookList struct {
	cookList      []Cook
	cookIdCounter int
}

func (cl CookList) start() {
	cl.cookIdCounter = 0
	for i := 0; i < cookN; i++ {
		randomCook := cookPersonas[rand.Intn(len(cookPersonas))]
		randomCook.id = cl.cookIdCounter
		cl.cookIdCounter++
		if i == 0 {
			randomCook.rank = 3
		}
		cl.cookList = append(cl.cookList, randomCook)
	}

	for _, cook := range cl.cookList {
		go cook.startWorking()
	}
}

var cookPersonas = []Cook{{
	rank:        1,
	proficiency: 1,
	name:        "Jimmy Cook",
	catchPhrase: "YES!",
}, {
	rank:        2,
	proficiency: 2,
	name:        "Andy",
	catchPhrase: "Why am i here?",
}, {
	rank:        1,
	proficiency: 3,
	name:        "Karen",
	catchPhrase: "Abolish the patriarchy",
}, {
	rank:        3,
	proficiency: 2,
	name:        "Vanessa",
	catchPhrase: "The cake is a lie",
}, {
	rank:        3,
	proficiency: 3,
	name:        "Gordon Ramsay",
	catchPhrase: "WHERE IS THE LAMB SAUCE?",
}}
