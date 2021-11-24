# Things TODO

## Functions to implement
yet unimplemented functions
```go
func getGame(w http.ResponseWriter, r *http.Request) {}

func getGameState(w http.ResponseWriter, r *http.Request) {}

func createGame(w http.ResponseWriter, r *http.Request) {}

func inviteToGame(w http.ResponseWriter, r *http.Request) {}

func joinGame(w http.ResponseWriter, r *http.Request) {}

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
