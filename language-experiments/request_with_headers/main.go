package request_with_headers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Event []struct {
	EventOrganizer string `json:"event_organizer"`
	UserEmail      string `json:"userEmail"`
	EventID        string `json:"event_id"`
	EventClub      string `json:"event_club"`
	EventName      string `json:"event_name"`
	EventLocation  string `json:"event_location"`
	EventYear      string `json:"event_year"`
	EventMetadata  struct {
		EventLocation string `json:"event_location"`
	} `json:"event_metadata"`
	EventHour             string    `json:"event_hour"`
	EventHands            string    `json:"event_hands"`
	EventMin              string    `json:"event_min"`
	EventImg              string    `json:"event_img"`
	EventDay              string    `json:"event_day"`
	EventDate             time.Time `json:"event_date"`
	EventRegistration     string    `json:"event_registration"`
	EventActivationStatus string    `json:"event_activation_status"`
	EventMonth            string    `json:"event_month"`
	EventResults          string    `json:"event_results"`
	EventNotes            string    `json:"event_notes"`
	Registered            []struct {
		Player2 string `json:"player2"`
		Player3 string `json:"player3"`
		Player1 string `json:"player1"`
		Player4 string `json:"player4"`
	} `json:"registered,omitempty"`
	EventTopThree []struct {
		Max   string `json:"Max"`
		Rank  string `json:"Rank"`
		Score string `json:"Score"`
		Names string `json:"Names"`
		Bds   string `json:"Bds"`
		Total string `json:"Total"`
	} `json:"event_top_three,omitempty"`
}

var myClient = &http.Client{}

func eventsHandlerWithHeaders(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("GET", "https://hyj6umtyma.execute-api.eu-west-1.amazonaws.com/Stage/events", nil)
	req.Header.Set("X-Api-Key", "oj7fp6P8mXas8zNN99jJU2FwE3QW26104hsA7okr")
	req.Header.Set("Content-Type", "application/json")

	resp, err := myClient.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return
	}
	defer resp.Body.Close()
	resp.Body = http.MaxBytesReader(w, resp.Body, 1048576)

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()

	var e Event

	if err := dec.Decode(&e);
	err!= nil {
		log.Println("Error ", err)
	}
	if dec.More() {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Event: %+v", e)
	log.Printf("%+v", e)
}

func FetchEventsWithHeaders(){
	r := mux.NewRouter()
	r.HandleFunc("/fetchEvents", eventsHandlerWithHeaders).Methods("GET")

	err := http.ListenAndServe("localhost:3002", r) //, GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}
