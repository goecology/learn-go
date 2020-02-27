package main

import (
	"fmt"
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cmd"
	"github.com/i2eco/muses/pkg/database/mongo"
)

type User struct {
	Uid  int
	Name string
}

func main() {
	app := muses.Container(
		cmd.Register,
		mongo.Register,
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}
	cnt, err := mongo.CopyDb("muses").C("muses").Find(nil).Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("cnt", cnt)

}
