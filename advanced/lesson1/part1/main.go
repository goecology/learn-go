package main

import (
	"fmt"
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cmd"
	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:  "hello",
	Long: `Show hello information`,
	Run:  helloCmd,
}

func helloCmd(cmd *cobra.Command, args []string) {
	fmt.Println("hello goecology")
}

func main() {
	app := muses.Container(
		cmd.Register,
	)
	app.SetRootCommand(func(cobraCommand *cobra.Command) {
		cobraCommand.AddCommand(HelloCmd)
	})

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
