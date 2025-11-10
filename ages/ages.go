package ages

import "fmt"

func Userages(age int) {
	user1 := age
	defer func() {
		if user1 >= 18 {
			fmt.Println("User ages:", user1)
		} else {
			fmt.Println("User != 18!!!!")
			return
		}

	}()

	if user1 < 18 {
		fmt.Println("User is not 18")
		return

	} else if user1 >= 18 {
		fmt.Println("User Age is done!")

	}

}
