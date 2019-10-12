package denormalized

import (
	"sort"
	"time"

	"github.com/dpb587/go-pairist/api"
)

func BuildLanes(historical *api.TeamPairing) Lanes {
	var lanes Lanes

	for laneID := range historical.Lanes {
		lane := Lane{}

		for _, entity := range historical.Entities {
			if entity.Location != laneID {
				continue
			}

			denormalizedEntity := Entity{
				Color:     entity.Color,
				Icon:      entity.Icon,
				Name:      entity.Name,
				Picture:   entity.Picture,
				UpdatedAt: entity.UpdatedAt,
			}

			switch entity.Type {
			case "person":
				lane.People = append(lane.People, denormalizedEntity)
			case "role":
				lane.Roles = append(lane.Roles, denormalizedEntity)
			case "track":
				lane.Tracks = append(lane.Tracks, denormalizedEntity)
			}
		}

		lanes = append(lanes, lane)
	}

	return lanes
}

func BuildHistory(historical api.TeamPairingHistory) []TeamPairing {
	var result []TeamPairing

	for dayIdx, day := range historical {
		result = append(result, TeamPairing{Timestamp: time.Unix(int64(dayIdx*3600), 0), Lanes: BuildLanes(&day)})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Timestamp.Before(result[j].Timestamp) })

	return result
}
