# Things TODO

## inviteToGame()
Check if game exists before inviting user

## Refactor models
Some models are not working out very well with databases

### User
Look into merging User and UserID

### Game
This has some problems

#### Active/Inactive Players
This needs to be moved to GameState which will cause breakage.

## Functions to implement
yet unimplemented functions
```go
func playInGame(w http.ResponseWriter, r *http.Request) {}

func checkGameState(gameId string) {}
```

## User checks
Handle same user being able to be added twice to stuff

## Add Databases
Currently only a mock in-memory repo exists but in prod we need DBs

### Postgres
Depends on time. Will probably use [this driver](https://github.com/jackc/pgx)

## Git commit and date vars
Add git commit and date variables to `/info` endpoint
