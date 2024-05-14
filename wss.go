// Copyright 2020 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package websocket implements the WebSocket transport for ØMQ sockets.
package websocket

import (
	"context"
	"net"

	"github.com/go-zeromq/zmq4/transport"
)

// Transport implements the zmq4 transport protocol for WebSockets.
type Transport struct{}

// Dial connects to the address on the named network using the provided
// context.
func (Transport) Dial(ctx context.Context, dialer transport.Dialer, addr string) (net.Conn, error) {
	panic("not implemented")
}

// Listen announces on the provided network address.
func (Transport) Listen(ctx context.Context, addr string) (net.Listener, error) {
	panic("not implemented")
}

// Addr returns the end-point address.
func (Transport) Addr(ep string) (addr string, err error) {
	panic("not implemented")
}

var (
	_ transport.Transport = (*Transport)(nil)
)
