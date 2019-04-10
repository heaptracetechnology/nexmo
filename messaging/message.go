package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-nexmo/result"
	"github.com/nexmo-community/nexmo-go"
	"net/http"
	"os"
)

//Send SMS
func Send(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
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
