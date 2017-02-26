package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main()  {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

}