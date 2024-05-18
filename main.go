package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func getAlbums(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(albums)
}

func postAlbums(w http.ResponseWriter, r *http.Request){
	var newAlbum album
	json.NewDecoder(r.Body).Decode(&newAlbum)
	albums = append(albums, newAlbum)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAlbum)
}

func getAlbumByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, album := range albums {
		if album.ID == params["id"] {
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
func deleteAlbumByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	for i, album := range albums {
		if album.ID == params["id"]{
			albums = append(albums[:i],  albums[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateAlbumByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	var newAlbum album
	json.NewDecoder(r.Body).Decode(&newAlbum)

	for i, album := range albums{
		if album.ID == params["id"]{
			albums[i] = newAlbum
			w.WriteHeader(http.StatusCreated)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)	
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/albums", postAlbums).Methods("POST")
	router.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
	router.HandleFunc("/albums/{id}", deleteAlbumByID).Methods("DELETE")
	router.HandleFunc("/albums/{id}", updateAlbumByID).Methods("PUT")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}