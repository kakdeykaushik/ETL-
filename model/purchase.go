package model

import (
	"encoding/json"
	"etl/db"
	"etl/utils"
)

type Purchase struct {
	Base
	ItemID string  `json:"item_id"`
	Amount float32 `json:"amount"`
}

func (p *Purchase) SaveToDB() {
	if err := db.GetSession().Query("INSERT INTO purchase_events(id, event_type, user_id, epoch, item_id, amount) VALUES (now(), ?, ?, ?, ?, ?)", p.EventType, p.UserID, p.Epoch, p.ItemID, p.Amount).Exec(); err != nil {
		utils.PanicErr(err, "Error while inserting to DB")
	}
}

func (p *Purchase) UnmarshalJSON(data []byte) error {
	type alias Purchase // alias is important else it will go in inf loop
	var purchase alias

	if err := json.Unmarshal(data, &purchase); err != nil {
		return err
	}

	purchase.Amount = float32(utils.Round(float64(purchase.Amount), 2))
	purchase.Epoch = utils.TsToEpoch(purchase.Timestamp)
	*p = Purchase(purchase)
	return nil
}
