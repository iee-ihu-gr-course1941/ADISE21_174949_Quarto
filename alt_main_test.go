package main

import (
	"bytes"
	rd "github.com/Pallinder/go-randomdata"
	"net/http"
	"strconv"
	"testing"
)

func TestWinInGame(t *testing.T) {
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

	//TODO: make u2 win and u not interfere
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
	str2 := `{"username": "` + u2.UserName + `", "user_id": "` + u2.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":15}` + `}`
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
	str2 = `{"username": "` + u2.UserName + `", "user_id": "` + u2.UserId + `", ` + `"position_x":` + strconv.Itoa(rd.Number(4)) + `, ` + `"position_y":` + strconv.Itoa(rd.Number(4)) + `, ` + `"next_piece": {"Id":14}` + `}`
	jsonPlayData2 = []byte(str2)
	// do a simple Post request with the above data
	res, err = http.Post(testURL, "application/json", bytes.NewBuffer(jsonPlayData2))
	// check for request errors
	if err != nil {
		t.Error("POST error:", err)
	}
	// be responsible and close the response some time
	defer res.Body.Close()
}
