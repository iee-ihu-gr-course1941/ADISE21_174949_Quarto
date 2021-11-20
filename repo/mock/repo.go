package mock

import (
	"fmt"
)

type MockDB struct {
	Users   []*User
	UserIds []*UserId
	Games   []*Game
}

var mymockdb MockDB = nil

func NewMockDB() (MockDB, error) {
	mymockdb = &MockDB{}
	return &MockDB{}, nil
}

func (m *MockDB) AddUser(u *User) error {
	m.Users = append(m.Users, u)
	return nil
}

func (m *MockDB) AddUserId(uid *UserId) error {
	m.UserIds = append(m.UserIds, uid)
	return nil
}

func (m *MockDB) GetUserId(userid string) (*UserId, error) {
	for _, u := range m.UserIds {
		if u.UserName == userid {
			return u, nil
		}
	}
	return nil, fmt.Error("user with id", userid, "not found")
}

func (m *MockDB) AddGame(g *Game) error {
	m.Games = append(m.Games, g)
	return nil
}

func (m *MockDB) GetGame(gameid string) (*Game, error) {
	for _, g := range m.Games {
		if g.GameId == gameid {
			return g, nil
		}
	}
	return nil, fmt.Error("game with id", gameid, "not found")
}

func (m *MockDB) GetGameState(gameid string) (*GameState, error) {
	for _, g := range m.Games {
		if g.GameId == gameid {
			return g.State, nil
		}
	}
	return nil, fmt.Error("game with id", gameid, "not found")
}

func (m *MockDB) GetAllGames() ([]*Game, error) {
	return m.Games, nil
}

func (m *MockDB) InviteUser(userid string, gameid string) error {
	u, err := GetUserId(userid)
	if err != nil {
		return err
	}
	g, err := GetGame(gameid)
	if err != nil {
		return err
	}
	g.InvitedUsers = append(g.InvitedUsers, u)
	return nil
}

func (m *MockDB) JoinUser(userid string, gameid string) error {
	u, err := GetUserId(userid)
	if err != nil {
		return err
	}
	g, err := GetGame(gameid)
	if err != nil {
		return err
	}
	for _, ip := range g.InvitedPlayers {
		if cap(g.ActivePlayers) == MaxPlayers {
			return fmt.Error("couldn't join because game is full")
		} else if cap(g.ActivePlayers) > MaxPlayers {
			return fmt.Error("I honestly don't know how this happened")
		} else {
			g.ActivePlayers = append(g.ActivePlayers, u)
			g.InvitedPlayers = g.InvitedPlayers[:len(g.InvitedPlayers)-1]
			return nil
		}
	}
}
