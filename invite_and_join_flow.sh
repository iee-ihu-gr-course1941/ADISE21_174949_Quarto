#!/bin/sh -x

# base URL
BASE_URL="localhost:8000"
#BASE_URL="https://users.iee.ihu.gr/~it174949/index.php"

# make user, get uid json
U=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "user", "password": "verybigsecret"}'\
	${BASE_URL}/user)

# make game, get raw gid
GID=$(curl -s -X POST -H "Content-Type: application/json" -d "${U}" ${BASE_URL}/game | jq -r '.game_id')

# print both
echo $U
echo $GID

LNK="${BASE_URL}/game/${GID}"

G=$(curl -s "${LNK}")

echo $G

# make another user, get raw name
U2=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "u2", "password": "hugesecret"}'\
	${BASE_URL}/user)

U2UN=$(echo ${U2} | jq -r '.username')

LNK="${BASE_URL}/game/${GID}/invite/${U2UN}"

INV=$(curl -s -X POST -H "Content-Type: application/json" -d "${U}" "${LNK}")

LNK="${BASE_URL}/game/${GID}"

G=$(curl -s "${LNK}")

#change link to join
LNK="${BASE_URL}/game/${GID}/join"

JOIN_RES1=$(curl -s -X POST -H "Content-Type: application/json" -d "${U}" "${LNK}")

echo $JOIN_RES1

JOIN_RES2=$(curl -s -X POST -H "Content-Type: application/json" -d "${U2}" "${LNK}")

echo $JOIN_RES2
