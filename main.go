package main

import (
	_ "bytes"
	"database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type room struct {
	Name  string `json:"name"`
	Queue queue  `json:"queue"`
}

type queue struct {
	TurnList []turn `json:"turnList"`
}

type turn struct {
	Attendee attendee `json:"attendee"`
}

type attendee struct {
	Name string `json:"name"`
}

//var rooms []room

var rooms map[string]*room
var db *sql.DB

// Display all from the rooms var
func DisplayRooms(w http.ResponseWriter, r *http.Request) {
	// Get latest data from DB
	GetRoomsDB(db, rooms)

	var room_array []room

	for _, m := range rooms {
		room_array = append(room_array, *m)
	}
	roomlist := make(map[string]interface{})
	roomlist["roomList"] = room_array
	json.NewEncoder(w).Encode(roomlist)
}

// Display a single data
func DisplayRoom(w http.ResponseWriter, r *http.Request) {
	// Get latest data from DB
	GetRoomsDB(db, rooms)

	params := mux.Vars(r)
	if room, ok := rooms[params["name"]]; ok {
		json.NewEncoder(w).Encode(room)
		return
	}
	json.NewEncoder(w).Encode(&room{})
}

func EnQueue(w http.ResponseWriter, r *http.Request) {
	type attendee_id struct {
		Attendee_id string `json:"attendee_id"`
	}

	var atid attendee_id
	params := mux.Vars(r)
	room_name := params["name"]
	_ = json.NewDecoder(r.Body).Decode(&atid)

	EnQueueDB(db, room_name, atid.Attendee_id)
	json.NewEncoder(w).Encode("Ok")
}

func DeQueue(w http.ResponseWriter, r *http.Request) {
	type attendee_id struct {
		Attendee_id string `json:"attendee_id"`
	}

	var atid attendee_id
	params := mux.Vars(r)
	room_name := params["name"]
	_ = json.NewDecoder(r.Body).Decode(&atid)

	DeQueueDB(db, room_name, atid.Attendee_id)
	json.NewEncoder(w).Encode("Ok")
}

func EmptyQueue(w http.ResponseWriter, r *http.Request) {
	var tlist []turn
	params := mux.Vars(r)
	room_name := params["name"]
	if room, ok := rooms[room_name]; ok {
		tlist = room.Queue.TurnList
		room.Queue.TurnList = room.Queue.TurnList[:0]
	}
	EmptyQueueDB(db, room_name)
	json.NewEncoder(w).Encode(tlist)
}

// main function to boot up everything
func main() {
	db = ConnectDB()
	router := mux.NewRouter()
	rooms = make(map[string]*room)

	router.HandleFunc("/room", DisplayRooms).Methods("GET")
	router.HandleFunc("/room/{name}", DisplayRoom).Methods("GET")
	router.HandleFunc("/room/{name}/queue", EnQueue).Methods("POST")
	router.HandleFunc("/room/{name}/queue", DeQueue).Methods("DELETE")
	router.HandleFunc("/room/{name}", EmptyQueue).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
	db.Close()
}
