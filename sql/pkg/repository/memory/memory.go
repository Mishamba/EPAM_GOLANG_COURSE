package memory

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/sql/pkg/model"
)

func NewContactsRepositoryInMemory(r *sql.DB) {
	connStr := "user=mishamba password=123 dbname=some-postgres sslmode=disable" //TODO
	r, _ = sql.Open("postgres", connStr)
}

func Save(r *sql.DB, contact model.Contact) (model.Contact, error) {
	rows, err := r.Query("SELECT (phone, email) FROM Contacts")
	if err != nil {
		return contact, err
	}

	defer rows.Close()
	var slice []model.Contact
	for rows.Next() {
		var tmp model.Contact
		err := rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
		if err != nil {
			return contact, err
		}

		slice = append(slice, tmp)
	}

	for _, c := range slice {
		if c.Email == contact.Email {
			return model.Contact{}, fmt.Errorf("contact with email %q already exists", c.Email)
		}

		if c.Phone == contact.Phone {
			return model.Contact{}, fmt.Errorf("contact with phone %q already exists", c.Phone)
		}
	}

	_, err = r.Exec("INSERT INTO Contacts (first_name, last_name, phone, email) VALUES ($1, $2, $3, $4, $5)", contact.FirstName, contact.LastName, contact.Phone, contact.Email)

	return contact, nil
}

func ListAll(r *sql.DB) ([]model.Contact, error) {
	var result []model.Contact
	rows, err := r.Query("SELECT (ID, first_name, last_name, phone, email) FROM Contacts")
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var tmp model.Contact
		err := rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
		if err != nil {
			return []model.Contact{}, err
		}

		result = append(result, tmp)
	}

	return result, nil
}

func GetByID(r *sql.DB, id uint) (model.Contact, error) {
	var contact model.Contact
	rows, err := r.Query("SELECT (ID, first_name, last_name, phone, email) FROM Contacts WHERE ID = ?", id)
	if err != nil {
		return model.Contact{}, err
	}

	defer rows.Close()
	err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return model.Contact{}, err
	}

	return contact, nil
}

func GetByPhone(r *sql.DB, phone string) (model.Contact, error) {
	var contact model.Contact
	rows, err := r.Query("SELECT (ID, first_name, last_name, phone, email) FROM Contacts WHERE phone = ?", phone)
	if err != nil {
		return model.Contact{}, err
	}

	defer rows.Close()
	err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return model.Contact{}, err
	}

	return contact, nil
}

func GetByEmail(r *sql.DB, email string) (model.Contact, error) {
	var contact model.Contact
	rows, err := r.Query("SELECT (ID, FirstName, LastName, phone, Email) FROM Contact WHERE Email = ?", email)
	if err != nil {
		return contact, err
	}

	defer rows.Close()
	err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return model.Contact{}, err
	}

	return model.Contact{}, fmt.Errorf("record not found")
}

func SearchByName(r *sql.DB, name string) ([]model.Contact, error) {
	var contacts []model.Contact
	rows, err := r.Query("SELECT (ID, first_name, last_name, phone, email) FROM Contacts WHERE first_name = ?", name)
	if err != nil {
		return contacts, err
	}

	defer rows.Close()
	for rows.Next() {
		var tmp model.Contact
		err := rows.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
		if err != nil {
			return []model.Contact{}, err
		}

		contacts = append(contacts, tmp)
	}
	return contacts, nil
}

func Delete(r *sql.DB, id uint) error {
	info, err := r.Exec("DELETE FROM Contacts WHERE ID = ?", id)
	if err != nil {
		return err
	}

	if count, _ := info.RowsAffected(); count > 1 {
		return errors.New("deleted " + string(count) + " strings, but expected one")
	}

	return nil
}
