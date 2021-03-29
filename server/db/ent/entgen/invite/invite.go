// Code generated by entc, DO NOT EDIT.

package invite

const (
	// Label holds the string label denoting the invite type in the database.
	Label = "invite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldUses holds the string denoting the uses field in the database.
	FieldUses = "uses"
	// FieldPossibleUses holds the string denoting the possible_uses field in the database.
	FieldPossibleUses = "possible_uses"
	// EdgeGuild holds the string denoting the guild edge name in mutations.
	EdgeGuild = "guild"
	// Table holds the table name of the invite in the database.
	Table = "invites"
	// GuildTable is the table the holds the guild relation/edge.
	GuildTable = "invites"
	// GuildInverseTable is the table name for the Guild entity.
	// It exists in this package in order to avoid circular dependency with the "guild" package.
	GuildInverseTable = "guilds"
	// GuildColumn is the table column denoting the guild relation/edge.
	GuildColumn = "guild_invite"
)

// Columns holds all SQL columns for invite fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldUses,
	FieldPossibleUses,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"guild_invite",
}

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

var (
	// DefaultUses holds the default value on creation for the "uses" field.
	DefaultUses int64
	// DefaultPossibleUses holds the default value on creation for the "possible_uses" field.
	DefaultPossibleUses int64
)
