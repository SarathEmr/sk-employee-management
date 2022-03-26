package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func (server Server) handleListEmployees(w http.ResponseWriter, r *http.Request) {

	employees, err := server.EmployeeRepo.ListEmployees(r.Context())
	if err != nil {
		log.Printf("handleListEmployees: EmployeeRepo.ListEmployees failed: %v", err)
		w.Write([]byte("error"))
		return
	}

	encoded, err := json.Marshal(employees)
	if err != nil {
		log.Printf("handleListEmployees: json.Marshal failed: %v", err)
		w.Write([]byte("error"))
		return
	}

	w.Write(encoded)
}
