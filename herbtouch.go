package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gen2brain/beeep"
)

func notify() {
	err := beeep.Notify("20 Minutes Have Passed!", "Please abide by the 20/20/20 rule and get up for 20 minutes", "gopher.jpg")
	if err != nil {
		panic(err)
	}
}
func play() {
	alarm, err := os.Open("AlarmClock.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(alarm)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}

}
func main() {
	for {
		time.Sleep(20 * time.Minute)
		go play()
		go notify()
		time.Sleep(20 * time.Minute)
	}
}
