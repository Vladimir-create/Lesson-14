
package main

import (
	"log"
	"time"
	"fmt"
	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	//nickname = []byte{'Player1'}
)


func readMess() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	
	for {
		_, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("recv: %s", p)
	}
	go readMess()
}


func writeMess() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	var s string
	
	for {
		fmt.Scan(&s)
		err := c.WriteMessage(websocket.TextMessage, []byte(s) )
		if err != nil {
			log.Fatal("dial:", err)
		}
	}
}


func main (){
	go writeMess()
	go readMess()
	time.Sleep(100*time.Second)
}

