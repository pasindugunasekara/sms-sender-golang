package main

import (
  "fmt"
  "strings"
  "math/rand"
  "time"
  "net/http"
  "net/url"
  "encoding/json"
)

func main() {
  
  accountSid := "ACc5df64312a9f5370600540f9ddf4887d"
  authToken := "d8e6db7e824c50dbc12c48f462a74492"
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

  
  quotes := [7]string{"I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.'",
                      "Peculiar travel suggestions are dancing lessons from God.",
                      "There's only one rule that I know of, babies—God damn it, you've got to be kind.",
                      "Many people need desperately to receive this message: 'I feel and think much as you do, care about many of the things you care about, although most people do not care about them. You are not alone.'",
                      "That is my principal objection to life, I think: It's too easy, when alive, to make perfectly horrible mistakes.",
                      "So it goes.",
                      "We must be careful about what we pretend to be."}

  
  rand.Seed(time.Now().Unix())

  
  msgData := url.Values{}
  msgData.Set("To","++94776494697")
  msgData.Set("From","+18055002838")
  msgData.Set("Body",quotes[rand.Intn(len(quotes))])
  msgDataReader := *strings.NewReader(msgData.Encode())

  
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  
  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}