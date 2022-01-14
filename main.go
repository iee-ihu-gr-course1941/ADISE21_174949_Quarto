package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
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
	//check if user exists
	uid, err = gamedb.GetUserIdFromUserId(uid.UserId)
	if err != nil {
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
		NextPlayer:     uid,
	}
	//automatically invite the game creator to the game
	g.InvitedPlayers = append(g.InvitedPlayers, uid)
	err = gamedb.AddGame(g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	} else {
		g, err = gamedb.GetGame(g.GameId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(ServerError))
			return
		}
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
	//check if game exists
	_, err := gamedb.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(GameNotFound))
		return
	}
	//user to be invited
	var uid *models.UserId = nil
	//get the name of the user to be invited from path param
	inviteeName, _ := params["username"]
	//see if user exists in the user database
	uid, err = gamedb.GetUserIdFromUserName(inviteeName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(UserNotFound))
		return
	}
	//append player to game
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
	//log.Println("playInGame called")
	w.Header().Set("Content-Type", "application/json")
	//get the path parameters
	params := mux.Vars(r)
	//get game_id from path param
	gameId, _ := params["game_id"]

	//get game
	g, err := gamedb.GetGame(gameId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(NotFound))
		return
	}
	//check if two players have joined
	if len(g.ActivePlayers) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "need exactly 2 players to play this game"}`))
		return
	}
	//get game move
	gameMove := &models.GameMove{}
	err = json.NewDecoder(r.Body).Decode(gameMove)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(BadReq))
		return
	}
	//user trying to join
	uid := &models.UserId{}
	uid.UserName = gameMove.UserName
	uid.UserId = gameMove.UserId
	uid, err = gamedb.GetUserIdFromUserId(uid.UserId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(UserNotFound))
		return
	}
	//make sure the game move's next piece has been set
	if gameMove.NextPiece == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ServerError))
		return
	} else {
		g.NextPiece = gameMove.NextPiece
	}
	// if requesting player is not player playing next, error out
	if g.NextPlayer.UserId != uid.UserId {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(Unauth))
		return
	}
	//if game move seems fine and user exists, put piece there
	g.Board[gameMove.PositionX][gameMove.PositionY] = g.NextPiece
	//TODO: maybe return game state somewhere here
	done := checkGameState(g.GameId)
	if done {
		g.ActivityStatus = false
		g.Winner = uid
		err := gamedb.ChangeGame(g, gameMove)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "couldn't register move"}`))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "` + uid.UserName + ` is the winner!"}`))
		return
	} else {
		if uid.UserName == g.ActivePlayers[0].UserName {
			g.NextPlayer = g.ActivePlayers[1]
		} else if uid.UserName == g.ActivePlayers[1].UserName {
			g.NextPlayer = g.ActivePlayers[0]
		}
		err := gamedb.ChangeGame(g, gameMove)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "couldn't register move"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

func ifQuarto(qp [4]*models.QuartoPiece) bool {
	//TODO: replace with x, y, z, w vars for readability
	if qp[0].Dark == qp[1].Dark == qp[2].Dark == qp[3].Dark {
		return true
	} else if qp[0].Short == qp[1].Short == qp[2].Short == qp[3].Short {
		return true
	} else if qp[0].Hollow == qp[1].Hollow == qp[2].Hollow == qp[3].Hollow {
		return true
	} else if qp[0].Round == qp[1].Round == qp[2].Round == qp[3].Round {
		return true
	} else {
		return false
	}
}

//TODO: fix and convert to use quartostorage interface
func checkGameState(gameId string) bool {
	g, err := gamedb.GetGame(gameId)
	if err != nil {
		return false
	}
	board := g.Board
	unusedPieces := g.UnusedPieces
	log.Println("unusedPieces", unusedPieces)
	// Statically define diagonal and reverse diagonal
	diag1 := [4]*models.QuartoPiece{board[0][0], board[1][1], board[2][2], board[3][3]}
	diag2 := [4]*models.QuartoPiece{board[0][3], board[1][2], board[2][1], board[3][0]}
	// Go through the board and check if anything qualifies as quarto
	for i, row := range board {
		log.Println(i, row)
		// Don't bother if 4 pieces haven't been on the board
		if cap(unusedPieces) > 12 {
			break
		}
		// Don't bother if row isn't full
		if cap(row) < 4 {
			break
		}
		// Check if current row has quarto
		if ifQuarto(row) {
			return true
		}
		// Collect items from column
		var col [4]*models.QuartoPiece
		for j, colItem := range row {
			log.Println(j, col)
			log.Println(j, colItem)
			col[j] = colItem
		}
		// Check if there are 4 pieces in the column
		if cap(col) == 4 && ifQuarto(col) {
			return true
		}
		// Check if there are 4 pieces in the diagonal
		if cap(diag1) == 4 && ifQuarto(diag1) {
			return true
		}
		// Check if there are 4 pieces in the reverse diagonal
		if cap(diag2) == 4 && ifQuarto(diag2) {
			return true
		}
	}
	// Return false if none of the above succeeded
	return false
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
		log.Println("logging Middleware", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func setupRouter(enableLoggingMiddleware bool) http.Handler {
	// Set up router
	router := mux.NewRouter()
	// Check if the app is running on the university server
	runningOnUsersIEEIHUGR := os.Getenv("QUARTO_USERS_IEE_IHU_GR")
	var userRouter *mux.Router
	var gameRouter *mux.Router
	if runningOnUsersIEEIHUGR != "" {
		// Set up weclome message at api root
		router.HandleFunc("/~it174949", getRoot)
		router.HandleFunc("/~it174949/", getRoot)
		router.HandleFunc("/~it174949/index.php", getRoot)
		router.HandleFunc("/~it174949/index.php/", getRoot)
		// Set up subrouter for user functions
		userRouter = router.PathPrefix("/~it174949/index.php/user").Subrouter()
		// Set up subrouter for game functions
		gameRouter = router.PathPrefix("/~it174949/index.php/game").Subrouter()
	} else {
		// Set up weclome message at api root
		router.HandleFunc("/", getRoot)
		// Set up subrouter for user functions
		userRouter = router.PathPrefix("/user").Subrouter()
		// Set up subrouter for game functions
		gameRouter = router.PathPrefix("/game").Subrouter()
	}
	// Set up routes for user API
	userRouter.HandleFunc("", createUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/register", createUser).Methods(http.MethodPost) //not REST-y
	// Set up routes for game API
	gameRouter.HandleFunc("", createGame).Methods(http.MethodPost)
	gameRouter.HandleFunc("/new", createGame).Methods(http.MethodPost) //not REST-y
	//gameRouter.HandleFunc("/all", getGames).Methods(http.MethodGet) //not REST-y
	gameRouter.HandleFunc("/{game_id}", getGame).Methods(http.MethodGet)
	gameRouter.HandleFunc("/{game_id}/join", joinGame).Methods(http.MethodPost)
	gameRouter.HandleFunc("/{game_id}/play", playInGame).Methods(http.MethodPost)
	//gameRouter.HandleFunc("/{game_id}/state", getGameState).Methods(http.MethodGet)
	gameRouter.HandleFunc("/{game_id}/invite/{username}", inviteToGame).Methods(http.MethodPost)
	if enableLoggingMiddleware {
		router.Use(loggingMiddleware)
	}
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
	// Print a message so there is feedback to the app admin
	log.Println("starting server at port", httpPort)
	// Start server with or without the logging middleware
	enableLoggingMiddleware := os.Getenv("ENABLE_LOGGING_MIDDLEWARE")
	if enableLoggingMiddleware != "" {
		router := setupRouter(true)
		log.Fatal(http.ListenAndServe(":"+httpPort, handlers.LoggingHandler(os.Stdout, router)))
	} else {
		router := setupRouter(false)
		log.Fatal(http.ListenAndServe(":"+httpPort, router))
	}
}
