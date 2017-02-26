package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http server address")
var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:",err)
		return
	}
	defer c.Close();
	for {
		mt, message, err := c.ReadMessage();
		if err != nil {
			fmt.Println("upgrade:",err)
			break
		}
		log.Printf("Recv: %s", message)
		err = c.WriteMessage(mt, message);
		if err != nil {
			fmt.Println("upgrade:",err)
			break
		}
	}
}

func main(){
	flag.Parse();
	log.SetFlags(0);
	http.Handle("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}