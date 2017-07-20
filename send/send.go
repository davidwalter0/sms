package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/davidwalter0/go-cfg"
)

type SMS struct {
	AccountSid string `usage:"Account from twilio: https://www.twilio.com/console"`
	AuthToken  string `usage:"Secret API tokey from twilio"`
	FromPhone  string `usage:"Twilio allocated phone number, https://www.twilio.com/console/phone-numbers/getting-started"`
	ToPhone    string `usage:"Send SMS message to this destination"`
	Text       string `usage:"Send text as SMS message"`
}

var app SMS

func missingArgError() {
	JSON, _ := json.MarshalIndent(&app, "", "  ")
	log.Fatalf("One or more required fields are not set \n%s\n%s\n", app.Text, string(JSON))
}

func init() {

	cfg.HelpText(`
Send an sms message from the command line or env variables.  Either
use the flag for text or env var. Override env or flag with command
line text. Args after the flags will be used as the text message. 

Skip flag/option arguments until last id marker "--" is seen `)

	var err error

	if err = cfg.Parse(&app); err != nil {
		log.Fatalf("%v\n", err)
	}
	var found bool
	os.Args = os.Args[1:]
	arglist := make([]string, 0)
	// Override env or flag with command line text
	// skip flag/option arguments until last id marker "--" is seen
	for _, arg := range os.Args {
		if !found {
			if len(arg) == 2 && arg[0:2] == "--" {
				found = true
			} else {
				if len(arg) > 0 && arg[0] == '-' {
					continue
				}
			}
		}
		arglist = append(arglist, arg)
	}

	if len(arglist) > 0 {
		app.Text = strings.Join(arglist, ":")
	}
	if len(app.Text) == 0 {
		missingArgError()
	}

	if len(app.AccountSid) == 0 ||
		len(app.AuthToken) == 0 ||
		len(app.FromPhone) == 0 ||
		len(app.ToPhone) == 0 ||
		len(app.Text) == 0 {
		log.Println("here")
		missingArgError()
	}
}

func main() {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + app.AccountSid + "/Messages.json"
	fmt.Println(urlStr)
	v := url.Values{}
	v.Set("To", app.ToPhone)
	v.Set("From", app.FromPhone)
	v.Set("Body", app.Text)
	rb := *strings.NewReader(v.Encode())
	fmt.Println(rb)
	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(app.AccountSid, app.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Printf("Response from Twilio SMS service\n%s\n", data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
