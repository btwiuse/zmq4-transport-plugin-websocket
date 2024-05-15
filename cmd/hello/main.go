package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	zmq "github.com/go-zeromq/zmq4"
)

func Run(args []string) error {
	return nil
}

func main() {
	go func() {
		err := hwserver()
		if err != nil {
			log.Fatalln("hwserver:", err)
		}
	}()
	time.Sleep(3 * time.Second)
	err := hwclient()
	if err != nil {
		log.Fatalln("hwclient:", err)
	}
}

var RELAY = getEnv("RELAY", "https://example.com/zmq/")

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func hwserver() error {
	ctx := context.Background()
	// Socket to talk to clients
	socket := zmq.NewRep(ctx)
	defer socket.Close()
	if err := socket.Listen(RELAY); err != nil {
		return fmt.Errorf("listening: %w", err)
	}

	for {
		msg, err := socket.Recv()
		if err != nil {
			return fmt.Errorf("receiving: %w", err)
		}
		fmt.Println("Received ", msg)

		// Do some 'work'
		time.Sleep(time.Second)

		reply := fmt.Sprintf("World")
		if err := socket.Send(zmq.NewMsgString(reply)); err != nil {
			return fmt.Errorf("sending reply: %w", err)
		}
	}
}

func hwclient() error {
	ctx := context.Background()
	socket := zmq.NewReq(ctx, zmq.WithDialerRetry(time.Second))
	defer socket.Close()

	fmt.Printf("Connecting to hello world server...")
	if err := socket.Dial(RELAY); err != nil {
		return fmt.Errorf("dialing: %w", err)
	}

	for i := 0; i < 10; i++ {
		// Send hello.
		m := zmq.NewMsgString("hello")
		fmt.Println("sending ", m)
		if err := socket.Send(m); err != nil {
			return fmt.Errorf("sending: %w", err)
		}

		// Wait for reply.
		r, err := socket.Recv()
		if err != nil {
			return fmt.Errorf("receiving: %w", err)
		}
		fmt.Println("received ", r.String())
	}
	return nil
}
