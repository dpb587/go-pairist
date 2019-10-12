package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/dpb587/go-pairist/api"
	"github.com/dpb587/go-pairist/denormalized"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: go run ./%s USERNAME PASSWORD\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	client := api.DefaultClient

	if os.Getenv("PAIRIST_API_KEY") != "" && len(os.Args) > 2 {
		client = api.NewClient(
			http.DefaultClient,
			api.DefaultFirebaseURL,
			&api.Auth{
				APIKey:   api.DefaultFirebaseAPIKey,
				Team:     os.Args[1],
				Password: os.Args[2],
			},
		)
	}

	historical, err := client.GetTeamHistorical(os.Args[1])
	if err != nil {
		panic(err)
	}

	pairPlans := denormalized.BuildHistory(*historical)

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'

	trackNames := []string{"Deploy Queue 🚀", "Support 😇", "Top of Backlog"}
	w.Write(append([]string{"Time"}, trackNames...))

	for _, plan := range pairPlans {
		record := []string{plan.Timestamp.Format(time.RFC3339)}

		for _, trackName := range trackNames {
			tracks := plan.Lanes.ByTrack(trackName)

			for _, track := range tracks {
				var people [3]string

				for peopleIdx, person := range track.People {
					if peopleIdx > 2 {
						fmt.Fprintf(os.Stderr, "WARN: skipping person #%d in role %s on %s\n", peopleIdx+1, trackName, plan.Timestamp.Format(time.RFC3339))

						continue
					}

					people[peopleIdx] = person.Name
				}

				record = append(record, people[0], people[1], people[2])
			}
		}

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
