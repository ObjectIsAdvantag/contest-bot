package main


import (
	"time"
	"net/http"
	"encoding/json"
	"strings"
	"fmt"

	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
"google.golang.org/appengine/urlfetch"
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
func processLaunch(ctx context.Context, message SparkMessage) {
	// Retrieve Room SIP number
	client, err := http.NewRequest("GET", "https://api.ciscospark.com/v1/rooms/" + message.RoomID + "?showSipAddress=true", nil)
	if err != nil {
		log.Errorf(ctx, "Unexpected error, retrieving room details for RoomID: %s ", message.RoomID)
		sendMessageToRoom(ctx, message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}
	client.Header.Add("Content-type", "application/json")
	client.Header.Add("Authorization", "Bearer " + env.sparkToken)

	response, err := urlfetch.Client(ctx).Do(client)
	if err != nil {
		log.Errorf(ctx, "Unexpected error while retrieving contents for room with id: %s ", message.RoomID)
		sendMessageToRoom(ctx, message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}

	decoder := json.NewDecoder(response.Body)
	var room SparkRoom
	if err := decoder.Decode(&room); err != nil {
		log.Errorf(ctx, "Could not parse json to decode SparkRoom")
		sendMessageToRoom(ctx, message.RoomID, "Cannot launch a new context for now, Sorry for that, Try again later")
		return
	}
	log.Debugf(ctx, "Retrieved room details, sip number is " + room.SipAddress)

	// Inform participants a contest is starting
	sendMessageToRoom(ctx, message.RoomID, "A new contest is starting, get ready !")

	// TODO: Pick a contest
	//contestAudio := "http://soundbible.com/mp3/I%20Love%20You%20Daddy-SoundBible.com-862095235.mp3"

	// Invoke Tropo script, see Readme and file newcontest.js
	params := fmt.Sprintf("room_sip=%s&replays=%d&botname=%s", room.SipAddress, 2, "ContestBot@mail.com")
	payload := strings.NewReader(params)
	req, _ := http.NewRequest("POST",
		"https://api.tropo.com/1.0/sessions?action=create&token=" + env.tropoToken,
		payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := urlfetch.Client(ctx).Do(req)
	if err != nil {
		log.Errorf(ctx,"Communication error with Tropo, err: %s", err)
		sendMessageToRoom(ctx, message.RoomID, "Contest failed to launch, Sorry for that, Try again later")
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf(ctx, "Tropo script invocation error: %s", err)
		sendMessageToRoom(ctx, message.RoomID, "Contest failed to launch, Sorry for that, Try again later")
		return
	}

	log.Infof(ctx, "New contest launched successfully")
}


func processAnswer(ctx context.Context, message SparkMessage) {
	log.Warningf(ctx, "Not implemented yet")
}


func processContribute(ctx context.Context, message SparkMessage) {
	log.Warningf(ctx, "Not implemented yet")
}


