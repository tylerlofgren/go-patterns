package main

import (
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db pg.DB

func main() {
	//Setup database
	//db, err := ConnectDB()
	//if err != nil {
	//	log.Fatal("Could not connect to Postgres")
	//}
	//defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/messages", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("You must be Tyler eh?"))
		writer.WriteHeader(http.StatusOK)
	})
	//r.HandleFunc("/messages", postMessages).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

//func ConnectDB() (*pg.DB, error) {
//	opts := pg.Options{
//		TLSConfig:    nil,
//		Addr:         "localhost:5432",
//		User:         "",
//		Password:     "",
//		Database:     "",
//		PoolSize:     1,
//		MinIdleConns: 1,
//	}
//
//	db := pg.Connect(&opts)
//	_, err := db.Exec("SELECT 1")
//	return db, err
//}
//
//type Message struct {
//	Timestamp   uint64
//	Symbol      string
//	Volume      uint16
//	Temperature uint16
//}
//
//func postMessages(w http.ResponseWriter, r *http.Request) {
//	//Check content type
//	if r.Header.Get("Content-Type") != "application/json" {
//		msg := "Content-Type header is not application/json"
//		http.Error(w, msg, http.StatusUnsupportedMediaType)
//		return
//
//	}
//
//	//Unmarshal JSON
//	var m Message
//	err := json.NewDecoder(r.Body).Decode(&m)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	//Save
//	_, err = saveMessage(m)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	err = json.NewEncoder(w).Encode(m)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//}
//
//func saveMessage(m Message) (Message, error) {
//	tx, err := db.Begin()
//	defer func() {
//		err = tx.Rollback()
//		if err != nil {
//			log.Println(err.Error())
//		}
//	}()
//	if err != nil {
//		return Message{}, err
//	}
//
//	_, err = tx.Model(&m).Returning("*").Insert()
//	if err != nil {
//		return Message{}, err
//	}
//
//	return m, tx.Commit()
//}
