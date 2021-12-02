package mysql

var createUserTableQuery = `CREATE TABLE if not exists Users (
	UserNickname VARCHAR(100) NOT NULL,
	UserPassword VARCHAR(100) NOT NULL,
	PRIMARY KEY (UserNickname)
);`

var createUserIdTableQuery = `CREATE TABLE if not exists UserIDs (
	UserNickname VARCHAR(100) NOT NULL REFERENCES Users(UserNickname),
	UserID VARCHAR(100) NOT NULL,
	PRIMARY KEY (UserNickname)
);`

var createGameTableQuery = `CREATE TABLE if not exists Games (
	GameID VARCHAR(100) PRIMARY KEY NOT NULL,
	ActivityStatus BOOLEAN NOT NULL DEFAULT FALSE,
	Winner VARCHAR(100) REFERENCES UserIDs(UserNickname),
	ActivePlayers JSON,
	InvitedPlayers JSON,
	NextPlayer VARCHAR(100) REFERENCES Users(UserNickname),
	NextPiece JSON,
	Board JSON NOT NULL,
	UnusedPieces JSON NOT NULL
);`

var createinvactPlayerTablesQuery = `CREATE TABLE if not exists ActivePlayers (
	GameID VARCHAR(100) REFERENCES Games(GameID) NOT NULL,
	UserNickname VARCHAR(100) REFERENCES UserIDs(UserNickname) NOT NULL,
	PRIMARY KEY (GameID)
);
CREATE TABLE if not exists InvitedPlayers (
	GameID VARCHAR(100) REFERENCES Games(GameID) NOT NULL,
	UserNickname VARCHAR(100) REFERENCES UserIDs(UserNickname),
	InvitationTime TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (GameID)
);`


// TODO: model board

var useridfromuseridRetrieveQuery = `SELECT * FROM UserIDs WHERE UserID = ?;`

var useridfromusernameRetrieveQuery = `SELECT * FROM UserIDs WHERE UserName = ?;`

var userRetrieveAllQuery = `SELECT * FROM Users;`

var useridRetrieveAllQuery = `SELECT * FROM UserIDs;`

var gameRetrieveQuery = `SELECT * FROM Games WHERE GameID = ?;`

var gameRetrieveAllQuery = `SELECT * FROM Games;`

var gamestateRetrieveQuery = `SELECT
	NextPlayer,
	NextPiece,
	Board,
	UnusedPieces
FROM Games WHERE GameID = ?;`

//TODO: order by timestamp
var invitedplayersRetrieveQuery = `SELECT * FROM InvitedPlayers WHERE GameID = ? ORDER BY InvitationTime DESCENDING;`

var activeplayersRetrieveQuery = `SELECT * FROM ActivePlayers WHERE GameID = ?;`

var userInsertQuery = `INSERT INTO Users (
	UserNickname,
	UserPassword
) VALUES (?, ?);`

var useridInsertQuery = `INSERT INTO UserIDs (
	UserNickname,
	UserID
) VALUES (?, ?);`

var gameInsertQuery = `INSERT INTO Games (
	GameID,
	ActivityStatus,
	InvitedPlayers,
	Board,
	UnusedPieces
) VALUES (?, ?, ?, ?, ?);`
