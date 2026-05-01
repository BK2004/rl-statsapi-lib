package main

import (
	"fmt"
	"os"
	"os/signal"
	"rl-statsapi-parser/events"
	_ "rl-statsapi-parser/listener"
)

func main() {
	for ch := range events.BallHit.Subscribe() {
		fmt.Printf("Received event: %v\n", ch.Ball.Location)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Exiting...")
}
