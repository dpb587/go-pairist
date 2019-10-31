# github.com/dpb587/go-pairist

A simple, unofficial Go module for reading [Pairist](https://pair.ist/) data.

## API

The [`api`](api/) package exposes the basic endpoints for reading current and historical pairings, as well as team lists. For anonymous access to the default Pairist server, use `api.DefaultClient`.

```go
import "github.com/dpb587/go-pairist/api"

client := api.DefaultClient

pairing, err := client.GetTeamCurrent("my-team-name")
```

To authenticate with team credentials, against a different environment, or API key, explicitly create and configure a client:

```go
client := api.NewClient(
  http.DefaultClient,
  api.DefaultFirebaseURL,
  &api.Auth{
    APIKey:   api.DefaultFirebaseAPIKey,
    Team:     os.Getenv("PAIRIST_TEAM_NAME"),
    Password: os.Getenv("PAIRIST_TEAM_PASSWORD"),
  },
)
```

The `api` package exposes the data in the raw form used by Pairist which is not very easy to interact with. For more useful views and methods, use the [`denormalized`](denormalized/) package.

```go
for _, role := range denormalized.BuildLanes(pairing).ByRole("interrupt") {
  for _, person := range role.People {
    fmt.Printf("%s\t%s\n", person.Name, person.Picture)
  }
}
```

## Development

For local use, ensure you have a recent version of Go installed before cloning and using.

```console
$ git clone git@github.com:dpb587/go-pairist.git
$ cd go-pairist
$ go test ./...
```

### CLI

The [`main`](main/) package provides a limited CLI for showing people in a role or track, showing lists, or exporting historical data as JSON.

```console
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
```

### Examples

The [`examples`](examples/) directory provides some other, customized examples of using this client.

```console
$ go run ./examples/pws-export.go "$PAIRIST_TEAM_NAME" "$PAIRIST_TEAM_PASSWORD" \
  > results.csv
```

## License

[MIT License](LICENSE)
