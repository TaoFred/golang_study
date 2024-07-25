package main

import (
	"fmt"
	"os/signal"
	"runtime"
	"syscall"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	// nc.Publish("hello", []byte("Just do it"))

	nc.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}
