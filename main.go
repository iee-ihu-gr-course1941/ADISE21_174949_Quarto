package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/repo/mock"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/repo/mysql"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"os"
)

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

// Constant for welcome message
const MsgWelcome string = `Welcome to my Quarto API written in Go`

var gamedb models.QuartoStorage

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(MsgWelcome + "\n"))
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	//log.Println("createUser called")
	w.Header().Set("Content-Type", "application/json")
	u := &models.User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	err = gamedb.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	uid := &models.UserId{
		UserName: u.UserName,
		UserId:   shortid.MustGenerate(),
	}
	err = gamedb.AddUserId(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	json.NewEncoder(w).Encode(uid)
	return
}

//TODO: if user authorized
func getGame(w http.ResponseWriter, r *http.Request) {
	//log.Println("getGame called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]
	g, err := gamedb.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(NotFound))
		return
	} else {
		json.NewEncoder(w).Encode(g)
	}
	return
}

/*
func getGameState(w http.ResponseWriter, r *http.Request) {
	//log.Println("getGameState called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]
	g, err := gamedb.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(NotFound))
		return
	} else {
		json.NewEncoder(w).Encode(g.State)
	}
	return
}
*/

func createGame(w http.ResponseWriter, r *http.Request) {
	//log.Println("createGame called")
	w.Header().Set("Content-Type", "application/json")
	//user that creates the game
	uid := &models.UserId{}
	err := json.NewDecoder(r.Body).Decode(uid)
	if err != nil || uid.UserId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	//create a new game instance
	g := &models.Game{
		GameId:         shortid.MustGenerate(),
		ActivityStatus: true,
		Board:          models.EmptyBoard,
		UnusedPieces:   models.AllQuartoPieces,
	}
	uid, err = gamedb.GetUserIdFromUserId(uid.UserId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	//automatically invite the game creator to the game
	g.InvitedPlayers = append(g.InvitedPlayers, uid)
	err = gamedb.AddGame(g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	} else {
		json.NewEncoder(w).Encode(g)
		return
	}
}

func inviteToGame(w http.ResponseWriter, r *http.Request) {
	//log.Println("inviteToGame called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]

	//user to be invited
	var uid *models.UserId = nil
	//get the name of the user to be invited from path param
	inviteeName, _ := params["username"]
	//see if user exists in the user database

	uid, err := gamedb.GetUserIdFromUserName(inviteeName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(UserNotFound))
		return
	}

	//append player to game if game exists
	err = gamedb.InviteUser(uid.UserId, gameId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(MsgSuccess))
		return
	}
}

func joinGame(w http.ResponseWriter, r *http.Request) {
	//log.Println("joinGame called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]

	//user trying to join
	uid := &models.UserId{}
	err := json.NewDecoder(r.Body).Decode(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	uid, err = gamedb.GetUserIdFromUserId(uid.UserId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(UserNotFound))
		return
	}
	g, err := gamedb.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(GameNotFound))
		return
	}

	err = gamedb.JoinUser(uid.UserId, g.GameId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(MsgSuccess))
	return
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func setupRouter() http.Handler {
	// Set up router
	router := mux.NewRouter()
	// Set up weclome message at api root
	//router.HandleFunc("/", getRoot)
	router.HandleFunc("/~it174949", getRoot)
	router.HandleFunc("/~it174949/", getRoot)
	router.HandleFunc("/~it174949/index.php", getRoot)
	router.HandleFunc("/~it174949/index.php/", getRoot)
	// Set up subrouter for user functions
	userRouter := router.PathPrefix("/~it174949/index.php/user").Subrouter()
	// Set up subrouter for game functions
	gameRouter := router.PathPrefix("/~it174949/index.php/game").Subrouter()
	// Set up routes for user API
	userRouter.HandleFunc("", createUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/register", createUser).Methods(http.MethodPost) //not REST-y
	// Set up routes for game API
	gameRouter.HandleFunc("", createGame).Methods(http.MethodPost)
	gameRouter.HandleFunc("/new", createGame).Methods(http.MethodPost) //not REST-y
	gameRouter.HandleFunc("/{game_id}", getGame).Methods(http.MethodGet)
	gameRouter.HandleFunc("/{game_id}/join", joinGame).Methods(http.MethodPost)
	gameRouter.HandleFunc("/{game_id}/play", playInGame).Methods(http.MethodPost)
	//gameRouter.HandleFunc("/{game_id}/state", getGameState).Methods(http.MethodGet)
	gameRouter.HandleFunc("/{game_id}/invite/{username}", inviteToGame).Methods(http.MethodPost)
	router.Use(loggingMiddleware)
	return router
}

func init() {
	// Set up storage
	mysqlURL := os.Getenv("MYSQL_URL")
	if mysqlURL != "" {
		if mysqlURL == "test" {
			db, err := mysql.NewMysqlRepo("tester:Apasswd@tcp(localhost:3306)/tester")
			if err != nil {
				log.Fatalf("error %v", err)
			}
			gamedb = db
		} else {
			db, err := mysql.NewMysqlRepo(mysqlURL)
			if err != nil {
				log.Fatalf("error %v", err)
			}
			gamedb = db
		}
		return
	}
	gamedb, _ = mock.NewMockDB() //Error ignored because it's always nil
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
