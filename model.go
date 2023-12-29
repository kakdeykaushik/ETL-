package main

import (
	"encoding/json"
	"fmt"
)

type SomeInterface interface {
	SaveToDB()
	UnmarshalJSON(data []byte) error
}

type Base struct {
	EventType string `json:"event_type"`
	UserID    string `json:"user_id"`
	Timestamp string `json:"timestamp"`
	Epoch     int64
}

// todo
func (Base) SaveToDB() {
	fmt.Println("saved to db")
}

type LevelUp struct {
	Base
	Level int `json:"level"`
}

func (lup *LevelUp) UnmarshalJSON(data []byte) error {
	type alias LevelUp // alias is important else it will go in inf loop
	var levelUp alias

	if err := json.Unmarshal(data, &levelUp); err != nil {
		return err
	}

	levelUp.Epoch = tsToEpoch(levelUp.Timestamp)
	*lup = LevelUp(levelUp)
	return nil

}

type Purchase struct {
	Base
	ItemID string  `json:"item_id"`
	Amount float64 `json:"amount"`
}

func (p *Purchase) UnmarshalJSON(data []byte) error {
	type alias Purchase // alias is important else it will go in inf loop
	var purchase alias

	if err := json.Unmarshal(data, &purchase); err != nil {
		return err
	}

	purchase.Amount = round(purchase.Amount, 2)
	purchase.Epoch = tsToEpoch(purchase.Timestamp)
	*p = Purchase(purchase)
	return nil
}

type Login struct {
	Base
	DeviceType string `json:"device_type"`
}

func (l *Login) UnmarshalJSON(data []byte) error {
	type alias Login // alias is important else it will go in inf loop
	var login alias

	if err := json.Unmarshal(data, &login); err != nil {
		return err
	}

	login.Epoch = tsToEpoch(login.Timestamp)
	*l = Login(login)
	return nil
}
