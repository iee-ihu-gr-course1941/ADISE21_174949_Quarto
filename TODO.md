# Things TODO

## Store every game move
Add function to QuartoStorage interface for storing GameMove

## Refactor models
Some models are not working out very well with databases

### User
Look into merging User and UserID

### Game
This has some problems

## Functions to implement
yet unimplemented functions
```go
func playInGame(w http.ResponseWriter, r *http.Request) {}

func checkGameState(gameId string) {}
```

## User checks
Handle same user being able to be added twice to stuff

## Add Databases
Currently only mock and mysql exist but I like postgres

### Postgres
Depends on time. Will probably use [this driver](https://github.com/jackc/pgx)

## Git commit and date vars
Add git commit and date variables to `/info` endpoint
