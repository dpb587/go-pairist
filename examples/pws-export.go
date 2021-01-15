package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/dpb587/go-pairist/api"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: go run ./%s USERNAME PASSWORD\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	client := api.NewClient(
		http.DefaultClient,
		"example-project-id",
		&api.Auth{
			APIKey:   "example-api-key",
			Email:    os.Args[1],
			Password: os.Args[2],
		},
	)

	historical, err := client.GetTeamHistorical(os.Args[1])
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'

	trackNames := []string{"Deploy Queue 🚀", "Support 😇", "Top of Backlog"}
	w.Write(append([]string{"Time"}, trackNames...))

	for rawTimestamp, pairing := range *historical {
		var timestampInt int64
		if timestampInt, err = strconv.ParseInt(rawTimestamp, 10, 64); err != nil {
			fmt.Fprintf(os.Stderr, "WARN: failed to parse timestamp %s; skipping\n", rawTimestamp)
			continue
		}

		formattedTimestamp := time.Unix(timestampInt/1000, 0).Format(time.RFC3339)
		record := []string{formattedTimestamp}

		for _, trackName := range trackNames {
			tracks := pairing.ByTrack(trackName)

			for _, track := range tracks {
				var people [3]string

				for peopleIdx, person := range track.People {
					if peopleIdx > 2 {
						fmt.Fprintf(os.Stderr, "WARN: skipping person #%d in role %s on %s\n", peopleIdx+1, trackName, formattedTimestamp)

						continue
					}

					people[peopleIdx] = person.DisplayName
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
