package main

import (
	"bytes"
	"encoding/json"
	rd "github.com/Pallinder/go-randomdata"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test server
var testServer *httptest.Server = httptest.NewServer(setupRouter())

// Function for creating a user for use only outside TestCreateUser
func randomUserCreation(t *testing.T) *UserId {
	// define URL
	testURL := testServer.URL + "/user"
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "` + rd.SillyName() + `", "password": "mypasswd"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// response should contain json that can maps to the UserId type
	u := &UserId{}
	// try to unmarshal
	err = json.Unmarshal(body, u)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	return u
}

// Function for creating a user for use only outside TestCreateUser
func userCreation(t *testing.T) *UserId {
	// define URL
	testURL := testServer.URL + "/user"
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "myself", "password": "mypasswd"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// response should contain json that can maps to the UserId type
	u := &UserId{}
	// try to unmarshal
	err = json.Unmarshal(body, u)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	return u
}

// Test creating a user
func TestCreateUser(t *testing.T) {
	WipeState()
	// define URL
	testURL := testServer.URL + "/user"
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "myself", "password": "mypasswd"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// log response
	t.Log(res)

	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// log response body
	t.Log(string(body))

	// response should contain json that can maps to the UserId type
	// set up empty UserId
	u := &UserId{}
	// try to unmarshal
	err = json.Unmarshal(body, u)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	// log UserId
	t.Log(u)
}

// Function for creating a game for use only outside TestCreateUser
func gameCreation(t *testing.T) *Game {
	// create a user
	u := randomUserCreation(t)
	// change URL
	testURL := testServer.URL + "/game"
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"user_id": "` + u.UserId + `"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}

	// response should contain json that can maps to the Game type
	// set up empty Game
	g := &Game{}
	// try to unmarshal
	err = json.Unmarshal(body, g)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	// check the amount of invited players
	if len(g.InvitedPlayers) > 1 {
		t.Error("more than 1 player is invited to the game")
		t.Log(g.InvitedPlayers[0], g.InvitedPlayers[1])
	} else if len(g.InvitedPlayers) < 1 {
		t.Error("less than 1 player is invited to the game")
	}
	firstInvPlayer := g.InvitedPlayers[0].UserName
	// first invited player should be the one we created
	if firstInvPlayer != u.UserName {
		t.Error("expected first invited player is not who they should be")
	}
	t.Log("gameCreation player 0", g.InvitedPlayers[0])
	return g
}

// Test creating a game
func TestCreateGame(t *testing.T) {
	WipeState()
	// create a user
	u := userCreation(t)
	// change URL
	testURL := testServer.URL + "/game"
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "` + u.UserName + `", "user_id": "` + u.UserId + `"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// log response
	t.Log(res)

	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// log response body
	t.Log(string(body))

	// response should contain json that can maps to the Game type
	// set up empty Game
	g := &Game{}
	// try to unmarshal
	err = json.Unmarshal(body, g)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	// log currently invited users (should only be user "myself")
	if len(g.InvitedPlayers) > 1 {
		t.Error("more than 1 player is invited to the game")
	} else if len(g.InvitedPlayers) < 1 {
		t.Error("less than 1 player is invited to the game")
	}
	t.Log(g.InvitedPlayers)
	firstInvPlayer := g.InvitedPlayers[0].UserName
	if firstInvPlayer != "myself" {
		t.Error("expected first invited player is not who they should be")
	}

	// log Game
	t.Log(g)
}
