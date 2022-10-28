package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Son0-0/Goginx/handlers"
)

func main() {

	target1Handler := &handlers.PortNumHandler{PortNum: "8081"}
	http.HandleFunc("/target1", target1Handler.Handler)

	target2Handler := &handlers.PortNumHandler{PortNum: "8082"}
	http.HandleFunc("/target2", target2Handler.Handler)

	fmt.Println("Goginx Running")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
