package main

import (
	"fmt"
	"os"

	"github.com/dpb587/go-pairist/api/authenticated"
	"github.com/dpb587/go-pairist/denormalized"
)

func main() {
	client, err := authenticated.CreateClient(os.Getenv("PAIRIST_API_KEY"), os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	// client := anonymous.DefaultClient

	curr, err := client.GetTeamCurrent(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, role := range denormalized.BuildLanes(curr).ByRole("Interrupt") {
		for _, person := range role.People {
			fmt.Printf("%s\n", person.Name)
		}
	}
}
