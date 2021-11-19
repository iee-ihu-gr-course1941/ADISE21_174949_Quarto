package main

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

//TODO: refactor into mock database
var testUsers []*User
var testUserIds []*UserId
var testPlayers []*UserId
var testGames []*Game
