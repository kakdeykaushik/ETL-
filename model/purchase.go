package model

import (
	"etl/db"
	"etl/utils"
	"fmt"
)

type Purchase struct {
	Base
	ItemID string  `json:"item_id"`
	Amount float32 `json:"amount"`
}

func (p *Purchase) SaveToDB() error {
	if err := db.GetSession().Query("INSERT INTO purchase_events(id, event_type, user_id, epoch, item_id, amount) VALUES (now(), ?, ?, ?, ?, ?)", p.EventType, p.UserID, p.Epoch, p.ItemID, p.Amount).Exec(); err != nil {
		return fmt.Errorf("error while inserting to DB - %w", err)
	}
	return nil
}

func (p *Purchase) Transform() error {
	epoch, err := utils.TsToEpoch(p.Timestamp)
	if err != nil {
		return err
	}
	p.Epoch = epoch
	p.Amount = float32(utils.Round(float64(p.Amount), 2))

	return nil
}
