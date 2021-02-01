package domain

import (
	"database/sql"
	"time"

	"github.com/PrasadR287/banking/errs"
	"github.com/PrasadR287/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDb struct {
	client *sqlx.DB
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	customers := make([]Customer, 0)
	var err error
	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// rows, err = d.client.Query(findAllSQL)
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		// rows, err = d.client.Query(findAllSQL, status)
		err = d.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexceptedError("Unexpected database error")
	}

	// sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexceptedError("Unexpected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexceptedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil

}

func (d customerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSQL, id)
	var c Customer
	err := d.client.Get(&c, customerSQL, id)
	// err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexceptedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() customerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:shrish287@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return customerRepositoryDb{client}
}
