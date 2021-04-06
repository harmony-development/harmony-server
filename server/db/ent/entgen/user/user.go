// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeLocalUser holds the string denoting the local_user edge name in mutations.
	EdgeLocalUser = "local_user"
	// EdgeForeignUser holds the string denoting the foreign_user edge name in mutations.
	EdgeForeignUser = "foreign_user"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// EdgeGuild holds the string denoting the guild edge name in mutations.
	EdgeGuild = "guild"
	// EdgeEmotepack holds the string denoting the emotepack edge name in mutations.
	EdgeEmotepack = "emotepack"
	// EdgeCreatedpacks holds the string denoting the createdpacks edge name in mutations.
	EdgeCreatedpacks = "createdpacks"
	// EdgeListentry holds the string denoting the listentry edge name in mutations.
	EdgeListentry = "listentry"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LocalUserTable is the table the holds the local_user relation/edge.
	LocalUserTable = "local_users"
	// LocalUserInverseTable is the table name for the LocalUser entity.
	// It exists in this package in order to avoid circular dependency with the "localuser" package.
	LocalUserInverseTable = "local_users"
	// LocalUserColumn is the table column denoting the local_user relation/edge.
	LocalUserColumn = "user_local_user"
	// ForeignUserTable is the table the holds the foreign_user relation/edge.
	ForeignUserTable = "foreign_users"
	// ForeignUserInverseTable is the table name for the ForeignUser entity.
	// It exists in this package in order to avoid circular dependency with the "foreignuser" package.
	ForeignUserInverseTable = "foreign_users"
	// ForeignUserColumn is the table column denoting the foreign_user relation/edge.
	ForeignUserColumn = "user_foreign_user"
	// ProfileTable is the table the holds the profile relation/edge.
	ProfileTable = "profiles"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "user_profile"
	// SessionsTable is the table the holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_sessions"
	// MessageTable is the table the holds the message relation/edge.
	MessageTable = "messages"
	// MessageInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessageInverseTable = "messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "user_message"
	// GuildTable is the table the holds the guild relation/edge. The primary key declared below.
	GuildTable = "user_guild"
	// GuildInverseTable is the table name for the Guild entity.
	// It exists in this package in order to avoid circular dependency with the "guild" package.
	GuildInverseTable = "guilds"
	// EmotepackTable is the table the holds the emotepack relation/edge.
	EmotepackTable = "emote_packs"
	// EmotepackInverseTable is the table name for the EmotePack entity.
	// It exists in this package in order to avoid circular dependency with the "emotepack" package.
	EmotepackInverseTable = "emote_packs"
	// EmotepackColumn is the table column denoting the emotepack relation/edge.
	EmotepackColumn = "user_emotepack"
	// CreatedpacksTable is the table the holds the createdpacks relation/edge.
	CreatedpacksTable = "emote_packs"
	// CreatedpacksInverseTable is the table name for the EmotePack entity.
	// It exists in this package in order to avoid circular dependency with the "emotepack" package.
	CreatedpacksInverseTable = "emote_packs"
	// CreatedpacksColumn is the table column denoting the createdpacks relation/edge.
	CreatedpacksColumn = "user_createdpacks"
	// ListentryTable is the table the holds the listentry relation/edge.
	ListentryTable = "guild_list_entries"
	// ListentryInverseTable is the table name for the GuildListEntry entity.
	// It exists in this package in order to avoid circular dependency with the "guildlistentry" package.
	ListentryInverseTable = "guild_list_entries"
	// ListentryColumn is the table column denoting the listentry relation/edge.
	ListentryColumn = "user_listentry"
	// RoleTable is the table the holds the role relation/edge. The primary key declared below.
	RoleTable = "role_members"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"guild_bans",
}

var (
	// GuildPrimaryKey and GuildColumn2 are the table columns denoting the
	// primary key for the guild relation (M2M).
	GuildPrimaryKey = []string{"user_id", "guild_id"}
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"role_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
