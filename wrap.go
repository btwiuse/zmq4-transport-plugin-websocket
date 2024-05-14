// Copyright 2020 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package websocket implements the WebSocket transport for ØMQ sockets.
package websocket

import (
	"net"
	"net/http"

	"github.com/btwiuse/wsconn"
)

func Wrap(ln net.Listener) *listener {
	l := &listener{
		ln,
		make(chan *acceptResult, 1),
	}
	go http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsconn.Wrconn(w, r)
		l.result <- &acceptResult{conn, err}
	}))
	return l
}

type listener struct {
	net.Listener
	result chan *acceptResult
}

type acceptResult struct {
	conn net.Conn
	err  error
}

func (ln *listener) Accept() (net.Conn, error) {
	res := <-ln.result
	return res.conn, res.err
}
