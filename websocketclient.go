
package main

import (
	"log"
	"fmt"
	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	//nickname = []byte{'Player1'}
)

func readMess(c *websocket.Conn) {
	
	for {
		_, p, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("recv: %s", p)
	}
	readMess(c)
}


func writeMess(c *websocket.Conn) {
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
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	
	go writeMess(c)
	readMess(c)
}

