package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"io"

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
	router.HandleFunc("/profile/{employeeId}", getProfile).Methods("GET")
	router.HandleFunc("/pastmeetings/{employeeId}", getPastMeetings).Methods("GET")
	router.HandleFunc("/upcoming/{employeeId}", getUpcoming).Methods("GET")
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
	
	vars := mux.Vars(r)
	userName := vars["employeeId"]
	
	for _, v := range blkEmployee.Employee {
		if v.EmployeeId == userName {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write(b)
			return
		}
	}
	for _, v := range efrontEmployee.Employee {
		if v.EmployeeId == userName {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write(b)
			return
		}
	}
	w.Write([]byte("User not found"))
}
func getPastMeetings(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	userName := vars["employeeId"]
	for _, v := range meeting.MeetingInfo {
		if v.EmployeeId == userName {
			b, err := json.Marshal(v.SpokenTo)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write(b)
			return
		}
	}
	w.Write([]byte("Details not found"))
}
func getUpcoming(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	userName := vars["employeeId"]
	for _, v := range meeting.MeetingInfo {
		if v.EmployeeId == userName {
			b, err := json.Marshal(v.Upcoming)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write(b)
			return
		}
	}
	w.Write([]byte("Details not found"))
}
func schedule(w http.ResponseWriter, r *http.Request) {
	organization:="Efront"
	employeeId:="yghonge"
	name:="yagnesh Ghonge"
	Team:="data"
	aboutMe:="i love coffee"
	location:="NY"
	email:="yghonge@blk.com"

	var emp models.EmployeeDetails
	emp = models.EmployeeDetails{EmployeeId: employeeId, Name:name,Team: Team,About: aboutMe,Location: location, Email: email}
	

if(organization =="Blackrock"){
    f, erro := os.OpenFile("../../resources/blkEmployees.json", os.O_WRONLY, 0666)
    if erro != nil {
      fmt.Println(erro)
    }

    blkEmployee.Employee=append(blkEmployee.Employee, emp)
    result, error := json.Marshal(blkEmployee)
    if error != nil {
      fmt.Println(error)
    }

    n, err := io.WriteString(f, string(result))
    if err != nil {
      fmt.Println(n, err)
    }


    scheduleMeeting(emp.EmployeeId, efrontEmployee)
}else{
	f, erro := os.OpenFile("../../resources/efrontEmployees.json", os.O_WRONLY, 0666)
    if erro != nil {
      fmt.Println(erro)
    }

    efrontEmployee.Employee=append(efrontEmployee.Employee, emp)
    result, error := json.Marshal(efrontEmployee)
    if error != nil {
      fmt.Println(error)
    }

    n, err := io.WriteString(f, string(result))
    if err != nil {
      fmt.Println(n, err)
    }


    scheduleMeeting(emp.EmployeeId, efrontEmployee)
}


	
}


	func scheduleMeeting(employeeId string, Employees models.Employees){

	}
