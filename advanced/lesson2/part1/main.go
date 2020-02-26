package main

import (
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cmd"
	"github.com/i2eco/muses/pkg/database/mysql"
)

type User struct {
	Uid  int
	Name string
}

func main() {
	app := muses.Container(
		cmd.Register,
		mysql.Register,
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}
	u := User{}
	if mysql.Caller("muses") != nil {
		mysql.Caller("muses").Where("uid=?", 1).Find(&u)
	}
}
