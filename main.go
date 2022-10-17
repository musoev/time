package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	ip := "localhost"
	port := ":8080"
	addres:=ip+port

	http.HandleFunc("/", HelloWeb)
	err:=http.ListenAndServe(addres, nil)
	if err !=nil {
		log.Println("Listen adn server error", err)
	}


}

func HelloWeb(response http.ResponseWriter, request *http.Request) {
	//time.
	
	t:=time.Now() 
	t2:=t.String()
	
	write, err:=response.Write([]byte(t2))
	if err !=nil {
		log.Println(err)
	}
	log.Println("lenth send",write)
}
