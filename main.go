package main

import (
	"encoding/json"
	"etl/middleware"
	"etl/model"
	"etl/utils"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", middleware.RecoverMiddleware(http.HandlerFunc(handleEvents)))
	log.Println("server is starting")
	utils.FatalErr(http.ListenAndServe(":8000", nil), "Error while starting server")
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	utils.PanicErr(err, "Error reading request body")

	m := map[string]any{}
	err = json.Unmarshal(data, &m)
	utils.PanicErr(err, "Error while unmarshaling request body []byte")

	eventType, ok := m["event_type"].(string)
	if !ok {
		utils.PanicErr(fmt.Errorf("invalid event_type"), "Invalid event_type received from request body")
		// todo this should be returned to caller with 4xx error
	}
	event, err := concreteEvent(eventType, data)
	utils.PanicErr(err, "Error while getting concrete event")

	err = event.SaveToDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message ingested successfully"))
}

func concreteEvent(eventType string, data []byte) (model.Event, error) {
	if eventType == "login" {
		var login model.Login
		err := json.Unmarshal(data, &login)
		if err != nil {
			return nil, err
		}
		return &login, nil
	}

	if eventType == "purchase" {
		var purchase model.Purchase
		err := json.Unmarshal(data, &purchase)
		if err != nil {
			return nil, err
		}
		return &purchase, nil
	}

	if eventType == "level_up" {
		var levelUp model.LevelUp
		err := json.Unmarshal(data, &levelUp)
		if err != nil {
			return nil, err
		}
		return &levelUp, nil
	}

	return nil, fmt.Errorf("invalid event type: %s", eventType)
}
