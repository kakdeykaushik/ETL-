package model

import (
	"etl/db"
	"etl/utils"
	"fmt"
)

type LevelUp struct {
	Base
	Level int `json:"level"`
}

func (lup *LevelUp) SaveToDB() error {
	if err := db.GetSession().Query("INSERT INTO levelup_events(id, event_type, user_id, epoch, level) VALUES	(now(), ?, ?, ?, ?)", lup.EventType, lup.UserID, lup.Epoch, lup.Level).Exec(); err != nil {
		return fmt.Errorf("error while inserting to DB - %w", err)
	}
	return nil
}

func (lup *LevelUp) Transform() error {
	epoch, err := utils.TsToEpoch(lup.Timestamp)
	if err != nil {
		return err
	}
	lup.Epoch = epoch
	return nil
}
