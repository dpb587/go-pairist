package api

type TeamPairingHistory map[int]TeamPairing

type TeamPairing struct {
	Entities map[string]TeamPairingEntity `json:"entities,omitempty"`
	Lanes    map[string]TeamPairingLane   `json:"lanes,omitempty"`
}

type TeamPairingEntity struct {
	Color     string   `json:"color,omitempty"`
	Icon      string   `json:"icon,omitempty"`
	Location  string   `json:"location,omitempty"`
	Name      string   `json:"name,omitempty"`
	Picture   string   `json:"picture,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Type      string   `json:"type,omitempty"`
	UpdatedAt uint     `json:"updatedAt,omitempty"`
}

type TeamPairingLane struct {
	Locked    bool `json:"locked,omitempty"`
	SortOrder uint `json:"sortOrder,omitempty"`
}

type TeamLists map[string]TeamList

type TeamList struct {
	Items TeamListItems `json:"items,omitempty"`
	Title string        `json:"title,omitempty"`
}

type TeamListItems map[string]TeamListItem

type TeamListItem struct {
	Checked bool   `json:"checked,omitempty"`
	Title   string `json:"title,omitempty"`
}
