package iugu_customers

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	iugu "github.com/fellippemendonca/gaia/lib/iugu"
	"github.com/fellippemendonca/gaia/services"
)

// Create model method
func Create(svc *services.Services, customerID string, customer *iugu.Customer) (string, error) {
	query := `INSERT INTO iugu_customers(
		id, iugu_id, email, name, notes, created_at, updated_at, cc_emails, phone_prefix,
		phone, zip_code, cpf_cnpj, district, state, city, street, number, complement, default_payment_method_id
	) VALUES(
		$1, $2, $3, $4, $5, $6, $7, $8, $9,
		$10, $11, $12, $13, $14, $15, $16, $17, $18, $19
	)`
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		customerID,
		customer.ID,
		customer.Email,
		customer.Name,
		customer.Notes,
		customer.CreatedAt,
		customer.UpdatedAt,
		customer.CcEmails,
		customer.PhonePrefix,
		customer.Phone,
		customer.ZipCode,
		customer.CpfCnpj,
		customer.District,
		customer.State,
		customer.City,
		customer.Street,
		customer.Number,
		customer.Complement,
		customer.DefaultPaymentMethodID,
	)
	if err != nil {
		log.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	log.Printf("row affected = %d\n", rowCnt)

	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		return "", nil
	default:
		panic(err)
	}
	return "", nil
}
