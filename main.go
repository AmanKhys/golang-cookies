package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	// Start a web server with two endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setCookieHandler)
	mux.HandleFunc("/get", getCookieHandler)

	log.Print("Listening... on port http://localhost:3000/")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal()
	}
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new cookie containing the string "Hello world!" and
	// some non-default attributes

	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "hello world!",
		Path:     "/", // ?
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true, //https only, in modern web browsers it also considers normal localhost http requests as secure
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// continaing the necessary cookie data

	http.SetCookie(w, &cookie)

	// Echo out the cookie value in the response body
	w.Write([]byte("cookie set!"))

}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is "exampleCookie"). If no matching cookie is found, this will
	// return a http.ErrNoCookie error. We check for this, and return a 400 Bad Request response to the client.

	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body
	w.Write([]byte(cookie.Value))
}
