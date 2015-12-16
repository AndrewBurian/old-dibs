package main

import (
	"fmt"
	"github.com/andrewburian/crypter"
	"io"
	"net/http"
	"os"
)

type loginHandler struct {
	crypto *crypter.Crypter
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("-- Login handler --")
	defer fmt.Println("-- end --\n")

	session := getSession(r, h.crypto)

	// check to see if user is already logged in
	if _, ok := session["uid"]; ok {
		fmt.Println("Redirecting logged in user")
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
		return
	}

	file, err := os.Open("views/login.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(w, file)

}

type logoutHandler struct {
}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("-- Logout handler --")
	defer fmt.Println("-- end --\n")

	// destory the session cookie by overwriting it with a delete cookie
	cookie, err := r.Cookie("sess")

	// if no cookie was found, just redirect
	if err != nil {
		fmt.Println("User not logged in")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("Setting delete cookie")
	cookie.MaxAge = -1

	w.Header().Set("Set-Cookie", cookie.String())

	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)

}
