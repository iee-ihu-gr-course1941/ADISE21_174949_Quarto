package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"os"
)

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
	testUsers = append(testUsers, u) //AddUser(u)
	uid := &UserId{
		UserName: u.UserName,
		UserId:   shortid.MustGenerate(),
	}
	testUserIds = append(testUserIds, uid) //AddUserId(uid)
	json.NewEncoder(w).Encode(uid)
	return
}

func getGame(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func getGameState(w http.ResponseWriter, r *http.Request) {
	// Empty for now
}

func createGame(w http.ResponseWriter, r *http.Request) {
	log.Println("createGame called")
	w.Header().Set("Content-Type", "application/json")
	//user that creates the game
	uid := &UserId{}
	err := json.NewDecoder(r.Body).Decode(uid)
	if err != nil || uid.UserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	//create a new game instance
	g := &Game{
		GameId:         shortid.MustGenerate(),
		ActivityStatus: true,
		State: &GameState{
			Board:        EmptyBoard,
			UnusedPieces: AllQuartoPieces,
		},
	}
	for _, u := range testUserIds {
		if u.UserId == uid.UserId {
			uid = u
			break
		}
	}
	//automatically invite the game creator to the game
	g.InvitedPlayers = append(g.InvitedPlayers, uid)
	testGames = append(testGames, g) //AddGame(g)
	json.NewEncoder(w).Encode(g)
}

func inviteToGame(w http.ResponseWriter, r *http.Request) {
	log.Println("inviteToGame called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]

	//user to be invited
	var uid *UserId = nil
	//get the name of the user to be invited from path param
	inviteeName, _ := params["username"]
	//see if user exists in the user database
	for idx, u := range testUserIds { //GetUserId(u)
		if u.UserName == inviteeName {
			uid = testUserIds[idx]
			break
		}
	}
	//return error if user with username can't be found
	if uid == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(UserNotFound))
		return
	}
	//append player to game if game exists
	for _, g := range testGames { //InviteUser(u)
		if g.GameId == gameId {
			g.InvitedPlayers = append(g.InvitedPlayers, uid)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(MsgSuccess))
			return
		}
	}
	//return error if game doesn't exist
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(GameNotFound))
	return
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
	userRouter.HandleFunc("/register", createUser) //not REST-y
	// Set up routes for game API
	gameRouter.HandleFunc("", createGame)
	gameRouter.HandleFunc("/new", createGame) //not REST-y
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
