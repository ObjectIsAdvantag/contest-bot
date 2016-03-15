package main


import (
	"net/http"
	"log"
	"fmt"
)


func main() {

	http.HandleFunc("/spark", sparkHandler)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "webhook is active")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Read new message
type NewMessagePayload struct {
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


func sparkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatal("Expecting POST for Spark Webhooks")
		fmt.Fprintf(w, "I am the ContestBot, listening to POST methods and only those")
		return
	}

	// Read incoming message
}