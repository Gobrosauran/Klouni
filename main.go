package main

import (
	"fmt"
	"klouni/ages"
	"klouni/usprob"
)

type User struct {
	Name   string
	Age    int
	Rating float64
}

func (u User) NameUSer(name string) {
	usprob.UserGo(u.Name)
}

func (u User) Greeting(rating float64) {
	if u.Rating-rating < 10.0 {
		fmt.Println("User Rating:", u.Rating)
		return

	} else if u.Rating+rating == 10.0 {
		fmt.Println(":", u.Rating)
		return

	}
}

func main() {
	user := User{
		Name:   "Sauran",
		Age:    18,
		Rating: 9.9,
	}
	

	fmt.Println("user:", user)
	user.NameUSer(user.Name)

	fmt.Println("")
	user.Greeting(user.Rating)
	fmt.Println("")

	ages.Userages(user.Age)

}
