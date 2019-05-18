package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dpb587/go-pairist/api"
	"github.com/dpb587/go-pairist/denormalized"
)

func main() {
	client := api.DefaultClient

	if os.Getenv("PAIRIST_API_KEY") != "" && len(os.Args) > 2 {
		client = api.NewClient(
			http.DefaultClient,
			api.DefaultDatabaseURL,
			&api.Auth{
				APIKey:   os.Getenv("PAIRIST_API_KEY"),
				Team:     os.Args[1],
				Password: os.Args[2],
			},
		)
	}

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
