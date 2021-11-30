package mysql

import (
	"database/sql"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlRepo struct {
	client *sql.DB
	mysqlURL  string
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
	return client, nil
}

func NewMysqlRepo(url string) (*mysqlRepo, error) {
	mysqlclient, err := newMysqlClient(url)
	if err != nil {
		return nil, err
	}
	repo := &mysqlRepo{
		mysqlURL:  url,
		client: mysqlclient,
	}
	return repo, nil
}

func (r *mysqlRepo) AddUser(u *models.User) error {
	return nil
}

func (r *mysqlRepo) AddUserId(uid *models.UserId) error {
	return nil
}

func (r *mysqlRepo) GetUserIdFromUserId(userid string) (*models.UserId, error) {
	return nil, nil
}

func (r *mysqlRepo) GetUserIdFromUserName(username string) (*models.UserId, error) {
	return nil, nil
}

func (r *mysqlRepo) AddGame(g *models.Game) error {
	return nil
}

func (r *mysqlRepo) GetGame(gameid string) (*models.Game, error) {
	return nil, nil
}

func (r *mysqlRepo) GetGameState(gameid string) (*models.GameState, error) {
	return nil, nil
}

func (r *mysqlRepo) GetAllGames() ([]*models.Game, error) {
	return nil, nil
}

func (r *mysqlRepo) InviteUser(userid string, gameid string) error {
	return nil
}

func (r *mysqlRepo) JoinUser(userid string, gameid string) error {
	return nil
}
