package main

import (
	"fmt"
	"os"

	"github.com/dpb587/go-pairist/api"
)

func main() {
	lists, err := api.DefaultClient.GetTeamLists(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, list := range *lists {
		fmt.Printf("%s\n\n", list.Title)

		for _, item := range list.Items {
			fmt.Printf("* [%s] %s\n", map[bool]string{false: " ", true: "x"}[item.Checked], item.Title)
		}

		fmt.Printf("\n")
	}
}
