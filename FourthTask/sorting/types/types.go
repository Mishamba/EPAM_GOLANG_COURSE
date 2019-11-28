//	Package contains person type and implements sort.Interface
package types

import (
	"time"
)

var StartDate time.Time = time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC) //user will set data relatively Jesus birthday

type Person struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	result := p[i].Birthday.Before(p[j].Birthday)
	return result
}

func (p People) Swap(i, j int) {
	var tmp Person
	tmp = p[i]
	p[i] = p[j]
	p[j] = tmp
}

func NewPeople(args ...Person) (result People) { //this fucn was created for tests
	for _, v := range args {
		result = append(result, v)
	}

	return result
}

func NewPerson(firstName string, lastName string, newYear, newMonth, newDay int) { //this func was created for tests
	newBirthday := StartDate.AddDate(newYear, newMonth, newDay)

	return Person{firstName, lastName, newBirthday}
}
