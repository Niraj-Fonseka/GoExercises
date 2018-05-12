package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TextToSpeechResponse struct {
	AudioContent string `json:"audioContent"`
}

func main() {
	url := "https://texttospeech.googleapis.com/v1beta1/text:synthesize?key=" + os.Getenv("KEY")
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{
		"audioConfig": {
		  "audioEncoding": "LINEAR16",
		  "pitch": "0.00",
		  "speakingRate": "1.00"
		},
		"input": {
		  "text": "Hello Kayla , Nolan and Carol ! This is google du!"
		},
		"voice": {
		  "languageCode": "en-US",
		  "name": "en-US-Wavenet-E"
		}
	  }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var responseTTS TextToSpeechResponse

	json.Unmarshal(body, &responseTTS)

	f, err := os.Create("audio.txt")
	defer f.Close()

	n3, err := f.WriteString(responseTTS.AudioContent)
	fmt.Printf("wrote %d bytes\n", n3)

}
