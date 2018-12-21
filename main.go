package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const apiServer = "http://api.pullword.com/get.php?source="

// Getting the rate
type RateJson struct {
	Words string `json:"t"`
	Rate  string `json:"p"`
}

// Rate Body Return
type RateBody struct {
	Body []RateJson
}

func requestPhrase(phrase string, precise, debug bool) (word string) {
	//by default precise is false and debug are true
	param1 := "0"
	param2 := "1"

	if precise == true {
		param1 = "1"
	}

	if debug == false {
		param2 = "0"
	}

	requestURL := apiServer + phrase + "&param1=" + param1 + "&param2=" + param2 + "&json=1"
	u, err := url.Parse(requestURL)
	if err != nil {
		log.Print(err)
	}
	log.Println(u)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	var data []RateJson
	json.Unmarshal(body, &data)
	//fmt.Printf("%s", data)

	for _, v := range data {
		log.Print(v.Words)
		if v.Rate == "1" {
			word = v.Words

		}
	}

	return string(word)
}

func main() {

	myWords := "李彦宏是马云最大威胁嘛?"

	data := requestPhrase(myWords, false, true)

	fmt.Print(data)
}
