package main

import (
	"io"
	"net/http"
)

func main() {

	// Let's work in the root path and let's send a greeting to it:
	http.HandleFunc("/", saludo)
	// Let's listen from some port of out computer, for example, the 8000:
	http.ListenAndServe(":8000", nil) // nil means that we won't return nothing when we execute this command
	// (example: in HandleFunc we returned "saludo").

}

func saludo(w http.ResponseWriter, peticion *http.Request) {

	// w implies that http.ResponseWriter will be an object.
	// http. specifies the type of the object and ResponseWriter send a response to the user.
	// The object peticion exists to manage the request of the user.

	io.WriteString(w, "Hello World!") // w allow us to write into the server.

}

// Windows will require administrator permissions to execute this script.

// It is recommended to start to build servers in go, to contibute some project as Hugo (gohugo.io) which allows you to create your own
// server from the current point in which Hugo has been developed (instead of starting from scratch).
