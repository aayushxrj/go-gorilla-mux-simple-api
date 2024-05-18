package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)


func router () *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)	
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/albums", postAlbums).Methods("POST")
	router.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
	router.HandleFunc("/albums/{id}", deleteAlbumByID).Methods("DELETE")
	router.HandleFunc("/albums/{id}", updateAlbumByID).Methods("PUT")

	return router
}

func Test_rootHandler(t *testing.T){
	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func Test_getAlbums(t *testing.T){
	req, _ := http.NewRequest("GET", "/albums", nil)

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusFound, w.Code)
	}
}

func Test_postAlbums(t *testing.T){
	newAlbum := album{ID: "4", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99}
	jsonValue,_ := json.Marshal(newAlbum)

	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}

func Test_getAlbumByID(t *testing.T){
	req, _ := http.NewRequest("GET", "/albums/1", nil)

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusFound, w.Code)
	}
}

func Test_deleteAlbumByID(t *testing.T){
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, w.Code)
	}
}

func Test_updateAlbumByID(t *testing.T){
	newAlbum := album{ID: "1", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99}
	jsonValue,_ := json.Marshal(newAlbum)
	
	req, _ := http.NewRequest("PUT", "/albums/1", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}