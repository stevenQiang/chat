package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"chat/db"
)

var addr = flag.String("port", ":8080", "http service address")
var hub = newHub()

func indexPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	username := getUserName(r)
	if username == ""{
		http.Redirect(w, r, "/login", 302)
	}
	http.ServeFile(w, r, "home.html")
}

func wsPage(w http.ResponseWriter, r *http.Request) {
	serveWs(hub, w, r)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		http.ServeFile(w, r, "login.html")	
	}else{
		username := r.FormValue("username")
		password := r.FormValue("password")
		setSession(username, password, w)
		http.Redirect(w, r, "/", 302)
	}
}

var router = mux.NewRouter()

func main() {
	db := db.InitDB()
	defer db.Close()
	flag.Parse()
	go hub.run()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/login", loginPage)
	router.HandleFunc("/ws", wsPage)
	http.Handle("/", router)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
