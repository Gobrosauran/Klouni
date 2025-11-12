package main

import (
	"fmt"
	"klouni/usprob"
)

type User struct {
	Name   string
	Age    int
	Rating float64
}

func NewUser(
	name string,
	age int,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if age < 0 || age > 150 {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}
	}
	return User{
		Name:   name,
		Age:    age,
		Rating: rating,
	}

}

func (u User) NameUSer(name string) {
	usprob.UserGo(u.Name)
}

func Greeting(u *User, rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("User Rating:", u.Rating)
		return

	} else {
		fmt.Println(u.Rating)
		return

	}
}

func main() {
	user := NewUser(
		"",
		160,
		0.0,
	)
	fmt.Println("User:", user)

}
