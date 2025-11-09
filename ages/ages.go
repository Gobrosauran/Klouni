package ages

import "fmt"

func Userages(age int) {
	user1 := age
	defer func() {
		fmt.Println("User ages:", user1)
	}()

	
	if user1 < 18 {
		fmt.Println("User is not 18")
		return

	} else if user1 >= 18 {
		fmt.Println("User age is :", user1)

	}

}
