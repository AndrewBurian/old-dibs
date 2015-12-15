package main

import (
//	"encoding/json"
	"net/http"
	"fmt"
//	"os"
//	"io"
)

const (
	BAD_REQUEST int = 400
	NOT_FOUND int = 404
	UNAUTHORIZED int = 401
)

type auth struct {
	name string `json:"username"`
	pass string `json:"password"`
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got auth request")

	// parse the form fields
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Decoding error!")
		w.WriteHeader(BAD_REQUEST)
		return
	}

	if r.PostFormValue("username") == "admin" && r.PostFormValue("password") == "pass" {
		fmt.Println("Welcome admin!")
		w.Write([]byte(`{"authentication":"success"}`))
	} else {
		w.WriteHeader(UNAUTHORIZED)
		fmt.Println(r.PostForm)
	}

	return
}

