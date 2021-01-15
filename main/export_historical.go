package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type ExportHistoricalCmd struct {
	*Opts `no-flag:"true"`

	Args ExportHistoricalArgs `positional-args:"true" required:"true"`
}

type ExportHistoricalArgs struct {
	Team string `positional-arg-name:"TEAM" description:"Team ID"`
}

type exportHistoricalPairing struct {
	Timestamp string                        `json:"timestamp"`
	Lanes     []exportHistoricalPairingLane `json:"lanes"`
}

type exportHistoricalPairingLane struct {
	People []string `json:"people,omitempty"`
	Roles  []string `json:"roles,omitempty"`
	Tracks []string `json:"tracks,omitempty"`
}

func (c *ExportHistoricalCmd) Execute(_ []string) error {
	historical, err := c.GetClient().GetTeamHistorical(c.Args.Team)
	if err != nil {
		return errors.Wrap(err, "fetching history")
	}

	for timestamp, pairing := range *historical {
		var timestampInt int64
		if timestampInt, err = strconv.ParseInt(timestamp, 10, 64); err != nil {
			return errors.Wrap(err, "parsing timestamp")
		}

		record := exportHistoricalPairing{
			Timestamp: time.Unix(timestampInt/1000, 0).Format(time.RFC3339),
		}

		for _, lane := range pairing.Pairs {
			exportLane := exportHistoricalPairingLane{}

			for _, person := range lane.People {
				exportLane.People = append(exportLane.People, person.DisplayName)
			}

			for _, role := range lane.Roles {
				exportLane.Roles = append(exportLane.Roles, role.Name)
			}

			for _, track := range lane.Tracks {
				exportLane.Tracks = append(exportLane.Tracks, track.Name)
			}

			record.Lanes = append(record.Lanes, exportLane)
		}

		b, err := json.Marshal(record)
		if err != nil {
			return errors.Wrap(err, "marshaling record")
		}

		fmt.Printf("%s\n", b)
	}

	return nil
}
