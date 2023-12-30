package model

import (
	"encoding/json"
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

func (lup *LevelUp) UnmarshalJSON(data []byte) error {
	type alias LevelUp // alias is important else it will go in inf loop
	var levelUp alias

	if err := json.Unmarshal(data, &levelUp); err != nil {
		return err
	}

	levelUp.Epoch = utils.TsToEpoch(levelUp.Timestamp)
	*lup = LevelUp(levelUp)
	return nil

}
