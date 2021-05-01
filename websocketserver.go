package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		log.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else if req.Method == "OPTIONS" {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(405)
	}
	
}

func Socketread(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("recv: %s", p)
		
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		log.Println(p)
	}
}

func Socketwrite(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var s string 
	
	for {
		fmt.Scan(&s)
		if err := conn.WriteMessage(1, []byte(s)); err != nil {
			log.Println(err)
			return
		}
		log.Println(s)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/socket2", Socketwrite)
	http.HandleFunc("/socket", Socketread)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

