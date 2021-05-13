package main

import (
	ns "github.com/HerringtonDarkholme/go-newspeak/pkg"
	"github.com/HerringtonDarkholme/go-newspeak/pkg/col"
)

type User struct {
	Name   string
	Age    int
	Salary int
}

type UserDB struct {
	Name, Age, Salary col.GenericColumn
}

func (*UserDB) ToTableText() string {
	return "user"
}

func test() {
	dbb := new(ns.DBB)
	u := new(UserDB)
	users := make([]*User, 0)

	dbb.
		Table(u).
		Where(u.Age.Eq(u.Age)).
		Group(u.Name).
		Find(&users)

    // works!
    dbb.Query("select * from absolutely_complicate_query")
    // does not work
    // s := "selct * from" + "bad_boy"
    // dbb.Query(s) // error
}
