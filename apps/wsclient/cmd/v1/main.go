package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/atharva29/go-template/apps/wsclient/internal/v1/client"
	"github.com/atharva29/go-template/apps/wsserver/internal/v1/database" // THIS IS NOT ALLOWING ME TO USE INTERNAL PACKAGE OF DIFF APP
	"github.com/atharva29/go-template/pkg/logger"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

type counter struct {
	count int
	mu    sync.Mutex
}

func main() {
	counter := &counter{count: 0}
	defer func() {
		fmt.Println("Total Live connections were", counter.count)
	}()
	fmt.Println("Start of Code")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/subscribe"}
	log.Printf("connecting to %s", u.String())
	database.InitDBModule()
	client.PrintClient()

	done := make(chan struct{})
	logger.Log()
	for i := 0; i < 100000; i++ {
		go func() {
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				fmt.Println("Total Live connections were", counter.count)
				log.Fatal("dial:", err)
			}
			defer c.Close()
			defer close(done)
			counter.mu.Lock()
			counter.count++
			counter.mu.Unlock()

			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}()
		time.Sleep(1 * time.Millisecond)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Total Live connections were", counter.count)
			return
		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")
			fmt.Println("Total Live connections were", counter.count)
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			// err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			// if err != nil {
			// 	log.Println("write close:", err)
			// 	return
			// }
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
