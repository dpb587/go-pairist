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
	FirebaseAPIKey    string `long:"firebase-api-key" description:"Firebase API key" env:"PAIRIST_FIREBASE_API_KEY"`
	FirebaseProjectID string `long:"firebase-project-id" description:"Firebase project ID" env:"PAIRIST_FIREBASE_PROJECT_ID"`
	Email             string `long:"email" description:"Team name" env:"PAIRIST_EMAIL"`
	Password          string `long:"password" description:"Team password" env:"PAIRIST_PASSWORD"`
}

func (o Opts) GetClient() *api.Client {
	auth := &api.Auth{
		APIKey:   o.FirebaseAPIKey,
		Email:    o.Email,
		Password: o.Password,
	}

	return api.NewClient(http.DefaultClient, o.FirebaseProjectID, auth)
}
