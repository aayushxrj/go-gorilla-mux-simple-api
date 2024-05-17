package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
}

func main(){
	http.HandleFunc("/", rootHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}