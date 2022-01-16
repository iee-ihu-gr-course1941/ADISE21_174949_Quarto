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

var createUnusedPieceTableQuery = `CREATE TABLE if not exists UnusedPieces (
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
	Winner VARCHAR(100) DEFAULT '' REFERENCES UserIDs(UserNickname),
	NextPlayer VARCHAR(100) REFERENCES UserIDs(UserNickname),
	NextPiece INTEGER,
	BoardID INTEGER REFERENCES Boards(BoardID),
	UnusedPiecesID INTEGER REFERENCES UnusedPieces(UnusedPiecesID)
);`

var createInvitedPlayerTableQuery = `CREATE TABLE if not exists InvitedPlayers (
	GameID VARCHAR(100) NOT NULL REFERENCES Games(GameID),
	UserName VARCHAR(100) NOT NULL REFERENCES UserIDs(UserNickname),
	InvitationTime TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (GameID, UserName)
);`

var createActivePlayerTableQuery = `CREATE TABLE if not exists ActivePlayers (
	GameID VARCHAR(100) NOT NULL REFERENCES Games(GameID),
	UserName VARCHAR(100) NOT NULL REFERENCES UserIDs(UserNickname),
	JoinTime TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (GameID, UserName)
);`

var createBoardTableQuery = `CREATE TABLE if not exists Boards (
	BoardID INTEGER AUTO_INCREMENT NOT NULL,
	x0y0 INTEGER DEFAULT -1,
	x0y1 INTEGER DEFAULT -1,
	x0y2 INTEGER DEFAULT -1,
	x0y3 INTEGER DEFAULT -1,
	x1y0 INTEGER DEFAULT -1,
	x1y1 INTEGER DEFAULT -1,
	x1y2 INTEGER DEFAULT -1,
	x1y3 INTEGER DEFAULT -1,
	x2y0 INTEGER DEFAULT -1,
	x2y1 INTEGER DEFAULT -1,
	x2y2 INTEGER DEFAULT -1,
	x2y3 INTEGER DEFAULT -1,
	x3y0 INTEGER DEFAULT -1,
	x3y1 INTEGER DEFAULT -1,
	x3y2 INTEGER DEFAULT -1,
	x3y3 INTEGER DEFAULT -1,
	PRIMARY KEY (BoardID)
);`

var createUnusedPiecesTableQuery = `CREATE TABLE if not exists UnusedPieces (
	UnusedPiecesID INTEGER AUTO_INCREMENT NOT NULL,
	up0 INTEGER DEFAULT 0,
	up1 INTEGER DEFAULT 1,
	up2 INTEGER DEFAULT 2,
	up3 INTEGER DEFAULT 3,
	up4 INTEGER DEFAULT 4,
	up5 INTEGER DEFAULT 5,
	up6 INTEGER DEFAULT 6,
	up7 INTEGER DEFAULT 7,
	up8 INTEGER DEFAULT 8,
	up9 INTEGER DEFAULT 9,
	up10 INTEGER DEFAULT 10,
	up11 INTEGER DEFAULT 11,
	up12 INTEGER DEFAULT 12,
	up13 INTEGER DEFAULT 13,
	up14 INTEGER DEFAULT 14,
	up15 INTEGER DEFAULT 15,
	PRIMARY KEY (UnusedPiecesID)
);`

var createEmptyBoardQuery = `INSERT INTO Boards () VALUES ();`

var createEmptyUnusedPiecesQuery = `INSERT INTO UnusedPieces () VALUES ();`

var useridfromuseridRetrieveQuery = `SELECT * FROM UserIDs WHERE UserID = ?;`

var useridfromusernameRetrieveQuery = `SELECT * FROM UserIDs WHERE UserNickName = ?;`

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
	UnusedPieces
) VALUES (?, ?, ?, ?, ?);`

var gameUpdateQuery = `UPDATE Games
	SET NextPlayer = ?,
		NextPiece = ?
	WHERE GameID = ?;`

var gameUpdateQueryWithWinner = `UPDATE Games
	SET ActivityStatus = ?,
		NextPlayer = ?,
		NextPiece = -1,
		Winner = ?
	WHERE GameID = ?;`

var boardUpdateQuery = `UPDATE Boards
	SET x0y0 = ?,
		x0y1 = ?,
		x0y2 = ?,
		x0y3 = ?,
		x1y0 = ?,
		x1y1 = ?,
		x1y2 = ?,
		x1y3 = ?,
		x2y0 = ?,
		x2y1 = ?,
		x2y2 = ?,
		x2y3 = ?,
		x3y0 = ?,
		x3y1 = ?,
		x3y2 = ?,
		x3y3 = ?
	WHERE BoardID = ?`
