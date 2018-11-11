package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type TextToSpeechResponse struct {
	AudioContent string `json:"audioContent"`
}

//base64 source_base64_text_file -d > dest_audio_file

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
			"text": "Hello ! How's it going"
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

	cmdBase := exec.Command("base64", "audio.txt -d > dest_audio_file.mp3")
	if err := cmdBase.Run(); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("play", "dest_audio_file.mp3")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
