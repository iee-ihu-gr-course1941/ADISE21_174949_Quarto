Table of Contents
=================
   * [Εγκατάσταση](#εγκατάσταση)
      * [Απαιτήσεις](#απαιτήσεις)
      * [Οδηγίες Εγκατάστασης](#οδηγίες-εγκατάστασης)
   * [Περιγραφή API](#περιγραφή-api)
      * [Methods](#methods)
         * [Board](#board)
            * [Ανάγνωση Board](#ανάγνωση-board)
            * [Αρχικοποίηση Board](#αρχικοποίηση-board)
         * [Piece](#piece)
            * [Ανάγνωση Θέσης/Πιονιού](#ανάγνωση-θέσηςπιονιού)
            * [Μεταβολή Θέσης Πιονιού](#μεταβολή-θέσης-πιονιού)
         * [Player](#player)
            * [Ανάγνωση στοιχείων παίκτη](#ανάγνωση-στοιχείων-παίκτη)
            * [Καθορισμός στοιχείων παίκτη](#καθορισμός-στοιχείων-παίκτη)
         * [Status](#status)
            * [Ανάγνωση κατάστασης παιχνιδιού](#ανάγνωση-κατάστασης-παιχνιδιού)
      * [Entities](#entities)
         * [Board](#board-1)
         * [Players](#players)
         * [Game_status](#game_status)


# Demo Page

Μπορείτε να κατεβάσετε τοπικά ή να επισκευτείτε την σελίδα: 
https://users.iee.ihu.gr/~it174949/index.php

# Εγκατάσταση

## Απαιτήσεις

* Go Compiler
* MySQL Server

## Οδηγίες Εγκατάστασης

* Κάντε clone το project σε κάποιον φάκελο
 ```
 $ git clone https://github.com/iee-ihu-gr-course1941/ADISE21_174949_Quarto.git
 $ cd ADISE21_174949_Quarto
 $ export MYSQL_URL="your mysql url"
 $ go run main.go
 ````
* Εαν τρεχει στο users υπαρχει README εκει για αλλες ρυθμισεις

# Περιγραφή Παιχνιδιού

Οι κανόνες είναι:
- a
- b
- c

Η βάση μας κρατάει τους εξής πίνακες και στοιχεία:
```sql
CREATE TABLE if not exists Users (
	UserNickname VARCHAR(100) NOT NULL,
	UserPassword VARCHAR(100) NOT NULL,
	PRIMARY KEY (UserNickname)
);

CREATE TABLE if not exists UserIDs (
	UserNickname VARCHAR(100) NOT NULL REFERENCES Users(UserNickname),
	UserID VARCHAR(100) NOT NULL,
	PRIMARY KEY (UserNickname)
);

CREATE TABLE if not exists Games (
	GameID VARCHAR(100) PRIMARY KEY NOT NULL,
	ActivityStatus BOOLEAN NOT NULL DEFAULT FALSE,
	Winner VARCHAR(100) DEFAULT '' REFERENCES UserIDs(UserNickname),
	NextPlayer VARCHAR(100) REFERENCES UserIDs(UserNickname),
	NextPiece INTEGER,
	BoardID INTEGER REFERENCES Boards(BoardID),
	UnusedPiecesID INTEGER REFERENCES UnusedPieces(UnusedPiecesID)
);

CREATE TABLE if not exists InvitedPlayers (
	GameID VARCHAR(100) NOT NULL REFERENCES Games(GameID),
	UserName VARCHAR(100) NOT NULL REFERENCES UserIDs(UserNickname),
	InvitationTime TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (GameID, UserName)
);

CREATE TABLE if not exists ActivePlayers (
	GameID VARCHAR(100) NOT NULL REFERENCES Games(GameID),
	UserName VARCHAR(100) NOT NULL REFERENCES UserIDs(UserNickname),
	JoinTime TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (GameID, UserName)
);

CREATE TABLE if not exists Boards (
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
);

CREATE TABLE if not exists UnusedPieces (
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
);
```

Η εφαρμογη αναπτύχθηκε σε σημειο που:
- μπορουν 2 ατομα να παιξουν κανονικα μεταξυ τους
- επιστρεφεται ενα error σε περιπτωση προβληματος
- μπορει να ξεκινησει να χρησιμοποιειται χωρις ιδιαιτερη δουλεια στη βαση (πινακες δημιουργουνται αυτοματα)
- υπαρχει script για test σεναριου οπου δυο παικτες παιζουν μεταξυ τους (δειτε `play.sh`)

Θα μπορουσε να βελτιωθει με:
- λιγο πιο περιγραφικα error
- τα αντικειμενα `QuartoPiece` να μην εχουν κεφαλαια στο json ονομα τους

## Συντελεστές

Προγραμματιστής 1: Εγω :')

# Περιγραφή API

## Methods

### User
#### Δημιουργια User

```
POST /user
```

Δημιουργει ενα νεο χρηστη και επιστρεφει το μοναδικο ID του χρηστη που πρεπει να χρησιμοποιειται σε καθε αιτημα.

### Game
#### Δημιουργια Game

```
GET /game/{game_id}
```

Επιστρέφει ολα τα στοιχεια του παιχνιδιου.

#### Δημιουργια Game
```
POST /game
```

Δημιουργει ενα νεο παιχνιδι, προσκαλει τον δημιουργο του παιχνιδιου αυτοματα σε αυτο και επιστρεφει το ID του παιχνιδιου.

#### Προσκληση χρηστη σε Game

```
POST /game/{game_id}/invite/{username}
```

Προσκληση υπαρχον χρηστη σε υπαρχον παιχνιδι με βαση το username αυτου.

#### Ενεργοποιηση παικτη σε Game

```
POST /game/{game_id}/join
```

Ενας χρηστης που εχει προσκληθει σε ενα παιχνιδι σημειωνεται ως ενεργος και μπορει να παιξει.

#### Κινηση στο Game

```
POST /game/{game_id}/play
```

Αποστελλεται η επιθυμητη κινηση και εαν ειναι δυνατο να γινει, επιστρεφεται η κατασταση του παιχνιδιου.

## Entities
Τα αντικειμενα που χρησιμοποιουνται στην εφαρμογη υπαρχουν παρακατω.
Οι πινακες της βασεις δεδομενων αναφερθηκαν πιο πανω.
Στο attribute αναγραφεται το json ονομα που θα εμφανιστει κατα τη χρηση της.

### QuartoPiece
Το καθε πιονι ειναι ενα αντικειμενο με τα παρακατω στοιχεια.
Η "κενη" κατασταση ενος πιονου ειναι οταν το id του ειναι `-1`.

| Attribute                | Description                                 | Values                              |
| ------------------------ | ------------------------------------------- | ----------------------------------- |
| `Id`                     | Μοναδικο ID που εχει το καθε πιονι          | 0...16
| `Dark`                   | Εαν ειναι σκουροχρωμο ή οχι                 | boolean
| `Dark`                   | Εαν ειναι κοντο ή οχι                       | boolean
| `Hollow`                 | Εαν ειναι κουφιο ή οχι                      | boolean
| `Round`                  | Εαν ειναι στρογγυλο ή οχι                   | boolean


### Game
Το καθε παιχνιδι ειναι ενα αντικειμενο με τα παρακατω στοιχεια.

| Attribute                | Description                                 | Values                              |
| ------------------------ | ------------------------------------------- | ----------------------------------- |
| `game_id`                | Μοναδικο ID παιχνιδιου                      | String                              |
| `active_players`         | Πινακας ενεργων στο παιχνιδι παικτων        | Yx1 pointer UserId                  |
| `active_players`         | Πινακας προσκεκλημενων στο παιχνιδι παικτων | Yx1 pointer UserId                  |
| `activity_status`        | Εαν το παιχνιδι ειναι ενεργο                | boolean                             |
| `next_player`            | Ποιος παικτης πρεπει να παιξει              | pointer UserId                      |
| `next_piece`             | Ποιο πιονι πρεπει να παιχτει                | pointer QuartoPiece                 |
| `board`                  | 4χ4 πινακας του ταμπλο με τα πιονα          | 4x4 pointer QuartoPiece             |
| `unused_pieces`          | Πινακας απο πιονια που δε χρησιμοποιηθηκαν  | 16x1 pointer QuartoPiece            |
| `winner`                 | Παικτης που κερδισε εαν υπαρχει             | pointer UserId                      |

### GameMove
Η καθε κινηση στο παιχνιδι περιεχει τα παρακατω στοιχεια.

| Attribute                | Description                                | Values      |
| ------------------------ | ------------------------------------------ | ----------- |
| `username`               | Όνομα παίκτη                               | string      |
| `user_id`                | Μοναδικο ID παικτη                         | string      |
| `position_x`             | Θεση στον αξονα X του πιονου που παιζεται  | 0...3       |
| `position_y`             | Θεση στον αξονα Y του πιονου που παιζεται  | 0...3       |
| `next_piece              | Πιονι που θα πρεπει να παιξει ο επομενος   | QuartoPiece |

### User
O καθε παικτης αντιπροσωπευεται απο παρακατω στοιχεια οταν κανει εγγραφη.

| Attribute                | Description     | Values  |
| ------------------------ | --------------- | ------- |
| `username`               | Όνομα παίκτη    | string  |
| `password`               | Κωδικος παικτη  | string  |

### UserId
O καθε παικτης αντιπροσωπευεται απο παρακατω στοιχεια οταν κανει οποιοδηποτε αλλο αιτημα.

| Attribute                | Description         | Values  |
| ------------------------ | ------------------- | ------- |
| `username`               | Όνομα παίκτη        | string  |
| `user_id`                | Μοναδικο ID παικτη  | string  |
