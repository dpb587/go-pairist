# go-pairist

A simple, unofficial client for reading [Pairist](https://pair.ist/) data.


## API

The [`api`](api/) package exposes the basic endpoints for reading current and historical pairings, as well as team lists. The package exposes the data in the raw form used by Pairist. Use the [`denormalized`](denormalized/) package to transform the raw data into structures which are slightly more usable.

For anonymous access to the default Pairist server, use `api.DefaultClient`. To authenticate with team credentials, explicitly create and configure a client:

    client := api.NewClient(
      http.DefaultClient,
      api.DefaultFirebaseURL,
      &api.Auth{
        APIKey:   api.DefaultFirebaseAPIKey,
        Team:     os.Getenv("PAIRIST_TEAM_NAME"),
        Password: os.Getenv("PAIRIST_TEAM_PASSWORD"),
      },
    )


## CLI

The [`main`](main/) package exposes some utility commands for showing people in a role or track, showing lists, or exporting historical data as JSON.

    $ go run ./main -h
    Usage:
      main [OPTIONS] <command>

    Application Options:
          --team-name=     Team name [$PAIRIST_TEAM_NAME]
          --team-password= Team password (required if team is private) [$PAIRIST_TEAM_PASSWORD]

    Help Options:
      -h, --help           Show this help message

    Available commands:
      export-historical  Export historical pairing data
      list-items         Show items of a list
      people-by-role     Show people having a specific role
      people-by-track    Show people having a specific track


## License

[MIT License](LICENSE)
