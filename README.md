# Quarto Go API
REST API written in Go to play Quarto

You can also see [extra documentation](DOCS.md).

## Technology Used
The following is a non-exhaustive list of technology used to build this
- `Go` as the language
- `gorilla/mux` for setting up routes
- built-in Go testing harness

## Running Application
from the root of the repository run:
```
go run main.go
```
## Running with test MySQL database
I use docker for this:
```bash
docker run --rm --name=mysql -p 3306:3306 -e MYSQL_INITDB_SKIP_TZINFO=yes -e MYSQL_USER=tester -e MYSQL_PASSWORD=Apasswd -e MYSQL_DATABASE=tester -e MYSQL_ROOT_PASSWORD=rootApasswd mysql:5
```

and then just:
```bash
MYSQL_URL="test" go run main.go
```

## Running tests
### For locally-running application
run the application and in another window run `./play.sh`

### For remote-running application
change the value of `BASE_URL` inside `play.sh` and then run it with `./play.sh`

# Usage
how to use with curl

## Register User
run:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "password": "verybigsecret"}'\
	localhost:8000/user
```

returns:
```json
{"username":"someuser","user_id":"6WZw9BJ7R"}
```

or if you have `jq` installed you can extract the `user_id` value using:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "password": "verybigsecret"}'\
	localhost:8000/user | jq '.user_id'
```

## Create Game
run:
```bash
curl -X POST \
	-H "Content-Type: application/json"\
	-d '{"username":"someuser","user_id":"6WZw9BJ7R"}'\
	localhost:8000/game
```

returns:
```json
{"game_id":"t3Rurf1ng","active_players":null,"invited_players":[{"username":"someuser","user_id":"6WZw9BJ7R"}],"activity_status":true,"next_player":{"username":"someuser","user_id":"6WZw9BJ7R"},"next_piece":{"Id":7,"Dark":false,"Short":true,"Hollow":false,"Round":true},"board":[[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}]],"unused_pieces":[{"Id":0,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":1,"Dark":true,"Short":false,"Hollow":false,"Round":false},{"Id":2,"Dark":false,"Short":true,"Hollow":false,"Round":false},{"Id":3,"Dark":false,"Short":false,"Hollow":true,"Round":false},{"Id":4,"Dark":false,"Short":false,"Hollow":false,"Round":true},{"Id":5,"Dark":true,"Short":true,"Hollow":false,"Round":false},{"Id":6,"Dark":false,"Short":true,"Hollow":true,"Round":false},{"Id":7,"Dark":false,"Short":true,"Hollow":false,"Round":true},{"Id":8,"Dark":true,"Short":false,"Hollow":true,"Round":false},{"Id":9,"Dark":true,"Short":false,"Hollow":false,"Round":true},{"Id":10,"Dark":false,"Short":false,"Hollow":true,"Round":true},{"Id":11,"Dark":false,"Short":true,"Hollow":true,"Round":true},{"Id":12,"Dark":true,"Short":false,"Hollow":true,"Round":true},{"Id":13,"Dark":true,"Short":true,"Hollow":false,"Round":true},{"Id":14,"Dark":true,"Short":true,"Hollow":true,"Round":false},{"Id":15,"Dark":true,"Short":true,"Hollow":true,"Round":true}],"winner":null}
```

## Invite User to Game
run:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "user_id": "6WZw9BJ7R"}'\
	localhost:8000/game/t3Rurf1ng/invite/otheruser
```

returns:
```json
{"message": "success"}
```

## Join Game
User 1:

run:
```bash

 curl -X POST\
 	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "user_id": "6WZw9BJ7R"}'\
	localhost:8000/game/t3Rurf1ng/join
 ```

 returns:
 ```json
 {"message": "success"}
 ```

User 2:
run:
```bash

 curl -X POST\
 	-H "Content-Type: application/json"\
	-d '{"username": "otheruser", "user_id": "-63C9B1nR"}'\
	localhost:8000/game/t3Rurf1ng/join
 ```

 returns:
 ```json
 {"message": "success"}
 ```

## Make a Move in Game
run:
```bash
curl -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "someuser", "user_id": "6WZw9BJ7R", "position_x":0, "position_y":2, "next_piece":{"Id":13}}'\
	localhost:8000/game/t3Rurf1ng/play
```

returns:
```json
{"game_id":"t3Rurf1ng","active_players":[{"username":"otheruser","user_id":"-63C9B1nR"},{"username":"someuser","user_id":"6WZw9BJ7R"}],"invited_players":[{"username":"otheruser","user_id":"-63C9B1nR"},{"username":"someuser","user_id":"6WZw9BJ7R"}],"activity_status":true,"next_player":{"username":"otheruser","user_id":"-63C9B1nR"},"next_piece":{"Id":13,"Dark":true,"Short":true,"Hollow":false,"Round":true},"board":[[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":7,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}],[{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false}]],"unused_pieces":[{"Id":0,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":1,"Dark":true,"Short":false,"Hollow":false,"Round":false},{"Id":2,"Dark":false,"Short":true,"Hollow":false,"Round":false},{"Id":3,"Dark":false,"Short":false,"Hollow":true,"Round":false},{"Id":4,"Dark":false,"Short":false,"Hollow":false,"Round":true},{"Id":5,"Dark":true,"Short":true,"Hollow":false,"Round":false},{"Id":6,"Dark":false,"Short":true,"Hollow":true,"Round":false},{"Id":-1,"Dark":false,"Short":false,"Hollow":false,"Round":false},{"Id":8,"Dark":true,"Short":false,"Hollow":true,"Round":false},{"Id":9,"Dark":true,"Short":false,"Hollow":false,"Round":true},{"Id":10,"Dark":false,"Short":false,"Hollow":true,"Round":true},{"Id":11,"Dark":false,"Short":true,"Hollow":true,"Round":true},{"Id":12,"Dark":true,"Short":false,"Hollow":true,"Round":true},{"Id":13,"Dark":true,"Short":true,"Hollow":false,"Round":true},{"Id":14,"Dark":true,"Short":true,"Hollow":true,"Round":false},{"Id":15,"Dark":true,"Short":true,"Hollow":true,"Round":true}],"winner":null}
```
