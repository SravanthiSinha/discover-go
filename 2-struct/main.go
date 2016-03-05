package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (u *user) printHello() {
	fmt.Println("Hello", u.Name)
}

func (u *user) printdetails() {

	dob := strings.Split(u.DOB, ", ")
	k, _ := strconv.Atoi(dob[1])
	fmt.Println(u.Name, "who was born in", u.City, "would be", time.Now().Year()-k, "years old today")

}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	u.printHello()
	u.printdetails()
}
