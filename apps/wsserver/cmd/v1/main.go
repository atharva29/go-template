package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/atharva29/go-template/apps/wsserver/internal/wsserver"
	"github.com/atharva29/go-template/pkg/kafka"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	fmt.Println("Start of Code")
	// Intialize public module KAFKA
	kafka.KafkaInitialize()
	flag.Parse()

	// Started Websocket server using internal WSSERVER module
	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		wsserver.ServeWs(w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
