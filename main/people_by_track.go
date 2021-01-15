package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type PeopleByTrackCmd struct {
	*Opts `no-flag:"true"`

	Args PeopleByTrackArgs `positional-args:"true" required:"true"`
}

type PeopleByTrackArgs struct {
	Team      string `positional-arg-name:"TEAM" description:"Team ID"`
	TrackName string `positional-arg-name:"TRACK-NAME" description:"Name of track to show"`
}

func (c *PeopleByTrackCmd) Execute(_ []string) error {
	pairing, err := c.GetClient().GetTeamCurrent(c.Args.Team)
	if err != nil {
		return errors.Wrap(err, "fetching team")
	}

	for _, role := range pairing.ByTrack(c.Args.TrackName) {
		for _, person := range role.People {
			fmt.Printf("%s\n", person.DisplayName)
		}
	}

	return nil
}
