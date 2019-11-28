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
	sortPersons := NewPeople(somePersons[2], somePersons[1], somePersons[3], somePersons[0])
	sort.Sort(somePersons)
	key := equalSlice(sortPersons, somePersons)
	assert.True(t, key, "sort works wrong way")
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
