package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
     _"bytes"
    _"fmt"
)

type room struct {
    Name   string `json:"name"`
    Queue  queue  `json:"queue"`
}

type queue struct {
   TurnList []turn `json:"turnList"`
}

type turn struct {
    Attendee attendee `json:"attendee"`
}    

type attendee struct {
    Name  string `json:"name"`
}

var rooms []room

// Display all from the rooms var
func GetRooms(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(rooms)
}

// Display a single data
func GetRoom(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range rooms {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&room{})
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
    rooms = append(rooms, room{Name: "Monkey Island", Queue: queue{TurnList: []turn{{Attendee: attendee{Name: "Maxi"}}, {Attendee: attendee{Name: "Manu"}}}}})
    rooms = append(rooms, room{Name: "Gotham", Queue: queue{TurnList: []turn{{Attendee: attendee{Name: "Maxi"}}, {Attendee: attendee{Name: "Manu"}}}}})
    rooms = append(rooms, room{Name: "New New York", Queue: queue{TurnList: []turn{{Attendee: attendee{Name: "Maxi"}}, {Attendee: attendee{Name: "Manu"}}}}})
    router.HandleFunc("/room", GetRooms).Methods("GET")    
    router.HandleFunc("/room/{name}", GetRoom).Methods("GET")
//    router.HandleFunc("/room/{name}/queue", EnQueue).Methods("POST")
//    router.HandleFunc("/room/{name}/queue", DeQueue).Methods("DELETE")
//    router.HandleFunc("/room//queue", EmptyQueue).Methods("DELETE")    
    log.Fatal(http.ListenAndServe(":8000", router))
}
