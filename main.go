package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"os"
)

// Variable of all Quarto pieces
var AllQuartoPieces = [16]*QuartoPiece{
	// All false
	&QuartoPiece{
		Dark:   false,
		Short:  false,
		Hollow: false,
		Round:  false,
	},
	// One true
	&QuartoPiece{
		Dark:   true,
		Short:  false,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  true,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  false,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  false,
		Hollow: false,
		Round:  true,
	},
	// Two true
	&QuartoPiece{
		Dark:   true,
		Short:  true,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  true,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  true,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Dark:   true,
		Short:  false,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Dark:   true,
		Short:  false,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Dark:   false,
		Short:  false,
		Hollow: true,
		Round:  true,
	},
	// Three true
	&QuartoPiece{
		Dark:   false,
		Short:  true,
		Hollow: true,
		Round:  true,
	},
	&QuartoPiece{
		Dark:   true,
		Short:  false,
		Hollow: true,
		Round:  true,
	},
	&QuartoPiece{
		Dark:   true,
		Short:  true,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Dark:   true,
		Short:  true,
		Hollow: true,
		Round:  false,
	},
	// All true
	&QuartoPiece{
		Dark:   true,
		Short:  true,
		Hollow: true,
		Round:  true,
	},
}

// Variable of empty game board
var EmptyBoard = [4][4]*QuartoPiece{
	{&QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}},
	{&QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}},
	{&QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}},
	{&QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}, &QuartoPiece{}},
}

// Constant for maximum amount of players per game
const MaxPlayers int = 2

// Constant for Bad Request
const BadReq string = `{"error": "bad request"}`

// Constant for Not Found
const NotFound string = `{"error": "not found"}`

// Constant for Unauthorized
const Unauth string = `{"error": "unauthorized"}`

// Constant for Unauthorized
const ServerError string = `{"error": "internal server error"}`

// Constant for success message
const MsgSuccess string = `{"message": "success"}`

// Constant for User Not Found
const UserNotFound string = `{"error": "user not found"}`

// Constant for Unauthorized
const UserUnauth string = `{"error": "user unauthorized"}`

// Constant for Game Not Found
const GameNotFound string = `{"error": "game not found"}`

// User struct with selected password
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// User struct with generated secret
type UserId struct {
	UserName string `json:"username"`
	UserId   string `json:"user_id"`
}

//TODO: rethink active/inactive players thing
type Game struct {
	GameId         string     `json:"game_id"`
	ActivePlayers  []*UserId  `json:"active_players"`
	InvitedPlayers []*UserId  `json:"invited_players"`
	ActivityStatus bool       `json:"activity_status"`
	State          *GameState `json:"game_state"`
	Winner         *UserId    `json:"winner"`
}

//TODO: fill in with fields
type GameState struct {
	NextPlayer   *UserId            `json:"next_player"`
	NextPiece    *QuartoPiece       `json:"next_piece"`
	Board        [4][4]*QuartoPiece `json:"board"`
	UnusedPieces [16]*QuartoPiece   `json:"unused_pieces"`
}

// Move in a Game
type GameMove struct {
	PositionX int32        `json:"position_x"`
	PositionY int32        `json:"position_y"`
	NextPiece *QuartoPiece `json:"next_piece"`
}

// Game Piece
type QuartoPiece struct {
	Dark   bool
	Short  bool
	Hollow bool
	Round  bool
}

//TODO: replace with database(s)
var testUsers []*User
var testUserIds []*UserId
var testPlayers []*UserId
var testGames []*Game

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("createUser called")
	w.Header().Set("Content-Type", "application/json")
	u := &User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	testUsers = append(testUsers, u)
	uid := &UserId{
		UserName: u.UserName,
		UserId:   shortid.MustGenerate(),
	}
	testUserIds = append(testUserIds, uid)
	json.NewEncoder(w).Encode(uid)
	return
}

func getGameState(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func createGame(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func inviteToGame(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func joinGame(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func playInGame(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func checkGameState(gameId string) {
	// Empty for now
}

// Function to set server HTTP port
func setupHTTPPort() string {
	httpPort := os.Getenv("QUARTO_HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}
	return httpPort
}

func setupRouter() http.Handler {
	// Set up router
	router := mux.NewRouter()
	// Set up subrouter for user functions
	userRouter := router.PathPrefix("/user").Subrouter()
	// Set up subrouter for game functions
	gameRouter := router.PathPrefix("/game").Subrouter()
	// Set up routes for user API
	userRouter.HandleFunc("", createUser)
	//userRouter.HandleFunc("/register", createUser) //not REST-y
	// Set up routes for game API
	gameRouter.HandleFunc("", createGame)
	//gameRouter.HandleFunc("/new", createGame) //not REST-y
	gameRouter.HandleFunc("/{game_id}", getGame)
	gameRouter.HandleFunc("/{game_id}/join", joinGame)
	gameRouter.HandleFunc("/{game_id}/play", playInGame)
	gameRouter.HandleFunc("/{game_id}/state", getGameState)
	gameRouter.HandleFunc("/{game_id}/invite/{username}", inviteToGame)
	return router
}

func main() {
	// Determine port to run at
	httpPort := setupHTTPPort()
	// Set up the router for the API
	router := setupRouter()
	// Print a message so there is feedback to the app admin
	log.Println("starting server at port", httpPort)
	// One-liner to start the server or print error
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
