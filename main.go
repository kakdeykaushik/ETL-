package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", recoverMiddleware(http.HandlerFunc(handleEvents)))
	log.Println("server is starting")
	fatalErr(http.ListenAndServe(":8000", nil), "Error while starting server")
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	panicErr(err, "Error reading request body")

	m := map[string]any{}
	err = json.Unmarshal(data, &m)
	panicErr(err, "Error while unmarshaling request body []byte")

	eventType, ok := m["event_type"].(string)
	if !ok {
		panicErr(fmt.Errorf("invalid event_type"), "Invalid event_type received from request body")
		// todo this should be returned to caller with 4xx error
	}
	event, err := concreteEvent(eventType, data)
	panicErr(err, "Error while getting concrete event")

	event.SaveToDB()

}

func concreteEvent(eventType string, data []byte) (Event, error) {
	if eventType == "login" {
		var login Login
		err := json.Unmarshal(data, &login)
		if err != nil {
			return nil, err
		}
		return &login, nil
	}

	if eventType == "purchase" {
		var purchase Purchase
		err := json.Unmarshal(data, &purchase)
		if err != nil {
			return nil, err
		}
		return &purchase, nil
	}

	if eventType == "level_up" {
		var levelUp LevelUp
		err := json.Unmarshal(data, &levelUp)
		if err != nil {
			return nil, err
		}
		return &levelUp, nil
	}

	return nil, fmt.Errorf("invalid event type: %s", eventType)
}
