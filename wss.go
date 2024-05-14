// Copyright 2020 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package websocket implements the WebSocket transport for ØMQ sockets.
package websocket

import (
	"context"
	"fmt"
	"net"

	"github.com/go-zeromq/zmq4/transport"
	"github.com/webteleport/webteleport/transport/websocket"
)

// Transport implements the zmq4 transport protocol for WebSockets.
type Transport struct {
	Secure bool
}

// Dial connects to the address on the named network using the provided
// context.
func (Transport) Dial(ctx context.Context, dialer transport.Dialer, addr string) (net.Conn, error) {
	return websocket.DialConn(ctx, addr, nil)
}

// Listen announces on the provided network address.
func (Transport) Listen(ctx context.Context, addr string) (net.Listener, error) {
	ln, err := websocket.Listen(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("websocket: %w", err)
	}
	return Wrap(ln), nil
}

// Addr returns the end-point address.
func (t Transport) Addr(ep string) (addr string, err error) {
	if !t.Secure {
		return "ws://" + ep, nil
	}
	return "wss://" + ep, nil
}

var (
	_ transport.Transport = (*Transport)(nil)
)
