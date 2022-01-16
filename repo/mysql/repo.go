package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
	"math/rand"
	"strconv"
	"time"
)

type mysqlRepo struct {
	client   *sql.DB
	mysqlURL string
}

func newMysqlClient(url string) (*sql.DB, error) {
	client, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createUserTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createUserIdTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createGameTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createBoardTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createUnusedPiecesTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createInvitedPlayerTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createActivePlayerTableQuery)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMysqlRepo(url string) (*mysqlRepo, error) {
	mysqlclient, err := newMysqlClient(url)
	if err != nil {
		return nil, err
	}
	repo := &mysqlRepo{
		mysqlURL: url,
		client:   mysqlclient,
	}
	return repo, nil
}

func (r *mysqlRepo) AddUser(u *models.User) error {
	err := r.client.QueryRow(userInsertQuery,
		u.UserName,
		u.Password,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) AddUserId(uid *models.UserId) error {
	err := r.client.QueryRow(useridInsertQuery,
		uid.UserName,
		uid.UserId,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) GetUserIdFromUserId(userid string) (*models.UserId, error) {
	var uid = &models.UserId{}
	rows, err := r.client.Query(useridfromuseridRetrieveQuery, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&uid.UserName,
			&uid.UserId,
		)
		if err != nil {
			return nil, err
		}
	}
	return uid, nil
}

func (r *mysqlRepo) GetUserIdFromUserName(username string) (*models.UserId, error) {
	var uid = &models.UserId{}
	rows, err := r.client.Query(useridfromusernameRetrieveQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&uid.UserName,
			&uid.UserId,
		)
		if err != nil {
			return nil, err
		}
	}
	return uid, nil
}

func (r *mysqlRepo) AddGame(g *models.Game) error {
	//create default/empty board and unusedpieces rows for use in game
	rs1, err := r.client.Exec(createEmptyBoardQuery)
	if err != nil {
		return err
	}
	bid, err := rs1.LastInsertId()
	if err != nil {
		return err
	}
	rs2, err := r.client.Exec(createEmptyUnusedPiecesQuery)
	if err != nil {
		return err
	}
	upid, err := rs2.LastInsertId()
	if err != nil {
		return err
	}
	//add new game to database
	rand.Seed(time.Now().Unix())
	rand.Intn(16) //random next piece TODO: use in the below query
	err = r.client.QueryRow(
		`INSERT INTO Games (GameId, ActivityStatus, NextPlayer, BoardId, UnusedPiecesId, NextPiece) VALUES (?, ?, ?, ?, ?, ?);`,
		g.GameId,
		g.ActivityStatus,
		g.NextPlayer.UserName,
		bid,
		upid,
		7,
	).Err()
	if err != nil {
		return err
	}
	//invite game creator to the game
	err = r.client.QueryRow(
		`INSERT INTO InvitedPlayers (GameId, UserName) VALUE (?, ?);`,
		g.GameId,
		g.InvitedPlayers[0].UserName,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) GetGame(gameid string) (*models.Game, error) {
	//load basic game data
	g := &models.Game{Board: models.EmptyBoard, UnusedPieces: models.EmptyPieces}
	rows, err := r.client.Query(
		`SELECT
			GameId,
			ActivityStatus,
			NextPlayer,
			NextPiece,
			Winner
		FROM Games WHERE GameId = ?;`,
		gameid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var np string //next player
	var npid int  //next piece id
	for rows.Next() {
		err = rows.Scan(
			&g.GameId,
			&g.ActivityStatus,
			&np,
			&npid,
			&g.Winner,
		)
		if err != nil {
			return nil, err
		}
	}
	g.NextPiece = models.AllQuartoPieces[npid]
	npuid, err := r.GetUserIdFromUserName(np)
	if err != nil {
		return nil, err
	}
	g.NextPlayer = npuid
	//load invitedplayers
	var ipuname string
	rows, err = r.client.Query(
		`SELECT UserName
			FROM InvitedPlayers AS ip
			JOIN Games AS g
			ON ip.GameID = g.GameID
		WHERE g.GameID = ?;`,
		g.GameId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ipuname)
		if err != nil {
			return nil, err
		}
		uid, err := r.GetUserIdFromUserName(ipuname)
		if err != nil {
			return nil, err
		}
		g.InvitedPlayers = append(g.InvitedPlayers, uid)
	}
	//load activeplayers
	var apuname string
	rows, err = r.client.Query(
		`SELECT UserName
			FROM ActivePlayers AS ip
			JOIN Games AS g
			ON ip.GameID = g.GameID
		WHERE g.GameID = ?;`,
		g.GameId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&apuname)
		if err != nil {
			return nil, err
		}
		uid, err := r.GetUserIdFromUserName(apuname)
		if err != nil {
			return nil, err
		}
		g.ActivePlayers = append(g.ActivePlayers, uid)
	}
	//load board data
	rows, err = r.client.Query(
		`SELECT b.*
			FROM Boards AS b
			JOIN Games AS g
			ON b.BoardID = g.BoardID
		WHERE g.GameID = ?;`,
		g.GameId,
	)
	if err != nil {
		return nil, err
	}
	var bid int
	for rows.Next() {
		err = rows.Scan(
			&bid,
			&g.Board[0][0].Id,
			&g.Board[0][1].Id,
			&g.Board[0][2].Id,
			&g.Board[0][3].Id,
			&g.Board[1][0].Id,
			&g.Board[1][1].Id,
			&g.Board[1][2].Id,
			&g.Board[1][3].Id,
			&g.Board[2][0].Id,
			&g.Board[2][1].Id,
			&g.Board[2][2].Id,
			&g.Board[2][3].Id,
			&g.Board[3][0].Id,
			&g.Board[3][1].Id,
			&g.Board[3][2].Id,
			&g.Board[3][3].Id,
		)
		if err != nil {
			return nil, err
		}
	}
	//TODO: load unusedpieces
	rows, err = r.client.Query(
		`SELECT up.*
			FROM UnusedPieces AS up
			JOIN Boards AS b
			ON up.UnusedPiecesID = b.BoardID
		WHERE b.BoardID = ?;`,
		bid,
	)
	if err != nil {
		return nil, err
	}
	var upid int
	for rows.Next() {
		err = rows.Scan(
			&upid,
			&g.UnusedPieces[0].Id,
			&g.UnusedPieces[1].Id,
			&g.UnusedPieces[2].Id,
			&g.UnusedPieces[3].Id,
			&g.UnusedPieces[4].Id,
			&g.UnusedPieces[5].Id,
			&g.UnusedPieces[6].Id,
			&g.UnusedPieces[7].Id,
			&g.UnusedPieces[8].Id,
			&g.UnusedPieces[9].Id,
			&g.UnusedPieces[10].Id,
			&g.UnusedPieces[11].Id,
			&g.UnusedPieces[12].Id,
			&g.UnusedPieces[13].Id,
			&g.UnusedPieces[14].Id,
			&g.UnusedPieces[15].Id,
		)
		if err != nil {
			println("oof")
			return nil, err
		}
	}
	for i, up := range g.UnusedPieces {
		if up.Id > -1 {
			g.UnusedPieces[i] = models.AllQuartoPieces[up.Id]
		} else {
			g.UnusedPieces[i] = models.EmptyQuartoPiece
		}
	}
	return g, nil
}

func (r *mysqlRepo) GetAllGames() ([]*models.Game, error) {
	return nil, nil
}

//TODO: check if piece is being placed in a correct place and has not already been placed
func (r *mysqlRepo) ChangeGame(g *models.Game, gm *models.GameMove) error {
	//board
	var bid int = -1
	err := r.client.QueryRow(`SELECT BoardID FROM Games WHERE GameID = ?;`, g.GameId).Scan(&bid)
	if err != nil || bid == -1 {
		return err
	}
	err = r.client.QueryRow(boardUpdateQuery,
		&g.Board[0][0].Id,
		&g.Board[0][1].Id,
		&g.Board[0][2].Id,
		&g.Board[0][3].Id,
		&g.Board[1][0].Id,
		&g.Board[1][1].Id,
		&g.Board[1][2].Id,
		&g.Board[1][3].Id,
		&g.Board[2][0].Id,
		&g.Board[2][1].Id,
		&g.Board[2][2].Id,
		&g.Board[2][3].Id,
		&g.Board[3][0].Id,
		&g.Board[3][1].Id,
		&g.Board[3][2].Id,
		&g.Board[3][3].Id,
		&bid,
	).Err()
	if err != nil {
		return err
	}
	//get id of unusedpieces row associated with this game/board
	rows, err := r.client.Query(
		`SELECT UnusedPiecesId FROM Games WHERE BoardID = ` + strconv.Itoa(bid) + `;`,
	)
	if err != nil {
		return err
	}
	var upid int
	for rows.Next() {
		err = rows.Scan(&upid)
		if err != nil {
			return err
		}
	}
	//remove piece played from unusedpieces
	npid := gm.NextPiece.Id
	err = r.client.QueryRow(
		`UPDATE UnusedPieces SET up`+strconv.Itoa(npid)+` = -1 WHERE UnusedPiecesID = ?;`,
		upid,
	).Err()
	if err != nil {
		return err
	}
	//update next piece
	err = r.client.QueryRow(
		`UPDATE Games SET NextPiece = `+strconv.Itoa(npid)+` WHERE GameID = ?;`,
		g.GameId,
	).Err()
	if err != nil {
		return err
	}
	//update next player and piece or declare winner
	if g.Winner != nil {
		err = r.client.QueryRow(gameUpdateQueryWithWinner,
			g.ActivityStatus,
			g.Winner.UserName,
			g.GameId,
		).Err()
		if err != nil {
			return err
		}
	} else {
		err = r.client.QueryRow(gameUpdateQuery,
			g.NextPlayer.UserName,
			g.NextPiece.Id,
			g.GameId,
		).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *mysqlRepo) InviteUser(userid string, gameid string) error {
	uid, err := r.GetUserIdFromUserId(userid)
	if err != nil {
		return err
	}
	err = r.client.QueryRow(`SELECT GameID FROM Games WHERE GameID = ?`, gameid).Err()
	if err != nil {
		return err
	}
	err = r.client.QueryRow(
		`INSERT INTO InvitedPlayers (GameID, UserName) VALUE (?, ?);`,
		gameid,
		uid.UserName,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) JoinUser(userid string, gameid string) error {
	uid, err := r.GetUserIdFromUserId(userid)
	if err != nil {
		return err
	}
	err = r.client.QueryRow(`SELECT GameId FROM Games WHERE GameId = ?`, gameid).Err()
	if err != nil {
		return err
	}
	//TODO: check if there is an open spot
	err = r.client.QueryRow(
		`INSERT INTO ActivePlayers (GameId, UserName) VALUE (?, ?);`,
		gameid,
		uid.UserName,
	).Err()
	if err != nil {
		return err
	}
	return nil
}
