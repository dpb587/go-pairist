# github.com/dpb587/go-pairist

A simple, unofficial Go module for reading [Pairist](https://github.com/pivotal-cf/pairist) data.

## API

The [`api`](api/) package exposes the basic endpoints for reading current and historical pairings, as well as team lists.

```go
import "github.com/dpb587/go-pairist/v2/api"

client := api.NewClient(
  http.DefaultClient,
  os.Getenv("PAIRIST_FIREBASE_PROJECT_ID"),
  &api.Auth{
    APIKey:   os.Getenv("PAIRIST_FIREBASE_API_KEY"),
    Team:     os.Getenv("PAIRIST_EMAIL"),
    Password: os.Getenv("PAIRIST_PASSWORD"),
  },
)

pairing, err := client.GetTeamCurrent("my-team-id")

for _, pair := range pairing.ByRole("interrupt") {
  for _, person := range pair.People {
    fmt.Printf("%s\n", person.DisplayName)
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
      --firebase-api-key=    Firebase API key [$PAIRIST_FIREBASE_API_KEY]
      --firebase-project-id= Firebase project ID [$PAIRIST_FIREBASE_PROJECT_ID]
      --email=               Team name [$PAIRIST_EMAIL]
      --password=            Team password [$PAIRIST_PASSWORD]

Help Options:
  -h, --help                 Show this help message

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
