# Quarto Go API
REST API written in Go to play Quarto

## Technology Used
The following is a non-exhaustive list of technology used to build this
- `Go` as the language
- `gorilla/mux` for setting up routes
- built-in Go testing harness

## Running Application
from the root of the repository run:
```
go run main.go vars.go
```

## Running Tests
from the root of the repository run:
```
go test
```
or if you want lots of debug output, run:
```
go test -v
```

# Usage
how to use with curl

## Register User
run:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "password": "verybigsecret"}'\
	localhost:8000/user
```

returns:
```json
{"username":"someuser","user_id":"G8boeMc7g"}
```

or if you have `jq` installed you can extract the `user_id` value using:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "password": "verybigsecret"}'\
	localhost:8000/user | jq '.user_id'
```

## Create Game
run:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"user_id": "G8boeMc7g"}' localhost:8000/game
```

returns:
```json
{"game_id":"NvFtm757g","players":[{"username":"","user_id":"G8boeMc7g"}],"activity_status":true,"game_state":{}}
```
