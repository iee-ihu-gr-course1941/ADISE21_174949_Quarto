#!/bin/sh -x

# base URL
BASE_URL="localhost:8000"
#BASE_URL="https://users.iee.ihu.gr/~it174949/index.php"

# make user, get uid json
U1=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "user", "password": "verybigsecret"}'\
	${BASE_URL}/user)

U1UN=$(echo ${U1} | jq -r '.username')
U1UID=$(echo ${U1} | jq -r '.user_id')

# make game, get raw gid
GID=$(curl -s -X POST -H "Content-Type: application/json" -d "${U1}" ${BASE_URL}/game | jq -r '.game_id')

LNK="${BASE_URL}/game/${GID}"

G=$(curl -s "${LNK}")

# make another user, get raw name
U2=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "u2", "password": "hugesecret"}'\
	${BASE_URL}/user)

U2UN=$(echo ${U2} | jq -r '.username')
U2UID=$(echo ${U2} | jq -r '.user_id')

LNK="${BASE_URL}/game/${GID}/invite/${U2UN}"

INV=$(curl -s -X POST -H "Content-Type: application/json" -d "${U1}" "${LNK}")

LNK="${BASE_URL}/game/${GID}"

G=$(curl -s "${LNK}")

#change link to join
LNK="${BASE_URL}/game/${GID}/join"

JOIN_RES1=$(curl -s -X POST -H "Content-Type: application/json" -d "${U1}" "${LNK}")

JOIN_RES2=$(curl -s -X POST -H "Content-Type: application/json" -d "${U2}" "${LNK}")

#change link to play
LNK="${BASE_URL}/game/${GID}/play"

PLAY_DATA='{"username":"user","user_id":"U1UID", "position_x":3, "position_y":2, "next_piece":{"Id":2,"Dark":true,"Short":false,"Hollow":false,"Round":false}}'

PLAY_DATA=$(echo $PLAY_DATA | sed s/U1UID/${U1UID}/g)

PLAY_RES=$(curl -s -X POST\
	-H 'Content-Type: application/json'\
	-d "${PLAY_DATA}"\
	"${LNK}")
