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

func main() {

	readFiles()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/profile", getProfile).Methods("GET")
	router.HandleFunc("/pastmeetings", getPastMeetings).Methods("GET")
	router.HandleFunc("/upcoming", getUpcoming).Methods("GET")
	router.HandleFunc("/schedule", schedule).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func readFiles(){
	file, err := os.Open("../../resources/blkEmployees.json")
	if err != nil {
        log.Println(err)
	}
	var blkEmployee models.Employees
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &blkEmployee)
	for i := 0; i < len(blkEmployee.Employee); i++ {
		fmt.Println("User Type: " + blkEmployee.Employee[i].EmployeeId)
}
		
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
