package main

import (
	"errors"
)

func getSheetByID(id int64) (*Sheet, error) {
	if 1 <= id && id <= 50 {
		return &Sheet{ID: id, Rank: "S", Num: id, Price: 5000}, nil
	} else if 51 <= id && id <= 200 {
		return &Sheet{ID: id, Rank: "A", Num: id - 50, Price: 3000}, nil
	} else if 201 <= id && id <= 500 {
		return &Sheet{ID: id, Rank: "B", Num: id - 200, Price: 1000}, nil
	} else if 501 <= id && id <= 1000 {
		return &Sheet{ID: id, Rank: "C", Num: id - 500, Price: 0}, nil
	}
	return nil, errors.New("No Column Error")
}
