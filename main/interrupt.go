package main

import (
	"fmt"
	"os"

	"github.com/dpb587/go-pairist/api"
	"github.com/dpb587/go-pairist/denormalized"
)

func main() {
	curr, err := api.DefaultClient.GetTeamCurrent(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, role := range denormalized.BuildLanes(curr).ByRole("interrupt") {
		for _, person := range role.People {
			fmt.Printf("%s\n", person.Name)
		}
	}
}
