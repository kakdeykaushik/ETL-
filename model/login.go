package model

import (
	"encoding/json"
	"etl/db"
	"etl/utils"
	"fmt"
)

type Login struct {
	Base
	DeviceType string `json:"device_type"`
}

func (l *Login) SaveToDB() error {
	if err := db.GetSession().Query("INSERT INTO login_events(id, event_type, user_id, epoch, device) VALUES (now(), ?, ?, ?, ?)", l.EventType, l.UserID, l.Epoch, l.DeviceType).Exec(); err != nil {
		return fmt.Errorf("error while inserting to DB - %w", err)
	}
	return nil
}

func (l *Login) UnmarshalJSON(data []byte) error {
	type alias Login // alias is important else it will go in inf loop
	var login alias

	if err := json.Unmarshal(data, &login); err != nil {
		return err
	}

	login.Epoch = utils.TsToEpoch(login.Timestamp)
	*l = Login(login)
	return nil
}
