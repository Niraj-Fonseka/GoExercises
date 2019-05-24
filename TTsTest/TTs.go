// # needs mplayer installed
// # apt install mplayer
// #
// # install htgo-tts
// # go get "github.com/hegedustibor/htgo-tts"

package main

import "github.com/hegedustibor/htgo-tts"

func main() {
    speech := htgotts.Speech{Folder: "audio", Language: "en"}
    speech.Speak("Person detected")
}