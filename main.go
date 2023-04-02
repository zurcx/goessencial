package main

import (
	"fmt"
	"log"

	"github.com/zurcx/goessencial/files"
)

var enabled bool
var name string = "Luiza"
var char rune
var age int
var total float32

type user struct {
	name string
	age  int
}

func (u *user) setName(name string) {
	u.name = name
}

func main() {

	myUser := user{
		name: "Luiz",
		age:  1,
	}

	var numPrt *int
	num := 10
	numPrt = &num
	*numPrt = 30
	fmt.Println(num)

	fmt.Println(myUser.name)
	myUser.setName("Luiza sN")

	var users []user
	users = append(users, myUser)
	users = append(users, user{name: "Luiza", age: 3})

	for i, user := range users {
		fmt.Println(i, user.name)
	}

	logins := make(map[string]user)
	logins["lza"] = user{name: "Luiza", age: 3}
	logins["lzz"] = user{name: "Luiz", age: 48}

	fmt.Println("aqui MAPA: ", logins["lzz"])

	filename, err := files.Upload("todo.txt", 100, "abc", "def")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(filename)

	enabled = true

	if enabled {
		fmt.Println("It's Enable")
	} else {
		fmt.Println("It's Disable")
	}
	fmt.Println("Here print!")

	if name == "Luiza" {
		fmt.Println("Hey Luiza")
	} else {
		fmt.Println("Hey")
	}

	switch name {
	case "Luiz":
		fmt.Println("Hey Luiz")
	case "Luiza":
		fmt.Println("Hey Luiza")
	case "default":
		fmt.Println("Hey")

	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
