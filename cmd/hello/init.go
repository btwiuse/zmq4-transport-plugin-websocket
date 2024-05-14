package main

import (
	"fmt"

	websocket "github.com/btwiuse/zmq4-transport-plugin-websocket"
	zmq "github.com/go-zeromq/zmq4"
)

func init() {
	must := func(err error) {
		if err != nil {
			panic(fmt.Errorf("%+v", err))
		}
	}

	transportWS := &websocket.Transport{}
	transportWSS := &websocket.Transport{
		Secure: true,
	}
	must(zmq.RegisterTransport("ws", transportWS))
	must(zmq.RegisterTransport("http", transportWS))
	must(zmq.RegisterTransport("wss", transportWSS))
	must(zmq.RegisterTransport("https", transportWSS))
}
