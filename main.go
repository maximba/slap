package main

import (
	_ "bytes"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"database/sql"
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
	json.NewEncoder(w).Encode(rooms)
}

// Display a single data
func DisplayRoom(w http.ResponseWriter, r *http.Request) {
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

	var att attendee
	params := mux.Vars(r)
	room_name := params["name"]
	if room, ok := rooms[room_name]; ok {
		var atid attendee_id
		_ = json.NewDecoder(r.Body).Decode(&atid)
		if atid.Attendee_id != "" {
			att = attendee{Name: atid.Attendee_id}
			room.Queue.TurnList = append(room.Queue.TurnList, turn{att})
		}
	}
	EnQueueDB(db, room_name, att.Name)
	json.NewEncoder(w).Encode(att)
}

func DeQueue(w http.ResponseWriter, r *http.Request) {
	var att attendee
	params := mux.Vars(r)
	room_name :=params["name"]
	if room, ok := rooms[room_name]; ok {
		att = room.Queue.TurnList[0].Attendee
		room.Queue.TurnList = room.Queue.TurnList[1:]
	}
	DeQueueDB(db, room_name, att.Name)
	json.NewEncoder(w).Encode(att)
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
	router := mux.NewRouter()
	rooms = make(map[string]*room)
	/*    rooms = map[string]*room {
	        "Monkey Island": &room{Name: "Monkey Island", Queue: queue{TurnList: []turn{}}},
	        "Gotham":        &room{Name: "Gotham",        Queue: queue{TurnList: []turn{}}},
	        "New New York":  &room{Name: "New New York",  Queue: queue{TurnList: []turn{}}},
	      }
	*/
	//  "New New York": &room{Name: "New New York", Queue: queue{TurnList: []turn{{Attendee: attendee{Name: "Maxi"}}, {Attendee: attendee{Name: "Manu"}}}}},

	router.HandleFunc("/room", DisplayRooms).Methods("GET")
	router.HandleFunc("/room/{name}", DisplayRoom).Methods("GET")
	router.HandleFunc("/room/{name}/queue", EnQueue).Methods("POST")
	router.HandleFunc("/room/{name}/queue", DeQueue).Methods("DELETE")
	router.HandleFunc("/room/{name}", EmptyQueue).Methods("DELETE")
	db = ConnectDB()
	defer db.Close()
	GetRoomsDB(db, rooms)
	log.Fatal(http.ListenAndServe(":8080", router))
}
