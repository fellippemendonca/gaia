package iugu_customers

import (
	log "github.com/sirupsen/logrus"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	"github.com/fellippemendonca/gaia/services"
)

// List model method
func List(svc *services.Services) ([]*iugu.Customer, error) {
	query := "SELECT id FROM iugu_customers;"
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	customers := make([]*iugu.Customer, 0)
	for rows.Next() {
		customer := new(iugu.Customer)
		err := rows.Scan(&customer.ID)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
