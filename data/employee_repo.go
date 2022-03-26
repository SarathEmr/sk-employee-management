package data

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type EmployeeRepo interface {
	CreateEmployeeTable(ctx context.Context) error
	ListEmployees(ctx context.Context) ([]Employee, error)
	// AddEmployee () error
	// EditEmployee (Employee) error
	// DeleteEmployee (ID int64) error
}

type employeeRepoImpl struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) EmployeeRepo {
	return &employeeRepoImpl{db}
}

const (
	queryCreateEmployeeTable = `
		CREATE TABLE IF NOT EXISTS employee(
			id 				SERIAL PRIMARY KEY,
			name 			TEXT NOT NULL,
			age 			INT NOT NULL,
			address 		TEXT NOT NULL,
			designation 	TEXT NOT NULL, 
			joining_date 	DATE NOT NULL
		)
	;`
	queryListEmployees = `
		SELECT id, name, age, address, designation, joining_date
		FROM employee
	;`
)

func (impl employeeRepoImpl) CreateEmployeeTable(ctx context.Context) error {

	_, err := impl.db.ExecContext(ctx, queryCreateEmployeeTable)
	if err != nil {
		return errors.Wrap(err, "queryCreateEmployeeTable failed")
	}
	return nil
}

func (impl employeeRepoImpl) ListEmployees(ctx context.Context) ([]Employee, error) {

	rows, err := impl.db.QueryContext(ctx, queryListEmployees)
	if err != nil {
		return nil, errors.Wrap(err, "queryListEmployees failed")
	}
	defer rows.Close()
	var employees []Employee
	var emp Employee

	for rows.Next() {
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Address, &emp.Designation,
			&emp.JoiningDate)
		if err != nil {
			return nil, errors.Wrap(err, "queryListEmployees scan failed")
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
