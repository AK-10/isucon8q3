package main

import ()

func getEvents(all bool) ([]*Event, error) {
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Title, &event.PublicFg, &event.ClosedFg, &event.Price); err != nil {
			return nil, err
		}
		if !all && !event.PublicFg {
			continue
		}
		e, err := getEventWithoutDetail(event)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func getEventWithoutDetail(event Event) (*Event, error) {
	(&event).initialize()

	rows, err := db.Query("SELECT * FROM reservations WHERE event_id = ? AND canceled_at IS NULL ", event.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Reservation
		if err := rows.Scan(&r.ID, &r.EventID, &r.SheetID, &r.UserID, &r.ReservedAt, &r.CanceledAt); err != nil {
			return nil, err
		}
		sheet, err := getSheetByID(r.SheetID)
		if err != nil {
			return nil, err
		}
		event.Remains--
		event.Sheets[sheet.Rank].Remains--
	}

	return &event, nil
}

func getEvent(eventID, loginUserID int64) (*Event, error) {
	var event Event
	if err := db.QueryRow("SELECT * FROM events WHERE id = ?", eventID).Scan(&event.ID, &event.Title, &event.PublicFg, &event.ClosedFg, &event.Price); err != nil {
		return nil, err
	}
	// initialize
	(&event).initialize()
	var i int64
	for i = 1; i <= 1000; i++ {
		sheet, _ := getSheetByID(i)
		// sheet.Mine = false
		event.Sheets[sheet.Rank].Detail = append(event.Sheets[sheet.Rank].Detail, sheet)
	}
	rows, err := db.Query("SELECT * FROM reservations WHERE event_id = ? AND canceled_at IS NULL ", event.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Reservation
		if err := rows.Scan(&r.ID, &r.EventID, &r.SheetID, &r.UserID, &r.ReservedAt, &r.CanceledAt); err != nil {
			return nil, err
		}
		sheet, err := getSheetByID(r.SheetID)
		if err != nil {
			return nil, err
		}
		event.Remains--
		event.Sheets[sheet.Rank].Remains--
		sheet.Mine = r.UserID == loginUserID
		sheet.Reserved = true
		sheet.ReservedAtUnix = r.ReservedAt.Unix()

		event.Sheets[sheet.Rank].Detail[sheet.ID-1] = sheet
	}

	return &event, nil
}

func (e *Event) initialize() {
	e.Total = 1000
	e.Remains = 1000
	e.Sheets = map[string]*Sheets{
		"S": &Sheets{Total: 50, Remains: 50, Price: e.Price + 5000},
		"A": &Sheets{Total: 150, Remains: 150, Price: e.Price + 3000},
		"B": &Sheets{Total: 300, Remains: 300, Price: e.Price + 1000},
		"C": &Sheets{Total: 500, Remains: 500, Price: e.Price + 0},
	}
}
