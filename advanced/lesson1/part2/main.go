package main

import (
	"fmt"
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cmd"
	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:  "hello info",
	Long: `Show hello information`,
	Run:  helloCmd,
}

var content string

func init() {
	HelloCmd.PersistentFlags().StringVar(&content, "content", "", "input content value")
}

func helloCmd(cmd *cobra.Command, args []string) {
	fmt.Println("hello " + content)
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
