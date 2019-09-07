package main

import (
	"fmt"
)

const (
	EVENTS_KEY = "EVENTS"
)

func (r *Redisful) initEvents() {
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		fmt.Println("ERROR INIT EVENTS: ", err)
	}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.PublicFg, &e.ClosedFg, &e.Price); err != nil {
			fmt.Println("ERROR INIT EVENTS: ", err)
		}
		err = r.SetHashToCache(EVENTS_KEY, e.ID, e)
		if err != nil {
			fmt.Println("ERROR INIT EVENTS: ", err)
		}
	}
}

func (r *Redisful) updateEventInCache(e Event) {
	err := r.SetHashToCache(EVENTS_KEY, e.ID, e)
	if err != nil {
		fmt.Println("ERROR UPDATE EVENT: ", err)
	}
}

func (r *Redisful) addEventInCache(e Event) {
	err := r.SetHashToCache(EVENTS_KEY, e.ID, e)
	if err != nil {
		fmt.Println("ERROR UPDATE EVENT: ", err)
	}
}
