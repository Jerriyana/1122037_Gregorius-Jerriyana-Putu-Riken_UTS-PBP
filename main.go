package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// 1. Routing Get All Rooms
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")

	// 2. Routing Insert Room
	router.HandleFunc("/join_room", controllers.InsertRoom).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8890")
	log.Println("Connected to port 8890")
	log.Fatal(http.ListenAndServe(":8890", router))
}
