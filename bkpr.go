package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/colvin/bkpr/internal"
	"github.com/urfave/cli"
	"encoding/json"
	"log"
	"os"
)

func main() {
        cli.VersionFlag = cli.BoolFlag{
                Name: "version, V",
                Usage: "print the version",
        }

	app := cli.NewApp()
	app.Name = "bkpr"
	app.Usage = "manage bhyve virtual machines"
	app.Version = bkpr.Version()
	app.Copyright = "(c) 2018 Colvin Wellborn"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Colvin Wellborn",
			Email: "colvinwellborn@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
                cli.BoolFlag{
                        Name: "verbose, v",
                        Usage: "increase output",
                },
                cli.BoolFlag{
                        Name: "quiet, q",
                        Usage: "decrease output",
                },
		cli.StringFlag{
			Name: "db",
			Usage: "`PATH` to the sqlite3 database",
			Value: bkpr.DEFAULT_DB_PATH,
			EnvVar: "BKPR_DB",
			Destination: &bkpr.Context.Db,
		},
        }

	app.Commands = []cli.Command{
		{
			Name:     "list",
			Usage:    "print known guests",
			Category: "Guest Management",
			Flags: []cli.Flag{
				bkpr.GuestArgs["id"],
				bkpr.GuestArgs["name"],
			},
			Action: func(c *cli.Context) error {
				db, err := bkpr.Db_connect()
				if err != nil {
					return err
				}
				defer db.Close()

				var rows = new(sql.Rows)
				if id := c.Int("id"); id != 0 {
					rows, err = db.Query("SELECT * FROM guest WHERE id = ?", id)
				} else if name := c.String("name"); name != "" {
					rows, err = db.Query("SELECT * FROM guest WHERE name = ?", name)
				} else {
					rows, err = db.Query("SELECT * FROM guest")
				}

				if err != nil {
					return err
				}
				defer rows.Close()
				guests := new([]bkpr.Guest)
				for rows.Next() {
					g := bkpr.Guest{}
					err = rows.Scan(
						&g.Id,
						&g.Name,
						&g.Os,
						&g.Loader,
					)
					if err != nil {
						return err
					}
					*guests = append(*guests, g)
				}
				j, err := json.Marshal(guests)
				fmt.Printf("%s\n", j)

				return nil
			},
		},
		{
			Name:     "create",
			Usage:    "create a new guest",
			Category: "Guest Management",
			Flags: []cli.Flag{
				bkpr.GuestArgs["name"],
				bkpr.GuestArgs["os"],
				bkpr.GuestArgs["loader"],
			},
			Action: func(c *cli.Context) error {
				g := bkpr.Guest{
					Name:   c.String("name"),
					Os:     c.String("os"),
					Loader: c.String("loader"),
				}
				fmt.Println(g)
				return nil
			},
		},
		{
			Name:     "modify",
			Usage:    "modify the configuration of a guest",
			Category: "Guest Management",
			Flags: []cli.Flag{
				bkpr.GuestArgs["id"],
				bkpr.GuestArgs["name"],
				bkpr.GuestArgs["os"],
				bkpr.GuestArgs["loader"],
			},
			Action: func(c *cli.Context) error {
				if id := c.Int("id"); id != 0 {
					fmt.Printf("would modify ID %d\n", id)
				} else if name := c.String("name"); name != "" {
					fmt.Printf("would modify NAME %s\n", name)
				} else {
					return errors.New("must provide either --id or --name")
				}
				return nil
			},
		},
		{
			Name:     "destroy",
			Usage:    "permanently remove a guest",
			Category: "Guest Management",
			Flags: []cli.Flag{
				bkpr.GuestArgs["id"],
				bkpr.GuestArgs["name"],
			},
			Action: func(c *cli.Context) error {
				if id := c.Int("id"); id != 0 {
					fmt.Printf("would destroy id %d\n", id)
				} else if name := c.String("name"); name != "" {
					fmt.Printf("would destroy name %s\n", name)
				} else {
					return errors.New("must provide either --id or --name")
				}
				return nil
			},
		},
		{
			Name:     "status",
			Usage:    "report on running guests",
			Category: "Runtime",
			Flags: []cli.Flag{
				bkpr.GuestArgs["id"],
				bkpr.GuestArgs["name"],
			},
			Action: func(c *cli.Context) error {
				if id := c.Int("id"); id != 0 {
					fmt.Printf("would report on status of id %d\n", id)
				} else if name := c.String("name"); name != "" {
					fmt.Printf("would report on status of name %s\n", name)
				} else {
					fmt.Printf("would report status of ALL\n")
				}
				return nil
			},
		},
		{
			Name:     "run",
			Usage:    "run a guest",
			Category: "Runtime",
			Flags: []cli.Flag{
				bkpr.GuestArgs["id"],
				bkpr.GuestArgs["name"],
				bkpr.GuestArgs["os"],
				bkpr.GuestArgs["loader"],
			},
			Action: func(c *cli.Context) error {
				if id := c.Int("id"); id != 0 {
					fmt.Printf("would run id %d\n", id)
				} else if name := c.String("name"); name != "" {
					fmt.Printf("would run name %s\n", name)
				} else {
					return errors.New("must provide either --id or --name")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
