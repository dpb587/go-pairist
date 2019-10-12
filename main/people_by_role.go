package main

import (
	"fmt"

	"github.com/dpb587/go-pairist/denormalized"
	"github.com/pkg/errors"
)

type PeopleByRoleCmd struct {
	*Opts `no-flag:"true"`

	Args PeopleByRoleArgs `positional-args:"true" required:"true"`
}

type PeopleByRoleArgs struct {
	Name string `positional-arg-name:"ROLE-NAME" description:"Name of role to show"`
}

func (c *PeopleByRoleCmd) Execute(_ []string) error {
	plan, err := c.GetClient().GetTeamCurrent(c.TeamName)
	if err != nil {
		return errors.Wrap(err, "fetching team")
	}

	for _, role := range denormalized.BuildLanes(plan).ByRole(c.Args.Name) {
		for _, person := range role.People {
			fmt.Printf("%s\t%s\n", person.Name, person.Picture)
		}
	}

	return nil
}
