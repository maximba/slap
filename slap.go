package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
     _"bytes"
    _"fmt"
)

type Attendee struct {
    Name  string `json:"name"`
}

type Room struct {
    Name   string `json:"name"`
    Queue  Queue  `json:"queue"`
}

type Queue struct {
   TurnList turnList `json:"turnList"`
}

type turnList []Attendee

var room []Room

// Display all from the room var
func GetRooms(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(room)
}

// Display a single data
func GetRoom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range room {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Room{})
}

/*  create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
} */

/* Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}*/

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    room = append(room, Room{Name: "Monkey Island", Queue: Queue{TurnList: []Attendee{{Name: "Perico"}, {Name: "Antoñico"}}}})
    room = append(room, Room{Name: "Gotham", Queue: Queue{TurnList: []Attendee{}}})    
    room = append(room, Room{Name: "New New York", Queue: Queue{TurnList: []Attendee{}}}) 
    router.HandleFunc("/room", GetRooms).Methods("GET")    
    router.HandleFunc("/room/{name}", GetRoom).Methods("GET")
//    router.HandleFunc("/room/{name}/queue", EnQueue).Methods("POST")
//    router.HandleFunc("/room/{name}/queue", DeQueue).Methods("DELETE")
//    router.HandleFunc("/room//queue", EmptyQueue).Methods("DELETE")    
    log.Fatal(http.ListenAndServe(":8000", router))
}
