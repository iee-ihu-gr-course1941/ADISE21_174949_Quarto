package main

import (
	"bytes"
	"encoding/json"
	rd "github.com/Pallinder/go-randomdata"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// Test HTTP server without logging middleware
var testServer *httptest.Server = httptest.NewServer(setupRouter(false))

// Function for creating a user for use only outside TestCreateUser
func randomUserCreation(t *testing.T) *models.UserId {
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
	u := &models.UserId{}
	// try to unmarshal
	err = json.Unmarshal(body, u)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	return u
}

// Function for creating a user for use only outside TestCreateUser
func userCreation(t *testing.T) *models.UserId {
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
	u := &models.UserId{}
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
	u := &models.UserId{}
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
func gameCreation(t *testing.T) *models.Game {
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
	g := &models.Game{}
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
	// create a user
	u := randomUserCreation(t)
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
	g := &models.Game{}
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
	if firstInvPlayer != u.UserName {
		t.Error("expected first invited player is not who they should be")
	}

	// log Game
	t.Log(g)
}

func gameInvitation(t *testing.T) (*models.Game, *models.UserId, *models.UserId) {
	// create a game which also creates random user
	g := gameCreation(t)
	// alias for the first invited player aka the game creator
	invPlayer1 := g.InvitedPlayers[0]
	// create a second random user
	u := randomUserCreation(t)
	// change URL, add the name of the user to be invited
	testURL := testServer.URL + "/game/" + g.GameId + "/invite/" + u.UserName
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "` + invPlayer1.UserName + `", "user_id": "` + invPlayer1.UserId + `"}`)
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
	// check if body has success message
	if string(body) != MsgSuccess {
		t.Error("inviting user did not yield success message")
	}

	// change URL to the game and its ID
	testURL = testServer.URL + "/game/" + g.GameId
	// do a simple Get request to see the game state
	res, err = http.Get(testURL)
	// check for request errors
	if err != nil {
		t.Error("GET error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// save response body to check later
	body, err = io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// try to unmarshal
	err = json.Unmarshal(body, g)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}
	return g, invPlayer1, u
}

func TestInviteToGame(t *testing.T) {
	// create a game which also creates random user
	g := gameCreation(t)
	// alias for the first invited player aka the game creator
	invPlayer1 := g.InvitedPlayers[0]
	// create a second random user
	u := randomUserCreation(t)
	// change URL, add the name of the user to be invited
	testURL := testServer.URL + "/game/" + g.GameId + "/invite/" + u.UserName
	// create some data in the form of an io.Reader from a string of json
	jsonData := []byte(`{"username": "` + invPlayer1.UserName + `", "user_id": "` + invPlayer1.UserId + `"}`)
	// do a simple Post request with the above data
	res, err := http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// log response
	t.Log("res", res)

	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// log response body
	t.Log("body", string(body))
	// check if body has success message
	if string(body) != MsgSuccess {
		t.Error("inviting user did not yield success message")
	}

	// change URL to the game and its ID
	testURL = testServer.URL + "/game/" + g.GameId
	// do a simple Get request to see the game state
	res, err = http.Get(testURL)
	// check for request errors
	if err != nil {
		t.Error("GET error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// log response
	t.Log("res", res)
	// save response body to check later
	body, err = io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	t.Log(string(body))
	// try to unmarshal
	err = json.Unmarshal(body, g)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}

	//TODO: replace with QuartoStorage calls
	//if len(testGames[0].InvitedPlayers) <= 1 || cap(testGames[0].InvitedPlayers) <= 1 {
	//	t.Error("second player wasn't added to the invitation list")
	//} else {
	//	t.Log("tG[0].IP[1]", testGames[0].InvitedPlayers[1])
	//	t.Log("g.IP[0]", g.InvitedPlayers[0])
	//	if g.GameId == testGames[0].GameId {
	//		t.Log("same Game ID so the below shouldn't explode since tG[0].IP[1] exists")
	//		t.Log("g.IP[1]", g.InvitedPlayers[1])
	//	}
	//}
}

func TestJoinGame(t *testing.T) {
	g, u, u2 := gameInvitation(t)
	testURL := testServer.URL + "/game/" + g.GameId + "/join"
	t.Log(testURL)
	t.Log(u, u2)
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
	t.Log("res", res)

	// save response body to check later
	body, err := io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// log response body
	t.Log("body", string(body))
	// check if body has success message
	if string(body) != MsgSuccess {
		t.Error("inviting user did not yield success message")
	}
}

func TestPlayInGame(t *testing.T) {
	g, u, u2 := gameInvitation(t)
	testURL := testServer.URL + "/game/" + g.GameId + "/join"

	//user 1 join game
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

	//user 2 join game
	// create some data in the form of an io.Reader from a string of json
	jsonData = []byte(`{"username": "` + u2.UserName + `", "user_id": "` + u2.UserId + `"}`)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	//user 1 play 1
	str1 := `{"username": "` + u.UserName + `", "user_id": "` + u.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":` + strconv.Itoa(rd.Number(16)) + `}` + `}`
	jsonPlayData1 := []byte(str1)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonPlayData1))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	//user 2 play 1
	str2 := `{"username": "` + u2.UserName + `", "user_id": "` + u2.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":` + strconv.Itoa(rd.Number(16)) + `}` + `}`
	jsonPlayData2 := []byte(str2)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonPlayData2))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	//user 1 play 2
	str1 = `{"username": "` + u.UserName + `", "user_id": "` + u.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":` + strconv.Itoa(rd.Number(16)) + `}` + `}`
	jsonPlayData1 = []byte(str1)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonPlayData1))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	//user 2 play 2
	str2 = `{"username": "` + u2.UserName + `", "user_id": "` + u2.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":` + strconv.Itoa(rd.Number(16)) + `}` + `}`
	jsonPlayData2 = []byte(str2)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonPlayData2))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	t.Log("TestPlayInGame not fully implemented yet")
}
