// This script has been created to do GET request to riskiq
// to read current account metadata and settings.
// Go to the setting section of your account to grab
// the Email and APIKEY.
// Location: https://community.riskiq.com/settings

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.passivetotal.org/v2/account"
	Email := "INSERT EMAIL"
	APIKey := "INSERT APIKEY"
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(Email, APIKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
