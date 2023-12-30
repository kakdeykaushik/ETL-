package model

type Event interface {
	SaveToDB() error
	UnmarshalJSON(data []byte) error
}

type Base struct {
	EventType string `json:"event_type"`
	UserID    string `json:"user_id"`
	Timestamp string `json:"timestamp"`
	Epoch     int64
}
