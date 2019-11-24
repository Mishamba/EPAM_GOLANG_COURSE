package tests

import (
	"errors"
	"sort"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/sorting/types"
)

func Errors() error {
	var somePersons types.People = types.People{
		types.Person{"g", "g", types.StartDate.AddDate(1990, 1, 1)},
		types.Person{"a", "a", types.StartDate.AddDate(1998, 1, 1)},
		types.Person{"z", "z", types.StartDate.AddDate(2000, 1, 1)},
		types.Person{"b", "b", types.StartDate.AddDate(1998, 1, 1)},
	}
	var sortPersons types.People = types.People{somePersons[2], somePersons[1], somePersons[3], somePersons[0]}
	sort.Sort(somePersons)
	if equalSlice(somePersons, sortPersons) {
		return errors.New("sort works wrong")
	}

	return nil
}

func equalSlice(fSlice types.People, sSlice types.People) bool {
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
