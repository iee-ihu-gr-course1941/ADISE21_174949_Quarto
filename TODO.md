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

## Data storage modularity
incomplete interface
```go
type QuartoStorage interface {
	AddUser(*User) error
	AddUserId(*UserId) error
	GetUserId(userid string) (*UserId, error)
	AddGame(*Game) error
	GetGame(gameid string) (*Game, error)
	GetAllGames() ([]*Game, error)
	InviteUser(userid string, gameid string) error
	JoinUser(userid string, gameid string) error
}
```

## Move structs, consts and vars out of main.go
too much clutter in main.go

### Structs
move models to dedicated package
```go
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserId struct {
	UserName string `json:"username"`
	UserId   string `json:"user_id"`
}

type Game struct {
	GameId         string     `json:"game_id"`
	ActivePlayers  []*UserId  `json:"active_players"`
	InvitedPlayers []*UserId  `json:"invited_players"`
	ActivityStatus bool       `json:"activity_status"`
	State          *GameState `json:"game_state"`
	Winner         *UserId    `json:"winner"`
}

type GameState struct {
	NextPlayer   *UserId            `json:"next_player"`
	NextPiece    *QuartoPiece       `json:"next_piece"`
	Board        [4][4]*QuartoPiece `json:"board"`
	UnusedPieces [16]*QuartoPiece   `json:"unused_pieces"`
}

type GameMove struct {
	PositionX int32        `json:"position_x"`
	PositionY int32        `json:"position_y"`
	NextPiece *QuartoPiece `json:"next_piece"`
}

type QuartoPiece struct {
	Dark   bool
	Short  bool
	Hollow bool
	Round  bool
}
```

### Consts
most globally-available constants can be moved to an `api` package since they're http-specific
```go
const MaxPlayers int = 2

const BadReq string = `{"error": "bad request"}`

const NotFound string = `{"error": "not found"}`

const Unauth string = `{"error": "unauthorized"}`

const ServerError string = `{"error": "internal server error"}`

const MsgSuccess string = `{"message": "success"}`

const UserNotFound string = `{"error": "user not found"}`

const UserUnauth string = `{"error": "user unauthorized"}`

const GameNotFound string = `{"error": "game not found"}`
```

### Vars

#### Mock db
refactor into mock in its own package that implements `QuartoStorage`
```go
var testUsers []*User
var testUserIds []*UserId
var testGames []*Game
```
example
```go
type MockDB struct {
	Users []*User
	UserIds []*UserId
	Players []*UserId
	Games []*Game
}
```

## Refactor tests actions to be reusable
example of snippet that could be made reusable
```go
testServer := httptest.NewServer(setupRouter())
path := "/user"
testURL := testServer.URL + path
jsonData := []byte(`{"username": "myself", "password": "mypasswd"}`)
res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
if err != nil {
	t.Error("POST error:", err)
}
defer res.Body.Close()
body, err := io.ReadAll(res.Body)
if err != nil {
	t.Error("resp.Body error:", err)
}
u := &UserId{}
err = json.Unmarshal(body, u)
if err != nil {
	t.Error("unmarshal error:", err)
}
testServer.Close()
```

## Split in packages
already mentioned above in some places but the service should be split into packages

## User checks
Handle same user being able to be added twice to stuff
