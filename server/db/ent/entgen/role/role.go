// Code generated by entc, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldHoist holds the string denoting the hoist field in the database.
	FieldHoist = "hoist"
	// FieldPingable holds the string denoting the pingable field in the database.
	FieldPingable = "pingable"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// EdgeGuild holds the string denoting the guild edge name in mutations.
	EdgeGuild = "guild"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgePermissionNode holds the string denoting the permission_node edge name in mutations.
	EdgePermissionNode = "permission_node"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// GuildTable is the table the holds the guild relation/edge. The primary key declared below.
	GuildTable = "guild_role"
	// GuildInverseTable is the table name for the Guild entity.
	// It exists in this package in order to avoid circular dependency with the "guild" package.
	GuildInverseTable = "guilds"
	// MembersTable is the table the holds the members relation/edge. The primary key declared below.
	MembersTable = "role_members"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
	// PermissionNodeTable is the table the holds the permission_node relation/edge.
	PermissionNodeTable = "permission_nodes"
	// PermissionNodeInverseTable is the table name for the PermissionNode entity.
	// It exists in this package in order to avoid circular dependency with the "permissionnode" package.
	PermissionNodeInverseTable = "permission_nodes"
	// PermissionNodeColumn is the table column denoting the permission_node relation/edge.
	PermissionNodeColumn = "role_permission_node"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldColor,
	FieldHoist,
	FieldPingable,
	FieldPosition,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"channel_role",
}

var (
	// GuildPrimaryKey and GuildColumn2 are the table columns denoting the
	// primary key for the guild relation (M2M).
	GuildPrimaryKey = []string{"guild_id", "role_id"}
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"role_id", "user_id"}
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
