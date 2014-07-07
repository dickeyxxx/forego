package main

import (
	"fmt"
	"github.com/ddollar/dist"
)

var cmdUpdate = &Command{
	Run:   runUpdate,
	Usage: "update",
	Short: "Update forego",
	Long: `
Update forego

Examples:

	forego update
`,
}

func init() {
}

func runUpdate(cmd *Command, args []string) {
	if version == "dev" {
		fmt.Println("ERROR: can't update dev version")
		return
	}
	d := dist.NewDist("ddollar/forego", version)
	to, err := d.Update()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Printf("updated to %s\n", to)
	}
}
