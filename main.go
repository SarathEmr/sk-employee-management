package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"sk-em/app"
	"sk-em/data"
)

const (
	dbUser     = "postgres"
	dbPassword = "password"
	dbHost     = "localhost"
	dbPort     = "5432"
	dbName     = "postgres"
)

func main() {

	server, err := setupServer()
	if err != nil {
		log.Fatalf("setupServer failed: %v", err)
	}

	r := server.InitRouter()
	log.Print("starting api server")
	http.ListenAndServe(":3002", r)
}

func setupServer() (*app.Server, error) {

	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, errors.Wrap(err, "create database connection failed")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "database ping failed")
	}
	log.Print("database connection successful")

	employeeRepo := data.NewEmployeeRepo(db)
	server := app.Server{EmployeeRepo: employeeRepo}
	err = server.EmployeeRepo.CreateEmployeeTable(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "EmployeeRepo.CreateEmployeeTable failed")
	}

	return &server, nil
}
