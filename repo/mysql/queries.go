package mysql

var createUserTableQuery = `CREATE TABLE if not exists Users (
	UserNickname VARCHAR PRIMARY KEY NOT NULL,
	UserPassword VARCHAR NOT NULL,
);`

var createUserIDsTableQuery = `CREATE TABLE if not exists UserIDs (
	UserNickname PRIMARY KEY VARCHAR REFERENCES Users(UserNickname) NOT NULL,
	UserID VARCHAR NOT NULL
);`

var createGameTableQuery = `CREATE TABLE if not exists Games (
	GameID VARCHAR PRIMARY KEY NOT NULL,
	ActivityStatus BOOLEAN NOT NULL DEFAULT FALSE,
	Winner VARCHAR REFERENCES Users(UserNickname)
	--ActivePlayers
	--InvitedPlayers
);`

var createGameStateTableQuery = `CREATE TABLE if not exists GameStates (
	GameID VARCHAR PRIMARY KEY NOT NULL,
	NextPlayer VARCHAR REFERENCES Users(UserNickname),
	--NextPiece
	--Board
	--UnusedPieces
);`
