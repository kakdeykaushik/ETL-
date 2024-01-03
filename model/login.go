package model

import (
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

func (l *Login) Transform() error {
	epoch, err := utils.TsToEpoch(l.Timestamp)
	if err != nil {
		return err
	}
	l.Epoch = epoch
	return nil
}
