package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("user name =", u.name, "user email =", u.email)
}

var (
	catOut  = make(chan string)
	dogOut  = make(chan string)
	fishOut = make(chan string)
)

func main() {
	// u := &user{
	// 	name:  "hts",
	// 	email: "hts_0000@sina.com",
	// }

	// sendNotify(u)
	// (*u).notify()
	// u.notify()

	// runtime.GOMAXPROCS(1)

	go catIn()
	go dogIn()
	go fishIn()

	for i := 0; i < 100; i++ {
		fmt.Println(<-catOut)
		fmt.Println(<-dogOut)
		fmt.Println(<-fishOut)
	}

}

func catIn() {
	defer close(catOut)
	for i := 0; i < 100; i++ {
		catOut <- "cat"
	}
}

func dogIn() {
	defer close(dogOut)
	for i := 0; i < 100; i++ {
		dogOut <- "dog"
	}
}

func fishIn() {
	defer close(fishOut)
	for i := 0; i < 100; i++ {
		fishOut <- "fish"
	}
}

func sendNotify(n notifier) {
	n.notify()
}
