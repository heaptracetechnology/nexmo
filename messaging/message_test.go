package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type SendSMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

var _ = Describe("Send SMS", func() {

	sms := SendSMS{From: "9876543210", To: "0123456789", Text: "Testing SMS"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sms)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send SMS message", func() {
		Context("send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Send SMS Negative", func() {

	sms := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sms)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send email message", func() {
		Context("send", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Receive SMS", func() {

	os.Setenv("API_KEY", "26ee1c8b")
	os.Setenv("API_SECRET", "gGJmYyeEh5Po05f2")

	sms := SendSMS{From: "9876543210", To: "0123456789", Text: "Testing SMS"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sms)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/receive", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Receiver)
	handler.ServeHTTP(recorder, request)

	Describe("Receive SMS message", func() {
		Context("receive", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
