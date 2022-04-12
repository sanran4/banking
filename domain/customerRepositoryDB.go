package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sanran4/banking/errs"
	"github.com/sanran4/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("error while querying database" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:letstalkdata@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	//row := d.client.QueryRow(customerSql, id)
	err := d.client.Get(&c, customerSql, id)

	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("error while scanning rowset" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}
