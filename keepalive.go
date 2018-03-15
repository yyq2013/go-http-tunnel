// Copyright (C) 2017 Michał Matczuk
// Use of this source code is governed by an AGPL-style
// license that can be found in the LICENSE file.

// +build !windows

package tunnel

import (
	"net"
	"time"

	"github.com/felixge/tcpkeepalive"
)

func keepAlive(conn net.Conn) error {
	return tcpkeepalive.SetKeepAlive(conn, DefaultKeepAliveIdleTime, DefaultKeepAliveCount, DefaultKeepAliveInterval)
}
