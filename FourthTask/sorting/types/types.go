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
	diffDays := p[j].Birthday.Sub(p[i].Birthday) / 24

	if diffDays > 0 {
		return false
	}
	if diffDays < 0 {
		return true
	}
	if p[j].LastName < p[i].LastName {
		return false
	}
	if p[j].FirstName < p[i].FirstName {
		return false
	}
	return true
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
