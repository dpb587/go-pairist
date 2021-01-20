package api

type TeamPairingHistory map[string]TeamPairing

type TeamPairing struct {
	Pairs []TeamPair `json:"pairs,omitempty"`
}

type TeamPair struct {
	People []TeamPerson `json:"people,omitempty"`
	Roles  []TeamRole   `json:"roles,omitempty"`
	Tracks []TeamTrack  `json:"tracks,omitempty"`
}

type TeamPerson struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type TeamRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TeamTrack struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TeamLists struct {
	Lists []TeamList `json:"lists,omitempty"`
}

type TeamList struct {
	Title string         `json:"title,omitempty"`
	Order uint           `json:"order,omitempty"`
	Items []TeamListItem `json:"items,omitempty"`
}

type TeamListItem struct {
	Text      string                          `json:"text,omitempty"`
	Order     uint                            `json:"order,omitempty"`
	Reactions map[string]TeamListItemReaction `json:"reactions,omitempty"`
}

type TeamListItemReaction struct {
	Count     uint `json:"count,omitempty"`
	Timestamp uint `json:"timestamp,omitempty"`
}

func (pairing TeamPairing) ByRole(roleName string) []TeamPair {
	var res []TeamPair

	for _, pair := range pairing.Pairs {
		for _, role := range pair.Roles {
			if role.Name == roleName {
				res = append(res, pair)

				break
			}
		}
	}

	return res
}

func (pairing TeamPairing) ByTrack(trackName string) []TeamPair {
	var res []TeamPair

	for _, pair := range pairing.Pairs {
		for _, track := range pair.Tracks {
			if track.Name == trackName {
				res = append(res, pair)

				break
			}
		}
	}

	return res
}
