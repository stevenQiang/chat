package main

import (
	"net/http"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (username string){
	cookie, err := r.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil{
			username = cookieValue["username"]
		}
	}
	return username
}

func setSession(username string, password string, r http.ResponseWriter){
	value := map[string]string{
		"username": username,
		"password": hashPassword(password),
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil{
		cookie := &http.Cookie{
			Name: "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(r, cookie)
	}
}

func hashPassword(password string) string{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func resolvePassowrd(password, hashword string) bool{
	return bcrypt.CompareHashAndPassword([]byte(hashword), []byte(password)) == nil
}