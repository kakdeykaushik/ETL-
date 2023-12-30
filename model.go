package main

import (
	"encoding/json"
)

type Event interface {
	SaveToDB()
	UnmarshalJSON(data []byte) error
}

type Base struct {
	EventType string `json:"event_type"`
	UserID    string `json:"user_id"`
	Timestamp string `json:"timestamp"`
	Epoch     int64
}

type LevelUp struct {
	Base
	Level int `json:"level"`
}

func (lup *LevelUp) SaveToDB() {
	if err := Session.Query("INSERT INTO levelup_events(id, event_type, user_id, epoch, level) VALUES	(now(), ?, ?, ?, ?)", lup.EventType, lup.UserID, lup.Epoch, lup.Level).Exec(); err != nil {
		panicErr(err, "Error while inserting to DB")
	}
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
	Amount float32 `json:"amount"`
}

func (p *Purchase) SaveToDB() {
	if err := Session.Query("INSERT INTO purchase_events(id, event_type, user_id, epoch, item_id, amount) VALUES (now(), ?, ?, ?, ?, ?)", p.EventType, p.UserID, p.Epoch, p.ItemID, p.Amount).Exec(); err != nil {
		panicErr(err, "Error while inserting to DB")
	}
}

func (p *Purchase) UnmarshalJSON(data []byte) error {
	type alias Purchase // alias is important else it will go in inf loop
	var purchase alias

	if err := json.Unmarshal(data, &purchase); err != nil {
		return err
	}

	purchase.Amount = float32(round(float64(purchase.Amount), 2))
	purchase.Epoch = tsToEpoch(purchase.Timestamp)
	*p = Purchase(purchase)
	return nil
}

type Login struct {
	Base
	DeviceType string `json:"device_type"`
}

func (l *Login) SaveToDB() {
	if err := Session.Query("INSERT INTO login_events(id, event_type, user_id, epoch, device) VALUES (now(), ?, ?, ?, ?)", l.EventType, l.UserID, l.Epoch, l.DeviceType).Exec(); err != nil {
		panicErr(err, "Error while inserting to DB")
	}
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
