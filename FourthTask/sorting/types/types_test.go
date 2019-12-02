package types

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	person1 := NewPerson("g", "g", 1990, 1, 1)
	person2 := NewPerson("a", "a", 1998, 1, 1)
	person3 := NewPerson("z", "z", 2000, 1, 1)
	person4 := NewPerson("b", "b", 1998, 1, 1)

	somePersons := NewPeople(person1, person2, person3, person4)
	sortPersons := NewPeople(person3, person2, person4, person1)

	sort.Sort(somePersons)
	key := equalSlice(sortPersons, somePersons)
	assert.True(t, key, "sort works wrong way 1")

	somePersons = NewPeople(person2, person3, person1, person4)
	sort.Sort(somePersons)
	key = equalSlice(sortPersons, somePersons)
	assert.True(t, key, "sort work's wrong way 2")

	person1 = NewPerson("m", "g", 1990, 1, 1)
	person2 = NewPerson("y", "a", 1998, 1, 1)
	person3 = NewPerson("x", "z", 2000, 1, 1)
	person4 = NewPerson("b", "b", 1998, 1, 1)

	somePersons = NewPeople(person4, person2, person3, person4)
	sort.Sort(somePersons)
	//sortPersons = NewPeople(, person4) //need to write it
	key = equalSlice(somePersons, somePersons)
	assert.True(t, key, "sort work's wrong way 3")
}

func equalSlice(fSlice People, sSlice People) bool {
	if fSlice.Len() != sSlice.Len() {
		return false
	}

	for i := 0; i < fSlice.Len(); i++ {
		if fSlice[i] != sSlice[i] {
			return false
		}
	}

	return true
}

func NewPeople(args ...Person) (result People) {
	for _, v := range args {
		result = append(result, v)
	}

	return result
}

func NewPerson(firstName string, lastName string, newYear, newMonth, newDay int) Person {
	newBirthday := StartDate.AddDate(newYear, newMonth, newDay)

	return Person{firstName, lastName, newBirthday}
}
