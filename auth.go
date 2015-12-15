package main

import (
	"fmt"
	"github.com/andrewburian/crypter"
	"io"
	"net/http"
)

const (
	BAD_REQUEST  int = 400
	NOT_FOUND    int = 404
	UNAUTHORIZED int = 401
	ERROR        int = 500
)

type authHandler struct {
	crypto *crypter.Crypter
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got auth request")
	session := getSession(r, h.crypto)

	// parse the form fields
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Decoding error!")
		w.WriteHeader(BAD_REQUEST)
		return
	}

	if len(r.FormValue("password")) < 6 {
		w.WriteHeader(UNAUTHORIZED)
		io.WriteString(w, "Password must be at least 6 characters")
	}

	// TODO DB auth
	auth := true

	// failed to login
	if !auth {
		w.WriteHeader(UNAUTHORIZED)
		w.Write([]byte("Username or Password incorrect"))
		return
	}

	session["uname"] = r.FormValue("username")

	setSession(w, h.crypto, session)

	io.WriteString(w, "Welome")
	return
}
