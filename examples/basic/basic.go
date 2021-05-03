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

func test() {
	dbb := ns.DBB{}
	u := UserDB{}
	users := make([]*User, 0)

	dbb.
		Where(u.Age.Eq(u.Age)).
		Find(&users)
}
