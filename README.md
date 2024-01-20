[![CI](https://github.com/carloscasalar/idle-fantasy-story/actions/workflows/main.yml/badge.svg)](https://github.com/carloscasalar/idle-fantasy-story/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/carloscasalar/idle-fantasy-story)](https://goreportcard.com/report/github.com/carloscasalar/idle-fantasy-story)

# idle-fantasy-story
This is meant to be a service that generates a random fantasy story. It will stream the story to all clients.

## Env vars
| Name                               | Description                                    | Default                          |
|------------------------------------|------------------------------------------------|----------------------------------|
| `API_PORT`                         | Port to listen to                              | 8080                             |
| `API_LOG_FORMATTER`                | Log formatter to use. Can be `json` or `text`. | json                             |
| `API_LOG_LEVEL`                    | Log level to use                               | info                             |
| `API_MEMORYSTORAGE_WORLDSFILEPATH` | Path to the file that contains the worlds      | init/storage/inmemory/worlds.yml |

## Start the service
To start the server you can use make:
```bash
make run
```

You can also use env vars to modify the behaviour of the server like this:
```bash
API_PORT=3000 API_LOG_FORMATTER=json API_LOG_LEVEL=info make run
```

## Call the GRPC endpoints
You can use the `grpcurl` tool to call the endpoints.

This way you can discover the grpc endpoints:
```bash
grpcurl -plaintext localhost:8080 list  
```

And you can call the endpoints like this:
```bash
grpcurl -plaintext -d '{}' localhost:8080 idlefantasystory.v1.StoryService/GetWorlds
```
