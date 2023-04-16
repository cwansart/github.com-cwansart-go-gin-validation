# Struct input validation with Gin Gonic

An example project with 2 REST endpoints which retrieves and updates a config file.

## Config struct

The config struct has 2 fields: FilePath and Port.

```go
type  Config  struct {
    FilePath string  `json:"file_path" binding:"required,dir"`  // binding for gin, validate for go-playground/validator
    Port uint16  `json:"port" binding:"required,max=9999,min=1000"`
}
```

FilePath represents a path on the server, whereas Port represents a numeric value of the server port.
FilePath triggers a validation error if the given path does not exist.
Port triggers a validation error, if the port is not between 1000 and 9999.

## REST endpoints

| Endpoint | Description |
|--|--|
| GET /config | retrieves the current config |
| PATCH /config | updates the current config |

## Server start

To start the server run: `go run .`

## Test curl call

### Valid call

```sh
curl http://localhost:2222/config --include --header "Content-Type: application/json" --request "PATCH" --data '{"file_path": "C:\\","port": 1000}'
```

### Invalid call, non existing file_path

```sh
curl http://localhost:2222/config --include --header "Content-Type: application/json" --request "PATCH" --data '{"file_path": "C:\\doesnotexist","port": 1000}'
```

### Invalid call, port lower than min

```sh
curl http://localhost:2222/config --include --header "Content-Type: application/json" --request "PATCH" --data '{"file_path": "C:\\","port": 200}'
```
