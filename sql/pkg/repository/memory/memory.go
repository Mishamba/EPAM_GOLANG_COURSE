package memory

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/sql/pkg/model"
)

type ContactsRepositoryInMemory struct {
	storage *sql.DB
	lastID  uint
}

func NewContactsRepositoryInMemory() *ContactsRepositoryInMemory {
	connStr := "user=mishamba password=123 dbname=some-postgres sslmode=disable" //TODO
	db, _ := sql.Open("postgres", connStr)
	return &ContactsRepositoryInMemory{
		storage: db,
	}
}

func (r *ContactsRepositoryInMemory) Save(contact model.Contact) (model.Contact, error) {
	row, err := r.storage.Query("select (Phone, Email) from Contacts")
	if err != nil {
		return contact, err
	}

	defer row.Close()
	var slice []model.Contact
	for row.Next() {
		var tmp model.Contact
		err := row.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
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

	r.lastID++
	contact.ID = r.lastID
	_, err = r.storage.Exec("insert into Contacts (ID, FirstName, LastName, Phone, Email) values ($1, $2, $3, $4, $5)", contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Email)

	return contact, nil
}

func (r *ContactsRepositoryInMemory) ListAll() ([]model.Contact, error) {
	var result []model.Contact
	row, err := r.storage.Query("select (ID, FirstName, LastName, Phone, Email) from Contacts")
	if err != nil {
		return result, err
	}

	defer row.Close()
	for row.Next() {
		var tmp model.Contact
		err := row.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
		if err != nil {
			return result, err
		}

		result = append(result, tmp)
	}

	return result, nil
}

func (r *ContactsRepositoryInMemory) GetByID(id uint) (model.Contact, error) {
	var contact model.Contact
	row, err := r.storage.Query("select (ID, FirstName, LastName, Phone, Email) from Contacts where ID = ?", id)
	if err != nil {
		return model.Contact{}, nil
	}

	defer row.Close()
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return contact, nil
	}

	return contact, nil
}

func (r *ContactsRepositoryInMemory) GetByPhone(phone string) (model.Contact, error) {
	var contact model.Contact
	row, err := r.storage.Query("select (ID, FirstName, LastName, Phone, Email) from Contacts where ID = ?", phone)
	if err != nil {
		return model.Contact{}, nil
	}

	defer row.Close()
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return contact, nil
	}

	return contact, nil
}

func (r *ContactsRepositoryInMemory) GetByEmail(email string) (model.Contact, error) {
	var contact model.Contact
	row, err := r.storage.Query("select (ID, FirstName, LastName, Phone, Email) from Contact where Email = ?", email)
	if err != nil {
		return contact, err
	}

	defer row.Close()
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
	if err != nil {
		return contact, nil
	}

	return model.Contact{}, fmt.Errorf("record not found")
}

func (r *ContactsRepositoryInMemory) SearchByName(name string) ([]model.Contact, error) {
	var contacts []model.Contact
	row, err := r.storage.Query("select (ID, FirstName, LastName, Phone, Email) from Contacts where Email = ?", name)
	if err != nil {
		return contacts, nil
	}

	defer row.Close()
	for row.Next() {
		var tmp model.Contact
		err := row.Scan(&tmp.ID, &tmp.FirstName, &tmp.LastName, &tmp.Phone, &tmp.Email)
		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, tmp)
	}
	return contacts, nil
}

func (r *ContactsRepositoryInMemory) Delete(id uint) error {
	info, err := r.storage.Exec("delete from Contacts where id = ?", id)
	if err != nil {
		return err
	}

	count, _ := info.RowsAffected()

	if count > 1 {
		return errors.New("deleted " + string(count) + " strings, but expected one")
	}

	return nil
}
