package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

const (
	EVENTS_KEY = "EVENTS"
)

func (r *Redisful) getEvents() ([]*Event, error) {
	data, err := r.GetAllHashFromCache(EVENTS_KEY)
	if err != nil {
		return nil, err
	}
	events := unmarshalEvents(data)
	return events, nil
}

func sortEvents(events []*Event) []*Event {
	sort.Slice(events, func(i, j int) bool { return events[i].ID < events[j].ID })
	return events
}

func unmarshalEvents(data [][]byte) []*Event {
	events := make([]*Event, 0, len(data))
	for i := range data {
		var event Event
		json.Unmarshal(data[i], &event)
		events = append(events, &event)
	}
	return events
}

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

func (e Event) toJson() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = e.ID
	res["title"] = e.Title
	res["public"] = e.PublicFg
	res["closed"] = e.ClosedFg
	res["price"] = e.Price
	return res
}

func (r *Redisful) updateEventInCache(e Event) {
	err := r.SetHashToCache(EVENTS_KEY, e.ID, e.toJson())
	if err != nil {
		fmt.Println("ERROR UPDATE EVENT: ", err)
	}
}

func (r *Redisful) addEventInCache(e Event) {
	err := r.SetHashToCache(EVENTS_KEY, e.ID, e.toJson())
	if err != nil {
		fmt.Println("ERROR UPDATE EVENT: ", err)
	}
}
