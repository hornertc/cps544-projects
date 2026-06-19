package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Process ID:", os.Getpid())
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c)

	// Block until any signal is received.
	for s := range c {
		fmt.Printf("Got signal: %s (%d)\n", s, s)
	}
}
