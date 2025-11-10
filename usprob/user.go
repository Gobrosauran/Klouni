package usprob

import (
	"fmt"
)

func UserGo(str string) {
	user1 := str
	if user1 == "" {
		fmt.Println("Err!")

	} else {
		fmt.Println("User name:", user1)

	}
}
