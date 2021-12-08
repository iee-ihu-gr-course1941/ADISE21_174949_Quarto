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

var createQuartoPieceTableQuery = `CREATE TABLE if not exists QuartoPieces (
	ID INTEGER NOT NULL,
	Dark BOOLEAN NOT NULL,
	Short BOOLEAN NOT NULL,
	Hollow BOOLEAN NOT NULL,
	Round BOOLEAN NOT NULL,
	PRIMARY KEY (ID)
);`

var createGameTableQuery = `CREATE TABLE if not exists Games (
	GameID VARCHAR(100) PRIMARY KEY NOT NULL,
	ActivityStatus BOOLEAN NOT NULL DEFAULT FALSE,
	Winner VARCHAR(100) REFERENCES UserIDs(UserNickname),
	ActivePlayers JSON,
	InvitedPlayers JSON,
	NextPlayer VARCHAR(100) REFERENCES Users(UserNickname),
	NextPiece JSON,
	BoardID INTEGER REFERENCES Boards(BoardID),
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


//could also be INTEGER UNIQUE
var createBoardTableQuery = `CREATE TABLE if not exists Boards (
	BoardID INTEGER AUTO_INCREMENT NOT NULL,

	x0y0 INTEGER REFERENCES QuartoPieces(ID),
	x0y1 INTEGER REFERENCES QuartoPieces(ID),
	x0y2 INTEGER REFERENCES QuartoPieces(ID),
	x0y3 INTEGER REFERENCES QuartoPieces(ID),

	x1y0 INTEGER REFERENCES QuartoPieces(ID),
	x1y1 INTEGER REFERENCES QuartoPieces(ID),
	x1y2 INTEGER REFERENCES QuartoPieces(ID),
	x1y3 INTEGER REFERENCES QuartoPieces(ID),

	x2y0 INTEGER REFERENCES QuartoPieces(ID),
	x2y1 INTEGER REFERENCES QuartoPieces(ID),
	x2y2 INTEGER REFERENCES QuartoPieces(ID),
	x2y3 INTEGER REFERENCES QuartoPieces(ID),

	x3y0 INTEGER REFERENCES QuartoPieces(ID),
	x3y1 INTEGER REFERENCES QuartoPieces(ID),
	x3y2 INTEGER REFERENCES QuartoPieces(ID),
	x3y3 INTEGER REFERENCES QuartoPieces(ID),

	PRIMARY KEY (BoardID)
);`

var useridfromuseridRetrieveQuery = `SELECT * FROM UserIDs WHERE UserID = ?;`

var useridfromusernameRetrieveQuery = `SELECT * FROM UserIDs WHERE UserName = ?;`

var userRetrieveAllQuery = `SELECT * FROM Users;`

var useridRetrieveAllQuery = `SELECT * FROM UserIDs;`

var gameRetrieveQuery = `SELECT * FROM Games WHERE GameID = ?;`

var gameRetrieveAllQuery = `SELECT * FROM Games;`

//TODO: order by timestamp
//var invitedplayersRetrieveQuery = `SELECT * FROM InvitedPlayers WHERE GameID = ? ORDER BY InvitationTime DESCENDING;`
//
//var activeplayersRetrieveQuery = `SELECT * FROM ActivePlayers WHERE GameID = ?;`

//alt impl
var invitedplayersRetrieveQuery = `SELECT InvitedPlayers FROM Games WHERE GameID = ?;`

var activeplayersRetrieveQuery = `SELECT ActivePlayers FROM Games WHERE GameID = ?;`

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

var gameUpdateQuery = `UPDATE Games SET InvitedPlayers = ? WHERE GameID = ?;`
