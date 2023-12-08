package db

type RecordRoom struct {
	Id          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	OwnerId     int32  `db:"owner_id"`
}

type RecordRoomMember struct {
	Id     int32 `db:"id"`
	RoomId int32 `db:"room_id"`
	UserId int32 `db:"user_id"`
	Joined int64 `db:"joined"`
}

func TableInitRoom(context TableInitContext) {
	/**
	 * room
	 */
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS room (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			owner_id INTEGER NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (owner_id) REFERENCES user_identity(id)
		);
	`)

	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_room_owner_id ON room (owner_id);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_room_name ON room (name);`)

	/**
	 * room_member
	 */
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS room_member (
			id SERIAL PRIMARY KEY,
			room_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			joined TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (room_id) REFERENCES room(id),
			FOREIGN KEY (user_id) REFERENCES user_identity(id)
		);
	`)

	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_room_member_room_id ON room_member (room_id);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_room_member_user_id ON room_member (user_id);`)
}

func InsertRoom(name string, description string, owner_id int32) {
	_, err := Con.NamedExec(`
		INSERT INTO room (name, description, owner_id)
		VALUES (:name, :description, :owner_id)
	`, map[string]interface{}{
		"name":        name,
		"description": description,
		"owner_id":    owner_id,
	})
	if err != nil {
		panic(err)
	}
}

func IsRoomMember(room_id int32, user_id int32) bool {
	var count int32
	err := Con.Get(&count, "SELECT COUNT(*) FROM room_member WHERE room_id = $1 AND user_id = $2", room_id, user_id)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (r *RecordRoom) IsMemberById(id int32) bool {
	var count int32
	err := Con.Get(&count,
		"SELECT COUNT(*) FROM room_member WHERE room_id = $1 AND user_id = $2",
		r.Id,
		id,
	)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (r *RecordRoom) IsMember(username string) bool {
	user, err := GetUser(username)
	if err != nil {
		panic(err)
	}
	return r.IsMemberById(user.Id)
}

func (r *RecordRoom) Owner() RecordUser {
	owner, err := GetUserById(r.OwnerId)
	if err != nil {
		panic(err)
	}
	return owner
}

func (r *RecordRoom) Members() []*RecordUser {
	var users []*RecordUser
	err := Con.Select(&users, "SELECT * FROM user WHERE id IN (SELECT user_id FROM room_members WHERE room_id = $1)", r.Id)
	if err != nil {
		panic(err)
	}
	return users
}
