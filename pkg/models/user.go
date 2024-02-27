package main

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	SecondName string `json:"secondName"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
