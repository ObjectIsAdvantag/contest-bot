package main


import (
	"time"
	"log"
	"net/http"
	"encoding/json"
)


type SparkRoom struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Created time.Time `json:"created"`
	LastActivity time.Time `json:"lastActivity"`
	IsLocked bool `json:"isLocked"`
	SipAddress string `json:"sipAddress"`
}

// launch a new contest
func processLaunch(message SparkMessage) {
	// Retrieve Room SIP number
	client, err := http.NewRequest("GET", "https://api.ciscospark.com/v1/rooms/" + message.RoomID + "?showSipAddress=true", nil)
	if err != nil {
		log.Printf("Unexpected error, retrieving room details for RoomID: %s ", message.RoomID)
		sendToRoom(message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}
	token := "SPARK-API-TOKEN-HERE"
	client.Header.Add("Content-type", "application/json")
	client.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(client)
	if err != nil {
		log.Printf("Unexpected error while retrieving contents for room with id: %s ", message.RoomID)
		sendToRoom(message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}

	decoder := json.NewDecoder(response.Body)
	var room SparkRoom
	if err := decoder.Decode(&room); err != nil {
		log.Print("Could not parse json to decode SparkRoom")
		sendToRoom(message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}

	log.Print("Retrieved room details, sip number is " + room.SipAddress)

	// Inform participants a contest is starting in XX minutes

	// Call Tropo

}

// Send a room
func sendToRoom(roomID string, message string) {
	log.Print("Not implemented yet")

}

func processAnswer(message SparkMessage) {
	log.Print("Not implemented yet")
}


func processContribute(message SparkMessage) {
	log.Print("Not implemented yet")
}


