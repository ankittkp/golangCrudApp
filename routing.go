package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/getEmployee", GetEmployees).Methods("GET")
	r.HandleFunc("/getEmployeeById/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/updateEmployeeById/{eid}", UpdateEmployeeById).Methods("PUT")
	r.HandleFunc("/deleteEmployeeById/{eid}", DeleteEmployeeById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
