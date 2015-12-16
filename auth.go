package main

import (
	"fmt"
	"github.com/andrewburian/crypter"
	"golang.org/x/crypto/bcrypt"
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
	db     *dbManager
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

	if len(r.FormValue("username")) == 0 {
		w.WriteHeader(UNAUTHORIZED)
		io.WriteString(w, "No username specified")
		return
	}

	if len(r.FormValue("password")) < 6 {
		w.WriteHeader(UNAUTHORIZED)
		io.WriteString(w, "Password must be at least 6 characters")
		return
	}

	// DB auth
	uid, err := h.db.getUserId(r.FormValue("username"))
	if err != nil {
		w.WriteHeader(UNAUTHORIZED)
		//io.WriteString(w, "Username or Password incorrect")
		io.WriteString(w, "User not found")
		fmt.Println("User not found")
		return
	}

	hash, err := h.db.getUserHash(uid)
	if err != nil {
		w.WriteHeader(ERROR)
		io.WriteString(w, "Error retrieving password hash")
		fmt.Println("Hash error")
		return
	}

	// password check
	err = bcrypt.CompareHashAndPassword(hash, []byte(r.FormValue("password")))
	if err != nil {
		w.WriteHeader(UNAUTHORIZED)
		io.WriteString(w, "Username or Password incorrect")
		fmt.Println("Incorrect password")
		return
	}

	// setup session
	session["uid"] = string(uid)

	setSession(w, h.crypto, session)

	io.WriteString(w, "Welome")
	return
}
