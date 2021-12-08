package mysql

import (
	"encoding/json"
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
	//TODO: invite game creator
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
			UnusedPieces,
			Winner
		FROM Games WHERE GameId = ?;
		`,
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
			&g.UnusedPieces,
			&g.Winner,
		)
		if err != nil {
			return nil, err
		}
	}
	return g, nil
}

func (r *mysqlRepo) GetAllGames() ([]*models.Game, error) {
	return nil, nil
}

//TODO: rewrite
func (r *mysqlRepo) InviteUser(userid string, gameid string) error {
	//TODO: check if user exists
	uid, err := r.GetUserIdFromUserName(userid)
	if err != nil {
		return err
	}
	var g = &models.Game{}
	var jap, jip, jnp, jb, jup []byte
	rows, err := r.client.Query(gameRetrieveQuery, gameid)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&g.GameId,
			&g.ActivityStatus,
			&g.Winner,
			&jap,
			&jip,
			&g.NextPlayer,
			&jnp,
			&jb,
			&jup,
		)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(jip, &g.InvitedPlayers)
	if err != nil {
		return err
	}
	g.InvitedPlayers = append(g.InvitedPlayers, uid)
	jip, err = json.Marshal(g.InvitedPlayers)
	if err != nil {
		return err
	}
	err = r.client.QueryRow(gameUpdateQuery, jip).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) JoinUser(userid string, gameid string) error {
	return nil
}
