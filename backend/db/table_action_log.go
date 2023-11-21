package db

// action_log table
// id | timestamp | user_id | target_id | action | action_type
// represents an action taken by a user and some logged metadata
// it's intended for this feature to disablable or with retention settings

const (
	ActionLogEntryTypeUserLogin   = 100
	ActionLogEntryTypeUserCreated = 101

	ActionLogEntryTypeRoomCreate = 15000
	ActionLogEntryTypeRoomDelete = 15001
	ActionLogEntryTypeRoomUpdate = 15002

	ActionLogEntryTypeUserCreate = 15010
	ActionLogEntryTypeUserDelete = 15011
	ActionLogEntryTypeUserUpdate = 15012

	ActionLogEntryTypeMessageCreate = 15020
	ActionLogEntryTypeMessageDelete = 15021
	ActionLogEntryTypeMessageUpdate = 15022

	ActionLogEntryTypeRoomMemberAdd    = 15030
	ActionLogEntryTypeRoomMemberRemove = 15031
)

type ActionLogEntry struct {
	Id         int32  `db:"id"`
	Timestamp  int64  `db:"timestamp"`
	UserId     int32  `db:"user_id"`
	TargetId   int32  `db:"target_id"`
	Action     string `db:"action"`
	ActionType int32  `db:"action_type"`
}
