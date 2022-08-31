package main

import(
	"fmt"
	"time"
	"net/http"
	"log"
	"bytes"
	"encoding/json"
)

const (
	siteUpCheckTimeInterval = 5
	siteUrl = "https://topiclist.xyz"
)

func main() {

	// Idk what i am coding at this point
	for range time.Tick(siteUpCheckTimeInterval * time.second) {
	// Send request to site 

	resp, err:= https.Get(siteUrl)
	if(err !=nil) {
		log.Fatal(err)	
	}
	// If responce is down then
	if resp.StatusCode != 200 {
	 fmt.Println("Website Response Is :", resp.StatusCode, http.StatusText(resp.StatusCode))
  } else {
	fmt.println("All Good Captin")
    }

  }


}

// Send Webhook Using Discord GO

func main() {
	hook := goWebhook.CreateWebhook()
	hook.AddField("New Title","New Value",true)
	hook.SendWebook("https://discordapp.com/api/yourwebhook") // no tests to check if webhook was successful
	// or
	webhookReq, err := hook.SendWebook("https://discordapp.com/api/yourwebhook")    // use variable to check if post request was successful
	
	if err !=nil{
	fmt.Println("Unexpected error sending webhook!")      
	
	if webhookReq.StatusCode == 204 { //204 status is successful webhook post
		if webhookReq.StatusCode == 200 { //200 Stats is send
	fmt.Println("Webhook sent")
	}else{
	fmt.Println("Webhook failed")
	fmt.Println(webhookReq.StatusCode)
  
            }

         }
      }
}

//sending a damn email to the owner


