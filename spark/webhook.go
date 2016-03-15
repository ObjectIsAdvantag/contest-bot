package main


import (
	"net/http"
	"encoding/json"

	"log"
	"fmt"
	"time"
)


func main() {

	port := "8080"
	log.Print("Starting webhook, listening at :", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


// Read new message
type NewMessageEvent struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Resource string `json:"resource"`
	Event string `json:"event"`
	Filter string `json:"filter"`
	Data struct {
		   ID string `json:"id"`
		   RoomID string `json:"roomId"`
		   PersonID string `json:"personId"`
		   PersonEmail string `json:"personEmail"`
		   Created time.Time `json:"created"`
	   } `json:"data"`
}

type SparkMessage struct {
	ID string `json:"id"`
	RoomID string `json:"roomId"`
	PersonID string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	Created time.Time `json:"created"`
	Text string `json:"text"`
}


func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Print("Expecting POST method as I am a Spark Webhook")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{ message:'I am the ContestBot, expecting POST as new messages are typed into the Spark Room' }"
		return
	}

	// Read incoming event
	decoder := json.NewDecoder(req.Body)
	var event NewMessageEvent
	if err := decoder.Decode(&event); err != nil {
		log.Print("Could not parse json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("Processing event: %v", event)

	// Retrieve message
	
	
	
	w.WriteHeader(http.StatusOK)
	return
}

