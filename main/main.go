package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	o := &Opts{}
	c := cli{
		Opts:             o,
		ExportHistorical: ExportHistoricalCmd{Opts: o},
		ListItems:        ListItemsCmd{Opts: o},
		PeopleByRole:     PeopleByRoleCmd{Opts: o},
		PeopleByTrack:    PeopleByTrackCmd{Opts: o},
	}

	var parser = flags.NewParser(&c, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
