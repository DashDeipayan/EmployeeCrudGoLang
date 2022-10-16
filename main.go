package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	EmployeeId string `json:"employeeid"`
	EmailId    string `json:"-"`
}

var employees []Employee

func main() {
	fmt.Println("Main function runs")

	employees = append(employees, Employee{FirstName: "Deipayan", LastName: "Dash", EmployeeId: "1", EmailId: "deipayan@example.com"})
	employees = append(employees, Employee{FirstName: "Gareth", LastName: "Keenan", EmployeeId: "2", EmailId: "gareth@office.com"})
	employees = append(employees, Employee{FirstName: "Tim", LastName: "Canterbury", EmployeeId: "3", EmailId: "tim@office.com"})

	r := mux.NewRouter()

	r.HandleFunc("/", createHomepage).Methods("GET")
	r.HandleFunc("/employees", getAllemployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getStudent).Methods("GET")
	r.HandleFunc("/employees", createStudent).Methods("POST")
	r.HandleFunc("/employees/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4300", r))
}

func createHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Go Lang Custom Service</h1>"))
}

func (employee *Employee) isEmpty() bool {
	return employee.FirstName == "" && employee.LastName == "" && employee.EmployeeId == "" && employee.EmailId == ""
}

func getAllemployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all employees")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all employees")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	var student Employee
	for _, stud := range employees {
		if stud.EmployeeId == params["id"] {
			student = stud
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode("No student found")
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all employees")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some id")
		return
	}
	var student Employee
	_ = json.NewDecoder(r.Body).Decode(&student)

	if student.isEmpty() {
		json.NewEncoder(w).Encode("Please send some valid data")
	}

	for _, stud := range employees {
		if stud.EmployeeId == student.EmployeeId {
			json.NewEncoder(w).Encode("Roll Number already exists")
			return
		}
	}

	employees = append(employees, student)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all employees")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some id")
		return
	}
	var student Employee
	_ = json.NewDecoder(r.Body).Decode(&student)

	if student.isEmpty() {
		json.NewEncoder(w).Encode("Please send some valid data")
	}

	params := mux.Vars(r)
	fmt.Println(params)
	for i, stud := range employees {
		if stud.EmployeeId == params["id"] {
			employees = append(employees[:i], employees[i+1:]...)
			employees = append(employees, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode("Please send some valid data")
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all employees")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	for i, stud := range employees {
		if stud.EmployeeId == params["id"] {
			employees = append(employees[:i], employees[i+1:]...)
			json.NewEncoder(w).Encode(stud)
			return
		}
	}
	json.NewEncoder(w).Encode("No Such student found")
}
