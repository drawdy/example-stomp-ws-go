/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"
	"net/url"

	stomp "github.com/drawdy/stomp-ws-go"
	"github.com/gorilla/websocket"
)

func main() {

	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/stomp-ws",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("Couldn't connect to %v: %v", u.String(), err)
	}
	log.Print("Websocket connection succeeded.")

	h := stomp.Headers{
		stomp.HK_ACCEPT_VERSION, "1.2,1.1,1.0",
		stomp.HK_HEART_BEAT, "3000,3000",
		stomp.HK_HOST, u.Host,
	}
	sc, err := stomp.ConnectOverWS(conn, h)
	if err != nil {
		log.Fatalf("Couldn't create stomp connection: %v", err)
	}

	err = DoSubscribe(sc)
	if err != nil {
		log.Fatalf("Failed to do subscribe: %v", err)
	}

	err = sc.Disconnect(stomp.NoDiscReceipt)
	if err != nil {
		log.Fatalf("Failed to disconnect: %v", err)
	}

	log.Print("Disconnected.")
}
