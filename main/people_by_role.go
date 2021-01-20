package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type PeopleByRoleCmd struct {
	*Opts `no-flag:"true"`

	Args PeopleByRoleArgs `positional-args:"true" required:"true"`
}

type PeopleByRoleArgs struct {
	Team     string `positional-arg-name:"TEAM" description:"Team ID"`
	RoleName string `positional-arg-name:"ROLE-NAME" description:"Name of role to show"`
}

func (c *PeopleByRoleCmd) Execute(_ []string) error {
	pairing, err := c.GetClient().GetTeamCurrent(c.Args.Team)
	if err != nil {
		return errors.Wrap(err, "fetching team")
	}

	for _, role := range pairing.ByRole(c.Args.RoleName) {
		for _, person := range role.People {
			fmt.Printf("%s\n", person.DisplayName)
		}
	}

	return nil
}
