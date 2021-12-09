#!/bin/sh

#go run main.go vars.go

# make user, get uid json
U=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "user", "password": "verybigsecret"}'\
	localhost:8000/user)

# make game, get raw gid
GID=$(curl -s -X POST -H "Content-Type: application/json" -d "$U" localhost:8000/game | jq -r '.game_id')

# print both
echo $U
echo $GID

LNK="localhost:8000/game/${GID}"

G=$(curl -s "${LNK}")

echo $G

# make another user, get raw name
U2=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "u2", "password": "hugesecret"}'\
	localhost:8000/user)

U2UN=$(echo ${U2} | jq -r '.username')

LNK="${LNK}/invite/${U2UN}"

INV=$(curl -s -X POST -H "Content-Type: application/json" -d "$U" "$LNK")

LNK="localhost:8000/game/${GID}"

G=$(curl -s "${LNK}")

#change link to join
LNK="localhost:8000/game/${GID}/join"

JOIN_RES1=$(curl -s -X POST -H "Content-Type: application/json" -d "$U" "$LNK")

echo $JOIN_RES1

JOIN_RES2=$(curl -s -X POST -H "Content-Type: application/json" -d "$U2" "$LNK")

echo $JOIN_RES2
