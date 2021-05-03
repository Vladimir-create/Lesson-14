package main

import (
	"net/http"
	"io"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var mas = make([]*websocket.Conn, 0, 0)


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



func Socket(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	mas = append(mas, conn)
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
		
		for i:=0; i<len(mas)-1; i++{
			if mas[i] == conn {i++}
			if err := mas[i].WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
			log.Println(p)
		}
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/socket", Socket)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}


