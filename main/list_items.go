package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type ListItemsCmd struct {
	*Opts `no-flag:"true"`

	Args ListItemsArgs `positional-args:"true" required:"true"`
}

type ListItemsArgs struct {
	Team  string `positional-arg-name:"TEAM" description:"Team ID"`
	Title string `positional-arg-name:"LIST-TITLE" description:"Title of list to show"`
}

func (c *ListItemsCmd) Execute(_ []string) error {
	lists, err := c.GetClient().GetTeamLists(c.Args.Team)
	if err != nil {
		return errors.Wrap(err, "fetching lists")
	}

	for _, list := range *&lists.Lists {
		if list.Title != c.Args.Title {
			continue
		}

		for _, item := range list.Items {
			fmt.Printf("%s\n", item.Text)
		}
	}

	return nil
}
