# Things TODO

## Refactor models
Some models are not working out very well with databases

### User
Look into merging User and UserID

### Game
This has multiple problems

#### Relational db issues
A lot of things don't work with relational databases so we need to get creative.
Investigate how quartopiece, board, player lists, quartopiece lists can be stored.
Look into [storing stuff as JSON](https://www.digitalocean.com/community/tutorials/working-with-json-in-mysql) as a possible solution

#### Active/Inactive Players
This needs to be moved to GameState which will cause breakage.
Needs to be done before DBs are added


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

### MySQL
Probably using [this driver](https://github.com/go-sql-driver/mysql)

### Postgres
Probably using [this driver](https://github.com/jackc/pgx)

## Git commit and date vars
Add git commit and date variables to `/info` endpoint
