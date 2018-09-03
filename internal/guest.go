package bkpr

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
)

type Guest struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Os string `json:"os"`
	Loader string `json:"loader"`
}

func (g Guest) String() string {
	j, _ := json.Marshal(g)
	return fmt.Sprintf("%s", j)
}

var GuestArgs = map[string]cli.Flag{
	"id":	cli.IntFlag{
		Name:	"id, i",
		Usage:	"guest `ID`",
	},
	"name": cli.StringFlag{
		Name:	"name, n",
		Usage:	"guest `NAME`",
	},
	"os":	cli.StringFlag{
		Name:	"os, o",
		Usage:	"guest `OS`",
		Value:	"FreeBSD",
	},
	"loader": cli.StringFlag{
		Name:	"loader, l",
		Usage:	"bhyve `LOADER`",
		Value:	"UEFI",
	},
}

