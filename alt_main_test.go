package main

import (
	"bytes"
	"github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto/models"
	"encoding/json"
	"io"
	"net/http"
	"runtime"
	"testing"
)

//TODO: figure out why the following happens
//
//=== RUN   TestAllReal
//    alt_main_test.go:92: gID yqT5bkp7Rz
//    alt_main_test.go:170: p 0 myself sqo5bkt7R
//--- PASS: TestAllReal (0.00s)
//=== RUN   TestGetGame
//    alt_main_test.go:207: p 0 myself sqo5bkt7R
//    alt_main_test.go:207: p 1 u222 -AlqJzt7Rz
//--- PASS: TestGetGame (0.00s)

var gid string

func TestAllReal(t *testing.T) {
	t.SkipNow()
	// clear the global storage
	WipeState()
	// define URL
	testURL := "http://localhost:8000" + "/user"
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

	testURL = "http://localhost:8000" + "/game"
	// create some data in the form of an io.Reader from a string of json
	jsonData = []byte(`{"username": "` + u.UserName + `", "user_id": "` + u.UserId + `"}`)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	// save response body to check later
	body, err = io.ReadAll(res.Body)
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
	//t.Log("game res", string(body))
	gid = g.GameId
	// log currently invited users (should only be user "myself")
	if len(g.InvitedPlayers) > 1 {
		t.Error("more than 1 player is invited to the game")
	} else if len(g.InvitedPlayers) < 1 {
		t.Error("less than 1 player is invited to the game")
	}
	firstInvPlayer := g.InvitedPlayers[0].UserName
	if firstInvPlayer != "myself" {
		t.Error("expected first invited player is not who they should be")
	}
	t.Log("gID", g.GameId)

	// define URL
	testURL = "http://localhost:8000" + "/user"
	// create some data in the form of an io.Reader from a string of json
	jsonData = []byte(`{"username": "u222", "password": "mypass"}`)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
	// save response body to check later
	body, err = io.ReadAll(res.Body)
	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// response should contain json that can maps to the UserId type
	u2 := &models.UserId{}
	// try to unmarshal
	err = json.Unmarshal(body, u2)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
	}

	testURL = "http://localhost:8000" + "/game/" + g.GameId + "/invite/" + u2.UserName
	jsonData = []byte(`{"username": "` + u.UserName + `", "user_id": "` + u.UserId + `"}`)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonData))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	// save response body to check later
	body, err = io.ReadAll(res.Body)

	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// check if body has success message
	if string(body) != MsgSuccess {
		t.Error("inviting user did not yield success message")
	}

	runtime.GC()

	for i, p := range g.InvitedPlayers {
		t.Log("p", i, p.UserName, p.UserId)
	}

	if len(g.InvitedPlayers) <= 1 && cap(g.InvitedPlayers) <= 1 {
		t.Error("second player wasn't added to the invitation list")
		t.Log(g.InvitedPlayers[1])
	}
}

/*
func TestGetGame(t *testing.T) {
	t.SkipNow()
	// clear the global storage
	WipeState()
	// define URL
	testURL := "http://localhost:8000" + "/game/" + gid
	res, err := http.Get(testURL)
	// check for request errors
	if err != nil {
		t.Error("GET error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()

	// save response body to check later
	body, err := io.ReadAll(res.Body)

	// check for response body read errors
	if err != nil {
		t.Error("resp.Body error:", err)
	}
	// set up empty Game
	g := &Game{}
	// try to unmarshal
	err = json.Unmarshal(body, g)
	// check for unmarshaling errors
	if err != nil {
		t.Error("unmarshal error:", err)
		t.Log("res", res)
		t.Log("bod", string(body))
	}
	for i, p := range g.InvitedPlayers {
		t.Log("p", i, p.UserName, p.UserId)
	}
}
*/
