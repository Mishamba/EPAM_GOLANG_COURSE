//it works, but i have one problem there
package construct

import (
	"fmt"
	"sync"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/sorting/types"
)

var wg sync.WaitGroup

func People() (res types.People) { //i'll left some questions
	cont := make(chan string)
	c := make(chan int) //chanel for data exchange
	wg.Add(1)
	go func(c chan int) { //here user will choose what to do(exit or enter new user)
		for {
			fmt.Print("\n1 - new person\n2 - show result\n")
			var tmp int
			_, err := fmt.Scan(&tmp)
			if err != nil {
				panic(err)
			}
			c <- tmp
			<-cont
		}
	}(c)

	go func(c chan int) { //here user will create new Person
		defer fmt.Println("u entered persons and now u will see result")
		defer wg.Done()
		for {
			if ans := <-c; ans == 2 { //when this condition works it still works till
				return
			}

			var newFirstName, newLastName string
			var newYear, newMonth, newDay int

			fmt.Println("Enter Person")
			fmt.Print("First and Last name: ") //here. i don't know why it happens(i can't deug my code, don't know, how to create config for goland)

			if _, err := fmt.Scan(&newFirstName, &newLastName); err != nil {
				panic(err)
			}

			fmt.Print("Birtday year, month, day: ")

			if _, err := fmt.Scan(&newYear, &newMonth, &newDay); err != nil {
				panic(err)
			}

			newBirthDay := types.StartDate.AddDate(newYear, newMonth, newDay) //creating Persons birthday

			res = append(res, types.Person{newFirstName, newLastName, newBirthDay})
			cont <- "can go"
		}
	}(c)
	wg.Wait()
	return res
}
