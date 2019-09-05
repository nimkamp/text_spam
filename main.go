package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	for {
		// Set account keys & information
		accountSid := "ACCOUNT_SID"
		authorizationToken := "AUTHORIZATION_TOKEN"
		urlString := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

		quotes := [7]string{"hello",
			"Message 1",
			"Message 2",
			"Message 3",
			"Message 4",
			"Message 5",
			"Message 6"}

		messageData := url.Values{}
		messageData.Set("To", `NUMBER_TO`)
		messageData.Set("From", `NUMBER_FROM`)
		messageData.Set("Body", quotes[rand.Intn(len(quotes))])
		messageDataReader := strings.NewReader(messageData.Encode())

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPost, urlString, messageDataReader)
		req.SetBasicAuth(accountSid, authorizationToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, _ := client.Do(req)
		if resp.StatusCode > 200 && resp.StatusCode < 300 {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			fmt.Println(data)
			if err == nil {
				fmt.Println(data["sid"])
			}
		} else {
			fmt.Println(resp.Status)
		}

		// Set up random generator
		rand.Seed(time.Now().Unix())
		fmt.Println(rand.Intn(len(quotes)))
	}
}
