/**
* Go File created on 13.01.23
* by Antonio Masotti (antonio)
* MIT License
 */

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var targetServer = "https://echo.zuplo.io/"

// ************** TYPES ************** //

type SampleMessage struct {
	Id         string `json:"_id"`
	CrmId      string `json:"crmId"`
	IsActive   bool   `json:"isActive"`
	Age        int    `json:"age"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Company    string `json:"company"`
	Email      string `json:"email"`
	UpdateTime int64  `json:"updateTime"`
}

// ************** SPIELEREI ************** //

func (p *SampleMessage) ToJsonString() (string, error) {
	s, err := json.Marshal(p)
	if err != nil {
		log.Println("Error marshalling SampleMessage")
		return "", err
	}

	return string(s), nil
}

func parseBodyAndAddProperty(message string) SampleMessage {
	var payload SampleMessage
	fmt.Println("Parsing message", message)
	//if err := json.Unmarshal([]byte(message), &message);
	if err := json.NewDecoder(strings.NewReader(message)).Decode(&payload); err != nil {
		log.Fatal("Error unmarshalling message", err)
	}

	payload.UpdateTime = time.Now().Unix()

	fmt.Println(payload)

	return payload
}

// ************** HAUPTLOGIK ************** //

func SendToApi(message string) {

	p := parseBodyAndAddProperty(message)

	resp, success := sendRequest(p)
	if !success {
		log.Println("Error sending request")
		return
	}

	fmt.Println("Request successfully sent")

	// Print the response
	headers := resp.Header
	fmt.Printf("Headers: %v", headers)

	body := resp.Body
	fmt.Printf("Body: %v", body)

}

func sendRequest(p SampleMessage) (*http.Response, bool) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request
	payload, _ := p.ToJsonString()
	req, err := http.NewRequest(http.MethodPost, targetServer, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return resp, true
}
