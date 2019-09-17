package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"io/ioutil"

	"github.com/gorilla/mux"
	"./models"
	"encoding/json"
)

var blkEmployee models.Employees
var efrontEmployee models.Employees
var meeting models.Meeting

func main() {
	loadData()
	configureRoutes()
}

func configureRoutes(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/profile", getProfile).Methods("GET")
	router.HandleFunc("/pastmeetings", getPastMeetings).Methods("GET")
	router.HandleFunc("/upcoming", getUpcoming).Methods("GET")
	router.HandleFunc("/schedule", schedule).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loadData(){
	fmt.Println("reading blk employees")
	readFiles(&blkEmployee,"../../resources/blkEmployees.json")

	fmt.Println("reading efront employees")	
	readFiles(&efrontEmployee,"../../resources/efrontEmployees.json")

	fmt.Println("reading meeting")
	readFiles(&meeting,"../../resources/meeting.json")
}

func readFiles(fileObject interface{}, filePath string) interface{} {
	file, err := os.Open(filePath)
	if err != nil {
        log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &fileObject)
	return fileObject
}
	

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func getPastMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func getUpcoming(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func schedule(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
