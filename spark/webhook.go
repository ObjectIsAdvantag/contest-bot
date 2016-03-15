package main


import (
	"net/http"
	"encoding/json"

	"log"
	"fmt"
	"time"
	"strings"
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
		w.Header().Add("Content-type", "application/json")
		fmt.Fprintf(w, `{ "message":"I am the ContestBot, expecting POST as new messages are typed into the Spark Room" }`)
		return
	}

	// Read incoming event
	decoder := json.NewDecoder(req.Body)
	var event NewMessageEvent
	if err := decoder.Decode(&event); err != nil {
		log.Print("Could not parse json to decode NewMessageEvent")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("Processing event: %v", event)

	// Retrieve message
	client, err := http.NewRequest("GET", "https://api.ciscospark.com/v1/messages/" + event.Data.ID, nil)
	if err != nil {
		log.Printf("Unexpected error while processing event: %s, retrieving message id: %s ", event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token := "MmIzYTk0MWYtYzY2My00MjgzLThlZDQtOWU5ZmU1MzdiOTNiZTVhZWExOGEtM2Rh"
	client.Header.Add("Content-type", "application/json")
	client.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(client)
	if err != nil {
		log.Printf("Unexpected error while retrieving contents for message id: %s ", event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Read message details
	decoder = json.NewDecoder(response.Body)
	var message SparkMessage
	if err := decoder.Decode(&message); err != nil {
		log.Print("Could not parse json to decode SparkMessage")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("New message: %v", message)

	// Process message
	go processMessage(message)

	w.WriteHeader(http.StatusOK)
	return
}

func processMessage(message SparkMessage) {
	// /launch
	if strings.HasPrefix(message.Text, "/launch") {
		log.Printf("Processing launch command")
		processLaunch(message)
		return
	}

	// /guess
	if strings.HasPrefix(message.Text, "/guess") {
		log.Printf("Processing guess command")

		processAnswer(message)
		return
	}

	// /contribute
	if strings.HasPrefix(message.Text, "/contribute") {
		log.Printf("Processing contribute command")
		processContribute(message)
		return
	}
}

