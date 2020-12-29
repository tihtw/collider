// Copyright (c) 2014 The WebRTC project authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"flag"
	"log"

	"github.com/tihtw/collider"
)

var tls = flag.Bool("tls", true, "whether TLS is used")
var port = flag.Int("port", 443, "The TCP port that the server listens on")
var roomSrv = flag.String("room-server", "https://appr.tc", "The origin of the room server")
var tlsCert = flag.String("tls-cert", "/cert/cert.pem", "The TLS cert file path")
var tlsKey = flag.String("tls-key", "/cert/key.pem", "The TLS server key file path")

func main() {
	flag.Parse()

	log.Printf("Starting collider: tls = %t, port = %d, room-server=%s, "+
		"tls-cert=%s, tls-key=%s", *tls, *port, *roomSrv, *tlsCert, *tlsKey)

	c := collider.NewCollider(*roomSrv)
	c.Run(*port, *tls, *tlsCert, *tlsKey)
}
