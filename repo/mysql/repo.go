package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
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
	rs, err := r.client.Exec(createEmptyBoardQuery)
	if err != nil {
		return err
	}
	bid, err := rs.LastInsertId()
	if err != nil {
		return err
	}
	err = r.client.QueryRow(
		`INSERT INTO Games (GameId, ActivityStatus, BoardId) VALUES (?, ?, ?);`,
		g.GameId,
		g.ActivityStatus,
		bid,
	).Err()
	if err != nil {
		return err
	}
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
	g := &models.Game{}
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
	for rows.Next() {
		err = rows.Scan(
			&g.GameId,
			&g.ActivityStatus,
			&g.NextPlayer,
			&g.NextPiece,
			&g.Winner,
		)
		if err != nil {
			return nil, err
		}
	}
	var uname string
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
		err = rows.Scan(&uname)
		if err != nil {
			return nil, err
		}
		uid, err := r.GetUserIdFromUserName(uname)
		if err != nil {
			return nil, err
		}
		g.InvitedPlayers = append(g.InvitedPlayers, uid)
	}
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
		err = rows.Scan(&uname)
		if err != nil {
			return nil, err
		}
		uid, err := r.GetUserIdFromUserName(uname)
		if err != nil {
			return nil, err
		}
		g.ActivePlayers = append(g.ActivePlayers, uid)
	}
	rows, err = r.client.Query(
		`SELECT b.*
			FROM Boards AS b
			JOIN Games AS g
			ON b.BoardID = g.BoardID
		WHERE g.GameID = ?;`,
		g.GameId,
	)
	var bid int
	for rows.Next() {
		err = rows.Scan(
			&bid,
			&g.Board[0][0],
			&g.Board[0][1],
			&g.Board[0][2],
			&g.Board[0][3],
			&g.Board[1][0],
			&g.Board[1][1],
			&g.Board[1][2],
			&g.Board[1][3],
			&g.Board[2][0],
			&g.Board[2][1],
			&g.Board[2][2],
			&g.Board[2][3],
			&g.Board[3][0],
			&g.Board[3][1],
			&g.Board[3][2],
			&g.Board[3][3],
		)
		if err != nil {
			//println("wowza batman, an error!: " + err.Error())
			return nil, err
		}
	}
	return g, nil
}

func (r *mysqlRepo) GetAllGames() ([]*models.Game, error) {
	return nil, nil
}

func (r *mysqlRepo) ChangeGame(g *models.Game) error {
	//board
	var bid int = -1
	err := r.client.QueryRow(`SELECT BoardID FROM Games WHERE GameID = ?`, g.GameId).Scan(&bid)
	if err != nil || bid == -1 {
		return err
	}
	err = r.client.QueryRow(`UPDATE Boards
		SET x0y0 = ?,
			x0y1 = ?,
			x0y2 = ?,
			x0y3 = ?,
			x1y0 = ?,
			x1y1 = ?,
			x1y2 = ?,
			x1y3 = ?,
			x2y0 = ?,
			x2y1 = ?,
			x2y2 = ?,
			x2y3 = ?,
			x3y0 = ?,
			x3y1 = ?,
			x3y2 = ?,
			x3y3 = ?
		WHERE BoardID = ?`,
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
	err = r.client.QueryRow(gameUpdateQuery,
		g.ActivityStatus,
		g.NextPlayer.UserName,
		g.NextPiece.Id,
		g.Winner.UserName,
	).Err()
	if err != nil {
		return err
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
