package messaging

import (
	bytes "bytes"
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-nexmo/result"
	"github.com/nexmo-community/nexmo-go"
	"log"
	"net/http"
	"os"
	"time"
)

type Payload struct {
	EventId     string       `json:"eventID"`
	EventType   string       `json:"eventType"`
	ContentType string       `json:"contentType"`
	Data        EmailMessage `json:"data"`
}

type Subscribe struct {
	Data          RequestParam `json:"data"`
	Endpoint      string       `json:"endpoint"`
	Id            string       `json:"id"`
	LastMessageId uint32
	IsTesting     bool `json:"istesting"`
}

type RequestParam struct {
}

var Listener = make(map[string]Subscribe)
var rtmStarted bool

//Send SMS
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var apiSecret = os.Getenv("API_SECRET")

	decoder := json.NewDecoder(request.Body)
	var requestBody nexmo.SendSMSRequest
	decodeErr := decoder.Decode(&requestBody)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(apiKey, apiSecret)
	client := nexmo.NewClient(http.DefaultClient, auth)

	requestBody.APIKey = apiKey
	requestBody.APISecret = apiSecret

	smsResponse, httpResponse, smsErr := client.SMS.SendSMS(requestBody)
	if smsErr != nil {
		result.WriteErrorResponse(responseWriter, smsErr)
		return
	}

	bytes, _ := json.Marshal(smsResponse)
	result.WriteJsonResponse(responseWriter, bytes, httpResponse.StatusCode)

}

//Receiver SMS
func Receiver(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var apiSecret = os.Getenv("API_SECRET")

	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(apiKey, apiSecret)
	client := nexmo.NewClient(http.DefaultClient, auth)

	//client.SMS.SetBaseURL("https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9")

	decoder := json.NewDecoder(request.Body)

	var sub Subscribe
	decodeError := decoder.Decode(&sub)
	if decodeError != nil {
		result.WriteErrorResponse(responseWriter, decodeError)
		return
	}

	Listener[sub.Data.Username] = sub
	if !rtmStarted {
		go SMSRTM()
		rtmStarted = true
	}

	bytes, _ := json.Marshal("Subscribed")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

func SMSRTM() {
	isTest := false
	for {
		if len(Listener) > 0 {
			for k, v := range Listener {
				go getMessageUpdates(k, v)
				isTest = v.IsTesting
			}
		} else {
			rtmStarted = false
			break
		}
		time.Sleep(5 * time.Second)
		if isTest == true {
			break
		}
	}
}

func getMessageUpdates(userid string, sub Subscribe) {

	Listener[userid] = sub
	var payload Payload
	payload.ContentType = "application" + "/" + "json"
	payload.EventType = "hears"
	payload.EventId = sub.Id
	payload.Data = data

	requestBody := new(bytes.Buffer)
	encodeError := json.NewEncoder(requestBody).Encode(payload)
	if encodeError != nil {
		log.Fatalln(encodeError)
		fmt.Println("err :", encodeError)
	}

	// sub.Endpoint = "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9"

	// fmt.Println("sub.Endpoint ::", sub.Endpoint)

	res, reserror := http.Post(sub.Endpoint, "application/json", requestBody)
	if reserror != nil {
		fmt.Println("err :", reserror)
	}
	fmt.Println("res :", res)
}
