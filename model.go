package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
)

func ConnectDB() *sql.DB {
	dbUser, ok := os.LookupEnv("DBUSER")
	if !ok {
		dbUser = "slapdb"
	}
	dbPass, ok := os.LookupEnv("DBPASS")
	if !ok {
		dbPass = "slapdb"
	}
	dbHost, ok := os.LookupEnv("DBHOST")
	if !ok {
		dbHost = "localhost"
	}
	dbName, ok := os.LookupEnv("DBNAME")
	if !ok {
		dbName = "slapdb"
	}
	dbPort, ok := os.LookupEnv("DBPORT")
	if !ok {
		dbPort = "3306"
	}
	connectURL := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", connectURL)
	if err != nil {
		os.Exit(1)
	}
	return db
}

func GetTurnListDB(db *sql.DB, room_id int) ([]turn){
	turnlist :=  []turn{}
	rows, err := db.Query("SELECT attendee FROM turn WHERE room_id=? ORDER BY priority ASC", room_id)
	if (err != nil) {
		log.Fatal(err)
	}	
	var attendee_id string
	var att attendee
	for rows.Next() {
               if err := rows.Scan(&attendee_id); err != nil {
                        log.Fatal(err)
                }
                att = attendee{Name: attendee_id}                
                turnlist = append(turnlist, turn{att})                
	}
	return turnlist
}

func GetRoomsDB(db *sql.DB, rooms map[string]*room) {
	rows, err := db.Query("SELECT id, name FROM rooms")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	var id int
	
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		turnlist := GetTurnListDB(db, id)
		rooms[name] = &room{Name: name, Queue: queue{TurnList: turnlist}}
	}
}

func EnQueueDB(db *sql.DB, room string, attendee string) {
	var room_id int
        err := db.QueryRow("SELECT id FROM rooms WHERE name=?", room).Scan(&room_id)
        if err != nil {
        	log.Fatal(err)
        }
        
	_,  err = db.Exec("INSERT INTO turn (room_id, attendee) VALUES (?,?)", room_id, attendee)
	if (err !=nil) {
		log.Fatal(err)
	}
} 

func DeQueueDB(db *sql.DB, room string, attendee string) {
	var room_id int
	
        err := db.QueryRow("SELECT id FROM rooms WHERE name=?", room).Scan(&room_id)
        if err != nil {
                log.Fatal(err)
        }
        _, err = db.Exec("DELETE FROM turn WHERE (room_id=? AND attendee=?)", room_id, attendee)
}

func EmptyQueueDB(db *sql.DB, room string) {
        var room_id int

        err := db.QueryRow("SELECT id FROM rooms WHERE name=?", room).Scan(&room_id)
        if err != nil {
                log.Fatal(err)
        }
        _, err = db.Exec("DELETE FROM turn WHERE room_id=?", room_id)
}
