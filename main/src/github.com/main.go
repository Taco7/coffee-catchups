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
	var blkEmployee models.Employees
	readFiles(&blkEmployee,"../../resources/blkEmployees.json")
		fmt.Println("User Type: " + blkEmployee.Employee[0].EmployeeId)

	fmt.Println("reading efront employees")
	var efrontEmployee models.Employees
	readFiles(&efrontEmployee,"../../resources/efrontEmployees.json")
		fmt.Println("User Type: " + efrontEmployee.Employee[0].EmployeeId)

	fmt.Println("reading meeting")
	var meeting models.Meeting
	readFiles(&meeting,"../../resources/meeting.json")
		fmt.Println("User Type: " + meeting.MeetingInfo[0].EmployeeId)
	
}

func readFiles(fileObject interface{}, filePath string) interface{} {
	file, err := os.Open(filePath)
	if err != nil {
        log.Println(err)
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
