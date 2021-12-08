package mysql

import (
	"encoding/json"
	"log"
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
	jip, err := json.Marshal(g.InvitedPlayers)
	if err != nil {
		return err
	}
	jb, err := json.Marshal(g.State.Board)
	if err != nil {
		return err
	}
	jup, err := json.Marshal(g.State.UnusedPieces)
	if err != nil {
		return err
	}
	err = r.client.QueryRow(gameInsertQuery,
		g.GameId,
		g.ActivityStatus,
		jip,
		jb,
		jup,
	).Err()
	log.Println("mysql game", g)
	if err != nil {
		return err
	}
	return nil
}

func (r *mysqlRepo) GetGame(gameid string) (*models.Game, error) {
	g := &models.Game{
		GameId:         gameid,
		ActivityStatus: true,
		State: &models.GameState{
			Board:        models.EmptyBoard,
			UnusedPieces: models.AllQuartoPieces,
		},
	}
	rows, err := r.client.Query(gameRetrieveQuery, gameid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jap, jip, jnp, jb, jup []byte
	for rows.Next() {
		err = rows.Scan(
			&g.GameId,
			&g.ActivityStatus,
			&g.Winner,
			&jap,
			&jip,
			&g.State.NextPlayer,
			&jnp,
			&jb,
			&jup,
		)
		if err != nil {
			return nil, err
		}
	}
	_ = json.Unmarshal(jap, &g.ActivePlayers)
	_ = json.Unmarshal(jip, &g.InvitedPlayers)
	_ = json.Unmarshal(jnp, &g.State.NextPiece)
	_ = json.Unmarshal(jb, &g.State.Board)
	_ = json.Unmarshal(jup, &g.State.UnusedPieces)
	return g, nil
}

func (r *mysqlRepo) GetGameState(gameid string) (*models.GameState, error) {
	g := &models.Game{
		GameId:         gameid,
		ActivityStatus: true,
		State: &models.GameState{
			Board:        models.EmptyBoard,
			UnusedPieces: models.AllQuartoPieces,
		},
	}
	var gsb []byte
	var gsup []byte
	rows, err := r.client.Query(gamestateRetrieveQuery, gameid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&g.State.NextPlayer,
			&g.State.NextPiece,
			&gsb,
			&gsup,
		)
		if err != nil {
			return nil, err
		}
	}
	_ = json.Unmarshal(gsb, &g.State.Board)
	_ = json.Unmarshal(gsup, &g.State.UnusedPieces)
	return g.State, nil
}

func (r *mysqlRepo) GetAllGames() ([]*models.Game, error) {
	return nil, nil
}

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
			&g.State.NextPlayer,
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
