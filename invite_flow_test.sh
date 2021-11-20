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

G=$(curl -s "$LNK")

echo $G

# make another user, get raw name
U2=$(curl -s -X POST\
	-H "Content-Type: application/json"\
	-d '{"username": "u2", "password": "hugesecret"}'\
	localhost:8000/user | jq -r '.username')

LNK="${LNK}/invite/${U2}"

INV=$(curl -s -X POST -H "Content-Type: application/json" -d "$U" "$LNK")

LNK="localhost:8000/game/${GID}"
echo $(curl -s "$LNK")
