# idle-fantasy-story
This is meant to be a service that generates a random fantasy story. It will stream the story to all clients.

## Env vars
| Name                  | Description                                    | Default |
|-----------------------|------------------------------------------------|---------|
| `API_PORT`            | Port to listen to                              | 8080    |
| `API_LOG_FORMATTER`   | Log formatter to use. Can be `json` or `text`. | json    |
| `API_LOG_LEVEL`       | Log level to use                               | info    |

## Start the service
To start the server you can use make:
```bash
make run
```

You can also use env vars to modify the behaviour of the server like this:
```bash
API_PORT=3000 API_LOG_FORMATTER=json API_LOG_LEVEL=info make run
```
