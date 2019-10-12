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
	Title string `positional-arg-name:"LIST-TITLE" description:"Title of list to show"`
}

func (c *ListItemsCmd) Execute(_ []string) error {
	lists, err := c.GetClient().GetTeamLists(c.TeamName)
	if err != nil {
		return errors.Wrap(err, "fetching lists")
	}

	for _, list := range *lists {
		if list.Title != c.Args.Title {
			continue
		}

		for _, item := range list.Items {
			fmt.Printf("%s\t%s\n", map[bool]string{false: "UNCHECKED", true: "CHECKED"}[item.Checked], item.Title)
		}
	}

	return nil
}
