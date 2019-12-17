package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/sql/pkg/model"
	"github.com/Mishamba/EPAM_GOLANG_COURSE/sql/pkg/repository/memory"
)

const (
	CommandSave = iota + 1
	CommandListAll
	CommandGetByID
	CommandGetByPhone
	CommandGetByEmail
	CommandSearchByName
	CommandDelete
)

func main() {
	var r *sql.DB
	memory.NewContactsRepositoryInMemory(r)

	for {
		fmt.Print(menu)
		var command int
		if _, err := fmt.Scanf("%d", &command); err != nil {
			log.Println(err)
		}

		switch command {
		case CommandSave:
			if err := Save(r); err != nil {
				log.Println(err)
			}
		case CommandListAll:
			if err := ListAll(r); err != nil {
				log.Println(err)
			}
		case CommandGetByID:
			if err := GetByID(r); err != nil {
				log.Println(err)
			}
		case CommandGetByPhone:
			if err := GetByPhone(r); err != nil {
				log.Println(err)
			}
		case CommandGetByEmail:
			if err := GetByEmail(r); err != nil {
				log.Println(err)
			}
		case CommandSearchByName:
			if err := SearchByName(r); err != nil {
				log.Println(err)
			}
		case CommandDelete:
			if err := Delete(r); err != nil {
				log.Println(err)
			}
		default:
			log.Printf("command not foumd for value %d\n", command)
		}

		printSeparator()
	}
}

func ListAll(r *sql.DB) error {
	records, err := memory.ListAll(r)
	if err != nil {
		return fmt.Errorf("error in ListAll: %q", err.Error())
	}

	fmt.Println("ListAll:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

func GetByID(r *sql.DB) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	record, err := memory.GetByID(r, id)
	if err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Println("GetByID")
	fmt.Println(record)

	return nil
}

func GetByPhone(r *sql.DB) error {
	phone := readString("Please enter an 'Phone' field and press Enter")

	record, err := memory.GetByPhone(r, phone)
	if err != nil {
		return fmt.Errorf("error in GetByPhone: %q", err.Error())
	}

	fmt.Println("GetByPhone:")
	fmt.Println(record)

	return nil
}

func GetByEmail(r *sql.DB) error {
	email := readString("Please enter an 'Email' field and press Enter")

	record, err := memory.GetByEmail(r, email)
	if err != nil {
		return fmt.Errorf("error in GetByEmail: %q", err.Error())
	}

	fmt.Println("GetByEmail:")
	fmt.Println(record)

	return nil
}

func SearchByName(r *sql.DB) error {
	email := readString("Please enter prefix for 'Name' field and press Enter")

	records, err := memory.SearchByName(r, email)
	if err != nil {
		return fmt.Errorf("error in SearchByName: %q", err.Error())
	}

	fmt.Println("SearchByName:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

func Delete(r *sql.DB) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	if err := memory.Delete(r, id); err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Printf("Delete:\nRecord with ID %d successfylly deleted\n", id)
	return nil
}

func Save(r *sql.DB) error {
	contact := model.Contact{
		FirstName: readString("Please enter an 'FirstName' field and press Enter"),
		LastName:  readString("Please enter an 'LastName' field and press Enter"),

		Phone: readString("Please enter an 'Phone' field and press Enter"),
		Email: readString("Please enter an 'Email' field and press Enter"),
	}

	result, err := memory.Save(r, contact)
	if err != nil {
		return err
	}

	fmt.Println("Save Contact:")
	fmt.Println(result)

	return nil
}

const menu = `
Please enter operation number:
  * 1 - Save
  * 2 - ListAll
  * 3 - GetByID
  * 4 - GetByPhone
  * 5 - GetByEmail
  * 6 - SearchByName
  * 7 - Delete 
  * Control + C - to exit 
`

func readString(message string) string {
	var r string

	for r == "" {
		fmt.Println(message)
		if _, err := fmt.Scanf("%s", &r); err != nil {
			fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
			printSeparator()
		}

	}

	return r
}

func readUint(message string) uint {
	var r uint

	for {
		fmt.Println(message)
		_, err := fmt.Scanf("%d", &r)
		if err == nil {
			break
		}

		fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
		printSeparator()
	}

	return r
}

func printSeparator() {
	for i := 0; i < 50; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}
