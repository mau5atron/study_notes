package main 

import(
	"fmt"
	// net/http package from Go standard library to enable web server stuff
	"net/http"
)

// Creating an HTTP Server in Go

// http.ResponseWriter assembles the servers response and writes to the client
// http.Request is the clients request

// function that handles response and requests
func handler(w http.ResponseWriter, r *http.Request){
	// This writes to the client, will show on a page in localhost
	fmt.Fprintf(w, "Hello world\n") 
	// in terminal type: go run create_http_server.go
		// THEN
	// visit http://localhost:8080/ in your browser 
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Earth\n")
	// in terminal type: go run create_http_server.go
		// THEN
	// visit http://localhost:8080/earth in your browser
}

func main(){
	// This calls for the function handler 's output to match the directory /
	http.HandleFunc("/", handler)

	// This calls for the function handler2 output to match directory /earth
	http.HandleFunc("/earth", handler2)

	// listen to port 8080 and handle requests
	http.ListenAndServe(":8080", nil)
}