package db

type Room struct {
	Id          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	OwnerId     int32  `db:"owner_id"`
}

type RoomMember struct {
	Id     int32 `db:"id"`
	RoomId int32 `db:"room_id"`
	UserId int32 `db:"user_id"`
	Joined int64 `db:"joined"`
}

func (r *Room) Owner() *User {
	return DBGetUser(r.OwnerId)
}

func (r *Room) Members() []*User {
	var users []*User
	err := Con.Select(&users, "SELECT * FROM users WHERE id IN (SELECT user_id FROM room_members WHERE room_id = $1)", r.Id)
	if err != nil {
		panic(err)
	}
	return users
}

// is user member of this room
func (r *Room) IsMember(username string) []*Message {
	var messages []*Message
	err := Con.Select(&messages, "SELECT * FROM messages WHERE room_id = $1", r.Id)
	if err != nil {
		panic(err)
	}
	return messages
}

func DBInitializeRoomTable() {
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS rooms (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		owner_id INTEGER NOT NULL,
		created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (owner_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		panic(err)
	}

	_, err = Con.Exec(`
	CREATE TABLE IF NOT EXISTS room_members (
		id SERIAL PRIMARY KEY,
		room_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		joined TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (room_id) REFERENCES rooms(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		panic(err)
	}
}
