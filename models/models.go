package models

// Variable of all Quarto pieces
var AllQuartoPieces = [16]*QuartoPiece{
	// All false
	&QuartoPiece{
		Id:     0,
		Dark:   false,
		Short:  false,
		Hollow: false,
		Round:  false,
	},
	// One true
	&QuartoPiece{
		Id:     1,
		Dark:   true,
		Short:  false,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Id:     2,
		Dark:   false,
		Short:  true,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Id:     3,
		Dark:   false,
		Short:  false,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Id:     4,
		Dark:   false,
		Short:  false,
		Hollow: false,
		Round:  true,
	},
	// Two true
	&QuartoPiece{
		Id:     5,
		Dark:   true,
		Short:  true,
		Hollow: false,
		Round:  false,
	},
	&QuartoPiece{
		Id:     6,
		Dark:   false,
		Short:  true,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Id:     7,
		Dark:   false,
		Short:  true,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Id:     8,
		Dark:   true,
		Short:  false,
		Hollow: true,
		Round:  false,
	},
	&QuartoPiece{
		Id:     9,
		Dark:   true,
		Short:  false,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Id:     10,
		Dark:   false,
		Short:  false,
		Hollow: true,
		Round:  true,
	},
	// Three true
	&QuartoPiece{
		Id:     11,
		Dark:   false,
		Short:  true,
		Hollow: true,
		Round:  true,
	},
	&QuartoPiece{
		Id:     12,
		Dark:   true,
		Short:  false,
		Hollow: true,
		Round:  true,
	},
	&QuartoPiece{
		Id:     13,
		Dark:   true,
		Short:  true,
		Hollow: false,
		Round:  true,
	},
	&QuartoPiece{
		Id:     14,
		Dark:   true,
		Short:  true,
		Hollow: true,
		Round:  false,
	},
	// All true
	&QuartoPiece{
		Id:     15,
		Dark:   true,
		Short:  true,
		Hollow: true,
		Round:  true,
	},
}

// Variable of empty game board
var EmptyBoard = [4][4]*QuartoPiece{
	{&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}},
	{&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}},
	{&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}},
	{&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}},
}

// Variable of empty allquartopieces
var EmptyPieces = [16]*QuartoPiece{
	&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1},
	&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1},
	&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1},
	&QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1}, &QuartoPiece{Id: -1},
}

// Variable of empty quartopiece
var EmptyQuartoPiece = &QuartoPiece{Id: -1}

// Constant for maximum amount of players per game
const MaxPlayers int = 2

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

type Game struct {
	GameId         string             `json:"game_id"`
	ActivePlayers  []*UserId          `json:"active_players"`
	InvitedPlayers []*UserId          `json:"invited_players"`
	ActivityStatus bool               `json:"activity_status"`
	NextPlayer     *UserId            `json:"next_player"`
	NextPiece      *QuartoPiece       `json:"next_piece"`
	Board          [4][4]*QuartoPiece `json:"board"`
	UnusedPieces   [16]*QuartoPiece   `json:"unused_pieces"`
	Winner         *UserId            `json:"winner"`
}

// Move in a Game
type GameMove struct {
	UserName  string       `json:"username"`
	UserId    string       `json:"user_id"`
	PositionX int32        `json:"position_x"`
	PositionY int32        `json:"position_y"`
	NextPiece *QuartoPiece `json:"next_piece"`
}

// Game Piece
type QuartoPiece struct {
	Id     int
	Dark   bool
	Short  bool
	Hollow bool
	Round  bool
}

// Database interface
type QuartoStorage interface {
	AddUser(*User) error
	AddUserId(*UserId) error
	GetUserIdFromUserId(userid string) (*UserId, error)
	GetUserIdFromUserName(userid string) (*UserId, error)
	AddGame(*Game) error
	GetGame(gameid string) (*Game, error)
	ChangeGame(*Game, *GameMove) error
	GetAllGames() ([]*Game, error)
	InviteUser(userid string, gameid string) error
	JoinUser(userid string, gameid string) error
}
