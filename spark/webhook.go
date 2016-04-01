// Spark webhook to be hosted on google app engine
//
package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
"google.golang.org/appengine/urlfetch"
)

func init() {
	// Handlers needs to be placed here for GAE compat
	http.HandleFunc("/", healthCheckHandler)
	http.HandleFunc("/spark", sparkHandler)
}


// NewMessageEvent defines the JSON structure posted by Spark to WebHooks
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

// SparkMessage defines the JSON structure retrieved when asking the Spark API for message details
// Long story short: the NewMessageEvent structure contains a Data property which contains all message properties except the Text
// A second call needs to be issued to get the message text in readable format (not crypted)
type SparkMessage struct {
	ID string `json:"id"`
	RoomID string `json:"roomId"`
	PersonID string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	Created time.Time `json:"created"`
	Text string `json:"text"`
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{ "message":"I am the ContestBot, you've hitted my HealthCheck endpoint, expecting GET method only here" }`)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{
			"description": "I am the ContestBot, you hitted my HealthCheck endpoint successfully",
			"settingsOK": %v,
			"sparkURI": "/spark" }`, env.isCorrect)
	return
}

func sparkHandler(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.Method != http.MethodPost {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{ "status":%d,
		 		"message":"I am the ContestBot, GET not supported, expecting POST as new messages are posted to a Spark Room" }`,
				http.StatusBadRequest)
		return
	}

	// Read incoming event
	decoder := json.NewDecoder(req.Body)
	var event NewMessageEvent
	if err := decoder.Decode(&event); err != nil {
		log.Errorf(ctx, "Could not parse json to decode NewMessageEvent, err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debugf(ctx, "Processing event: %v", event)

	// Call Spark to retrieve message details, the text essentially
	client, err := http.NewRequest("GET", "https://api.ciscospark.com/v1/messages/" + event.Data.ID, nil)
	if err != nil {
		log.Errorf(ctx, "Unexpected error; %s, while processing event: %s, retrieving message id: %s ", err, event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client.Header.Add("Content-type", "application/json")
	client.Header.Add("Authorization", "Bearer " + env.sparkToken)

	response, err := urlfetch.Client(ctx).Do(client)
	if err != nil {
		log.Errorf(ctx, "Unexpected error: %s, while retrieving contents for message id: %s ", err, event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder = json.NewDecoder(response.Body)
	var message SparkMessage
	if err := decoder.Decode(&message); err != nil {
		log.Errorf(ctx, "Could not decode SparkMessage, err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Infof(ctx, "New message: %v", message)

	// Process message
	// WORKAROUND : removed // execution for compliance with GAE
	// go processMessage(ctx, message)
	processMessage(ctx, message)

	w.WriteHeader(http.StatusOK)
	return
}

func processMessage(ctx context.Context, message SparkMessage) {
	// /launch
	if strings.HasPrefix(message.Text, "/launch") {
		log.Debugf(ctx, "Processing launch command")
		processLaunch(ctx, message)
		return
	}

	// /guess
	if strings.HasPrefix(message.Text, "/guess") {
		log.Debugf(ctx, "Processing guess command")
		processAnswer(ctx, message)
		return
	}

	// /contribute
	if strings.HasPrefix(message.Text, "/contribute") {
		log.Debugf(ctx, "Processing contribute command")
		processContribute(ctx, message)
		return
	}
}

