package main

import (
	"net/http"

	"github.com/dpb587/go-pairist/api"
)

type cli struct {
	*Opts

	PeopleByRole     PeopleByRoleCmd     `command:"people-by-role" description:"Show people having a specific role"`
	PeopleByTrack    PeopleByTrackCmd    `command:"people-by-track" description:"Show people having a specific track"`
	ListItems        ListItemsCmd        `command:"list-items" description:"Show items of a list"`
	ExportHistorical ExportHistoricalCmd `command:"export-historical" description:"Export historical pairing data"`
}

type Opts struct {
	FirebaseAPIKey string `long:"firebase-api-key" description:"Firebase API key" hidden:"true"`
	FirebaseURL    string `long:"firebase-url" description:"Firebase API key" hidden:"true"`

	TeamName     string `long:"team-name" description:"Team name" env:"PAIRIST_TEAM_NAME"`
	TeamPassword string `long:"team-password" description:"Team password (required if team is private)" env:"PAIRIST_TEAM_PASSWORD"`
}

func (o Opts) GetClient() *api.Client {
	var auth *api.Auth

	if o.TeamPassword != "" {
		var firebaseAPIKey = api.DefaultFirebaseAPIKey

		if v := o.FirebaseAPIKey; v != "" {
			firebaseAPIKey = v
		}

		auth = &api.Auth{
			APIKey:   firebaseAPIKey,
			Team:     o.TeamName,
			Password: o.TeamPassword,
		}
	}

	var firebaseURL = api.DefaultFirebaseURL

	if v := o.FirebaseURL; v != "" {
		firebaseURL = v
	}

	return api.NewClient(http.DefaultClient, firebaseURL, auth)
}
