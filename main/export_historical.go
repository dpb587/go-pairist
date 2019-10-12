package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dpb587/go-pairist/denormalized"
	"github.com/pkg/errors"
)

type ExportHistoricalCmd struct {
	*Opts `no-flag:"true"`
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
	rawHistorical, err := c.GetClient().GetTeamHistorical(c.TeamName)
	if err != nil {
		return errors.Wrap(err, "fetching lists")
	}

	for _, pairing := range denormalized.BuildHistory(*rawHistorical) {
		record := exportHistoricalPairing{
			Timestamp: pairing.Timestamp.Format(time.RFC3339),
		}

		for _, lane := range pairing.Lanes {
			exportLane := exportHistoricalPairingLane{}

			for _, person := range lane.People {
				exportLane.People = append(exportLane.People, person.Name)
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
