package iugu_customers

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	"github.com/fellippemendonca/gaia/services"
)

// GetByID model method
func GetByID(svc *services.Services, id string) (*iugu.Customer, error) {
	query := `SELECT id, email, name, notes, created_at, updated_at, cc_emails, phone, phone_prefix, zip_code, cpf_cnpj,
		district, state, city, street, number, complement, default_payment_method_id
	 	FROM iugu_customers where id = $1;`
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	customer := new(iugu.Customer)
	row := stmt.QueryRow(id)
	err = row.Scan(
		&customer.ID,
		&customer.Email,
		&customer.Name,
		&customer.Notes,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CcEmails,
		&customer.Phone,
		&customer.PhonePrefix,
		&customer.ZipCode,
		&customer.CpfCnpj,
		&customer.District,
		&customer.State,
		&customer.City,
		&customer.Street,
		&customer.Number,
		&customer.Complement,
		&customer.DefaultPaymentMethodID)
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		return customer, nil
	default:
		panic(err)
	}
	return nil, nil
}
