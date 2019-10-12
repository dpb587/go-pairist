package main

import (
	"fmt"

	"github.com/dpb587/go-pairist/denormalized"
	"github.com/pkg/errors"
)

type PeopleByTrackCmd struct {
	*Opts `no-flag:"true"`

	Args PeopleByTrackArgs `positional-args:"true" required:"true"`
}

type PeopleByTrackArgs struct {
	Name string `positional-arg-name:"TRACK-NAME" description:"Name of track to show"`
}

func (c *PeopleByTrackCmd) Execute(_ []string) error {
	plan, err := c.GetClient().GetTeamCurrent(c.TeamName)
	if err != nil {
		return errors.Wrap(err, "fetching team")
	}

	for _, role := range denormalized.BuildLanes(plan).ByTrack(c.Args.Name) {
		for _, person := range role.People {
			fmt.Printf("%s\t%s\n", person.Name, person.Picture)
		}
	}

	return nil
}
