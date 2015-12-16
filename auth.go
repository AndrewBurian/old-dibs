package main

import (
	"fmt"
	"github.com/andrewburian/crypter"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type authHandler struct {
	crypto *crypter.Crypter
	db     *dbManager
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- Auth handler --")
	defer fmt.Println("-- end --\n")

	session := getSession(r, h.crypto)

	// parse the form fields
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Decoding error!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(r.FormValue("username")) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "No username specified")
		return
	}

	if len(r.FormValue("password")) < 6 {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Password must be at least 6 characters")
		return
	}

	// DB auth
	uid, err := h.db.getUserId(r.FormValue("username"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "User not found")
		fmt.Println("User not found")
		fmt.Println(err.Error())
		return
	}

	hash, err := h.db.getUserHash(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error retrieving password hash")
		fmt.Println("Hash error")
		fmt.Println(err.Error())
		return
	}

	// password check
	err = bcrypt.CompareHashAndPassword(hash, []byte(r.FormValue("password")))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Username or Password incorrect")
		fmt.Println("Incorrect password")
		return
	}

	fmt.Printf("Successfully authed user %s\n", r.FormValue("username"))

	// setup session
	session["uid"] = string(uid)

	setSession(w, h.crypto, session)

	io.WriteString(w, "Welome")
	return
}
