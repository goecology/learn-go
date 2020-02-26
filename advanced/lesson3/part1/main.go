package main

import (
	"fmt"
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cache/redis"
	"github.com/i2eco/muses/pkg/cmd"
)

func main() {
	app := muses.Container(
		cmd.Register,
		redis.Register,
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}

	resp1, err := redis.Caller("muses").Set("hello", "world", 0)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("resp1", resp1)

	resp2, err := redis.Caller("muses").GetString("hello")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("resp2", resp2)

}
