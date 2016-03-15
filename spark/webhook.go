package main


import (
	"net/http"
	"encoding/json"

	"log"
	"fmt"
	"time"
	"strings"

	"github.com/spf13/viper"
)

func main() {

	// load env variables : BOT_SPARK_TOKEN and BOT_TROPO_TOKEN
	viper.SetEnvPrefix("bot") // will be uppercased automatically
	viper.BindEnv("spark_token") // will be uppercased automatically
	viper.BindEnv("tropo_token") // will be uppercased automatically

	// launch server
	port := "8080"
	log.Print("Starting webhook, listening at :", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
		log.Printf("Could not parse json to decode NewMessageEvent, err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print("Processing event: %v", event)

	// Call Spark to retrieve message details, the text essentially
	client, err := http.NewRequest("GET", "https://api.ciscospark.com/v1/messages/" + event.Data.ID, nil)
	if err != nil {
		log.Printf("Unexpected error; %s, while processing event: %s, retrieving message id: %s ", err, event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client.Header.Add("Content-type", "application/json")
	client.Header.Add("Authorization", "Bearer " + viper.GetString("spark_token"))

	response, err := http.DefaultClient.Do(client)
	if err != nil {
		log.Printf("Unexpected error: %s, while retrieving contents for message id: %s ", err, event.ID, event.Data.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder = json.NewDecoder(response.Body)
	var message SparkMessage
	if err := decoder.Decode(&message); err != nil {
		log.Printf("Could not decode SparkMessage, err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("New message: %v", message)

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

