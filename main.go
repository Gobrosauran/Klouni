package main

import (
	"fmt"
	"klouni/usprob"
)

type User struct {
	Name   string
	Rating float64
}

func (u User) NameUSer(name string) {
	fmt.Println("Меня зовут:", u.Name)
}

func (u User) Greeting(rating float64) {
	if u.Rating-rating < 10.0 {
		fmt.Println("Мой рейтинг:", u.Rating)
		fmt.Println(u.Rating)
		return
	} else if u.Rating+rating == 10.0 {
		fmt.Println("Мой рейтинг:", u.Rating)
		return

	}
}

func main() {
	user := User{
		Name:   "Sauran",
		Rating: 10.0,
	}
	fmt.Println("user:", user)
	user.NameUSer(user.Name)
	user.Greeting(user.Rating)

	usprob.UserGo(user.Name)

}
